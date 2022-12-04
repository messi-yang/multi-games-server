package gamehandler

import (
	"encoding/json"
	"net/http"
	"sync"

	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/apiserver/application/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/apiserver/port/adapter/notification/redis"
	commonnotification "github.com/dum-dum-genius/game-of-liberty-computer/server/common/notification"
	commonredisdto "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification/redis/dto"
	commonservice "github.com/dum-dum-genius/game-of-liberty-computer/server/common/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/ztrue/tracerr"
)

type clientSession struct {
	liveGameId            livegamemodel.LiveGameId
	playerId              gamecommonmodel.PlayerId
	socketSendMessageLock sync.RWMutex
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type HandlerConfiguration struct {
	GameApplicationService service.GameApplicationService
	NotificationPublisher  commonnotification.NotificationPublisher
}

func NewHandler(configuration HandlerConfiguration) func(c *gin.Context) {
	return func(c *gin.Context) {
		conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.Error(err)
			return
		}
		defer conn.Close()
		closeConnFlag := make(chan bool)

		gameId, err := configuration.GameApplicationService.GetFirstGameId()
		if err != nil {
			return
		}

		clientSession := &clientSession{
			liveGameId:            livegamemodel.NewLiveGameId(gameId.GetId()),
			playerId:              gamecommonmodel.NewPlayerId(uuid.New()),
			socketSendMessageLock: sync.RWMutex{},
		}

		redisGameInfoUpdatedSubscriber, _ := redis.NewRedisGameInfoUpdatedSubscriber(clientSession.liveGameId, clientSession.playerId)
		redisGameInfoUpdatedSubscriberUnsubscriber := redisGameInfoUpdatedSubscriber.Subscribe(
			func(event commonredisdto.RedisGameInfoUpdatedEvent) {
				handleRedisGameInfoUpdatedEvent(conn, clientSession, event)
			},
		)
		defer redisGameInfoUpdatedSubscriberUnsubscriber()

		redisAreaZoomedSubscriber, _ := redis.NewRedisAreaZoomedSubscriber(clientSession.liveGameId, clientSession.playerId)
		redisAreaZoomedSubscriberUnsubscriber := redisAreaZoomedSubscriber.Subscribe(func(event commonredisdto.RedisAreaZoomedEvent) {
			handleRedisAreaZoomedEvent(conn, clientSession, event)
		})
		defer redisAreaZoomedSubscriberUnsubscriber()

		redisZoomedAreaUpdatedSubscriber, _ := redis.NewRedisZoomedAreaUpdatedSubscriber(clientSession.liveGameId, clientSession.playerId)
		redisZoomedAreaUpdatedSubscriberUnsubscriber := redisZoomedAreaUpdatedSubscriber.Subscribe(func(event commonredisdto.RedisZoomedAreaUpdatedEvent) {
			handleRedisZoomedAreaUpdatedEvent(conn, clientSession, event)
		})
		defer redisZoomedAreaUpdatedSubscriberUnsubscriber()

		configuration.NotificationPublisher.Publish(
			commonredisdto.NewRedisAddPlayerRequestedEventChannel(),
			commonredisdto.NewRedisAddPlayerRequestedEvent(clientSession.liveGameId, clientSession.playerId),
		)

		go func() {
			defer func() {
				closeConnFlag <- true
			}()

			for {
				_, compressedMessage, err := conn.ReadMessage()
				if err != nil {
					emitErrorEvent(conn, clientSession, err)
					break
				}

				gzipCompressor := commonservice.NewGzipService()
				message, err := gzipCompressor.Ungzip(compressedMessage)
				if err != nil {
					emitErrorEvent(conn, clientSession, err)
					break
				}

				eventType, err := gameHandlerPresenter.ExtractEventType(message)
				if err != nil {
					emitErrorEvent(conn, clientSession, err)
					break
				}

				switch eventType {
				case RedisZoomAreaRequestedEventType:
					area, err := gameHandlerPresenter.ExtractRedisZoomAreaRequestedEvent(message)
					if err != nil {
						emitErrorEvent(conn, clientSession, err)
						return
					}

					configuration.NotificationPublisher.Publish(
						commonredisdto.NewRedisZoomAreaRequestedEventChannel(),
						commonredisdto.NewRedisZoomAreaRequestedEvent(clientSession.liveGameId, clientSession.playerId, area),
					)
				case RedisReviveUnitsRequestedEventType:
					coordinatePresenters, err := gameHandlerPresenter.ExtractRedisReviveUnitsRequestedEvent(message)
					if err != nil {
						emitErrorEvent(conn, clientSession, err)
						return
					}

					configuration.NotificationPublisher.Publish(
						commonredisdto.NewRedisReviveUnitsRequestedEventChannel(),
						commonredisdto.NewRedisReviveUnitsRequestedEvent(clientSession.liveGameId, coordinatePresenters),
					)
				default:
				}
			}
		}()

		for {
			<-closeConnFlag

			configuration.NotificationPublisher.Publish(
				commonredisdto.NewRedisRemovePlayerRequestedEventChannel(),
				commonredisdto.NewRedisRemovePlayerRequestedEvent(clientSession.liveGameId, clientSession.playerId),
			)

			return
		}
	}
}

func sendJSONMessageToClient(conn *websocket.Conn, clientSession *clientSession, message any) {
	clientSession.socketSendMessageLock.Lock()
	defer clientSession.socketSendMessageLock.Unlock()

	messageJsonInBytes, _ := json.Marshal(message)

	gzipCompressor := commonservice.NewGzipService()
	compressedMessage, err := gzipCompressor.Gzip(messageJsonInBytes)
	if err != nil {
		emitErrorEvent(conn, clientSession, err)
		return
	}

	conn.WriteMessage(2, compressedMessage)
}

func emitErrorEvent(conn *websocket.Conn, clientSession *clientSession, err error) {
	tracerr.Print(tracerr.Wrap(err))
	sendJSONMessageToClient(conn, clientSession, gameHandlerPresenter.CreateErroredEvent(err.Error()))
}

func handleRedisGameInfoUpdatedEvent(conn *websocket.Conn, clientSession *clientSession, event commonredisdto.RedisGameInfoUpdatedEvent) {
	sendJSONMessageToClient(conn, clientSession, gameHandlerPresenter.CreateInformationUpdatedEvent(event.Dimension))
}

func handleRedisAreaZoomedEvent(conn *websocket.Conn, clientSession *clientSession, event commonredisdto.RedisAreaZoomedEvent) {
	sendJSONMessageToClient(conn, clientSession, gameHandlerPresenter.CreateRedisAreaZoomedEvent(event.Area, event.UnitBlock))
}

func handleRedisZoomedAreaUpdatedEvent(conn *websocket.Conn, clientSession *clientSession, event commonredisdto.RedisZoomedAreaUpdatedEvent) {
	sendJSONMessageToClient(conn, clientSession, gameHandlerPresenter.CreateRedisZoomedAreaUpdatedEvent(event.Area, event.UnitBlock))
}
