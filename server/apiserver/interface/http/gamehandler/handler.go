package gamehandler

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/module/common/compression"
	"github.com/dum-dum-genius/game-of-liberty-computer/module/common/notification"
	"github.com/dum-dum-genius/game-of-liberty-computer/module/game/domain/model/gamecommonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/module/game/domain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/apiserver/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/apiserver/port/adapter/notification/apiredis"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/gameserver/port/adapter/notification/gameredis"
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
	GameApplicationService applicationservice.GameApplicationService
	NotificationPublisher  notification.NotificationPublisher
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

		redisGameInfoUpdatedSubscriber, _ := apiredis.NewRedisGameInfoUpdatedSubscriber(clientSession.liveGameId, clientSession.playerId)
		redisGameInfoUpdatedSubscriberUnsubscriber := redisGameInfoUpdatedSubscriber.Subscribe(
			func(event apiredis.RedisGameInfoUpdatedIntegrationEvent) {
				handleRedisGameInfoUpdatedEvent(conn, clientSession, event)
			},
		)
		defer redisGameInfoUpdatedSubscriberUnsubscriber()

		redisAreaZoomedSubscriber, _ := apiredis.NewRedisAreaZoomedSubscriber(clientSession.liveGameId, clientSession.playerId)
		redisAreaZoomedSubscriberUnsubscriber := redisAreaZoomedSubscriber.Subscribe(func(event apiredis.RedisAreaZoomedIntegrationEvent) {
			handleRedisAreaZoomedEvent(conn, clientSession, event)
		})
		defer redisAreaZoomedSubscriberUnsubscriber()

		redisZoomedAreaUpdatedSubscriber, _ := apiredis.NewRedisZoomedAreaUpdatedSubscriber(clientSession.liveGameId, clientSession.playerId)
		redisZoomedAreaUpdatedSubscriberUnsubscriber := redisZoomedAreaUpdatedSubscriber.Subscribe(func(event apiredis.RedisZoomedAreaUpdatedIntegrationEvent) {
			handleRedisZoomedAreaUpdatedEvent(conn, clientSession, event)
		})
		defer redisZoomedAreaUpdatedSubscriberUnsubscriber()

		configuration.NotificationPublisher.Publish(
			gameredis.RedisAddPlayerRequestedSubscriberChannel,
			gameredis.NewRedisAddPlayerRequestedIntegrationEvent(clientSession.liveGameId, clientSession.playerId),
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

				gzipCompressor := compression.NewGzipCompressor()
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
						gameredis.RedisZoomAreaRequestedSubscriberChannel,
						gameredis.NewRedisZoomAreaRequestedIntegrationEvent(clientSession.liveGameId, clientSession.playerId, area),
					)
				case RedisReviveUnitsRequestedEventType:
					coordinatePresenters, err := gameHandlerPresenter.ExtractRedisReviveUnitsRequestedEvent(message)
					if err != nil {
						emitErrorEvent(conn, clientSession, err)
						return
					}

					configuration.NotificationPublisher.Publish(
						gameredis.RedisReviveUnitsRequestedSubscriberChannel,
						gameredis.NewRedisReviveUnitsRequestedIntegrationEvent(clientSession.liveGameId, coordinatePresenters),
					)
				default:
				}
			}
		}()

		for {
			<-closeConnFlag

			configuration.NotificationPublisher.Publish(
				gameredis.RedisRemovePlayerRequestedSubscriberChannel,
				gameredis.NewRedisRemovePlayerRequestedIntegrationEvent(clientSession.liveGameId, clientSession.playerId),
			)

			return
		}
	}
}

func sendJSONMessageToClient(conn *websocket.Conn, clientSession *clientSession, message any) {
	clientSession.socketSendMessageLock.Lock()
	defer clientSession.socketSendMessageLock.Unlock()

	messageJsonInBytes, _ := json.Marshal(message)

	gzipCompressor := compression.NewGzipCompressor()
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

func handleRedisGameInfoUpdatedEvent(conn *websocket.Conn, clientSession *clientSession, event apiredis.RedisGameInfoUpdatedIntegrationEvent) {
	sendJSONMessageToClient(conn, clientSession, gameHandlerPresenter.CreateInformationUpdatedEvent(event.Dimension))
}

func handleRedisAreaZoomedEvent(conn *websocket.Conn, clientSession *clientSession, event apiredis.RedisAreaZoomedIntegrationEvent) {
	sendJSONMessageToClient(conn, clientSession, gameHandlerPresenter.CreateRedisAreaZoomedEvent(event.Area, event.UnitBlock))
}

func handleRedisZoomedAreaUpdatedEvent(conn *websocket.Conn, clientSession *clientSession, event apiredis.RedisZoomedAreaUpdatedIntegrationEvent) {
	sendJSONMessageToClient(conn, clientSession, gameHandlerPresenter.CreateRedisZoomedAreaUpdatedEvent(event.Area, event.UnitBlock))
}
