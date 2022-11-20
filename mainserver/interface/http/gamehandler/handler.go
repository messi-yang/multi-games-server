package gamehandler

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/application/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/common/notification"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/gamecommonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/mainserver/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/messaging/redissubscriber"
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
	GameApplicationService *applicationservice.GameApplicationService
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

		redisGameInfoUpdatedSubscriber, _ := redissubscriber.NewRedisGameInfoUpdatedSubscriber(clientSession.liveGameId, clientSession.playerId)
		redisGameInfoUpdatedSubscriberUnsubscriber := redisGameInfoUpdatedSubscriber.Subscribe(
			func(event redissubscriber.RedisGameInfoUpdatedIntegrationEvent) {
				handleRedisGameInfoUpdatedEvent(conn, clientSession, event)
			},
		)
		defer redisGameInfoUpdatedSubscriberUnsubscriber()

		redisAreaZoomedSubscriber, _ := redissubscriber.NewRedisAreaZoomedSubscriber(clientSession.liveGameId, clientSession.playerId)
		redisAreaZoomedSubscriberUnsubscriber := redisAreaZoomedSubscriber.Subscribe(func(event redissubscriber.RedisAreaZoomedIntegrationEvent) {
			handleRedisAreaZoomedEvent(conn, clientSession, event)
		})
		defer redisAreaZoomedSubscriberUnsubscriber()

		redisZoomedAreaUpdatedSubscriber, _ := redissubscriber.NewRedisZoomedAreaUpdatedSubscriber(clientSession.liveGameId, clientSession.playerId)
		redisZoomedAreaUpdatedSubscriberUnsubscriber := redisZoomedAreaUpdatedSubscriber.Subscribe(func(event redissubscriber.RedisZoomedAreaUpdatedIntegrationEvent) {
			handleRedisZoomedAreaUpdatedEvent(conn, clientSession, event)
		})
		defer redisZoomedAreaUpdatedSubscriberUnsubscriber()

		configuration.NotificationPublisher.Publish(
			redissubscriber.RedisAddPlayerRequestedSubscriberChannel,
			redissubscriber.NewRedisAddPlayerRequestedIntegrationEvent(clientSession.liveGameId, clientSession.playerId),
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

				compressionApplicationService := service.NewCompressionApplicationService()
				message, err := compressionApplicationService.Ungzip(compressedMessage)
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
						redissubscriber.RedisZoomAreaRequestedSubscriberChannel,
						redissubscriber.NewRedisZoomAreaRequestedIntegrationEvent(clientSession.liveGameId, clientSession.playerId, area),
					)
				case RedisReviveUnitsRequestedEventType:
					coordinatePresenters, err := gameHandlerPresenter.ExtractRedisReviveUnitsRequestedEvent(message)
					if err != nil {
						emitErrorEvent(conn, clientSession, err)
						return
					}

					configuration.NotificationPublisher.Publish(
						redissubscriber.RedisReviveUnitsRequestedSubscriberChannel,
						redissubscriber.NewRedisReviveUnitsRequestedIntegrationEvent(clientSession.liveGameId, coordinatePresenters),
					)
				default:
				}
			}
		}()

		for {
			<-closeConnFlag

			configuration.NotificationPublisher.Publish(
				redissubscriber.RedisRemovePlayerRequestedSubscriberChannel,
				redissubscriber.NewRedisRemovePlayerRequestedIntegrationEvent(clientSession.liveGameId, clientSession.playerId),
			)

			return
		}
	}
}

func sendJSONMessageToClient(conn *websocket.Conn, clientSession *clientSession, message any) {
	clientSession.socketSendMessageLock.Lock()
	defer clientSession.socketSendMessageLock.Unlock()

	messageJsonInBytes, _ := json.Marshal(message)

	compressionApplicationService := service.NewCompressionApplicationService()
	compressedMessage, err := compressionApplicationService.Gzip(messageJsonInBytes)
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

func handleRedisGameInfoUpdatedEvent(conn *websocket.Conn, clientSession *clientSession, event redissubscriber.RedisGameInfoUpdatedIntegrationEvent) {
	sendJSONMessageToClient(conn, clientSession, gameHandlerPresenter.CreateInformationUpdatedEvent(event.Dimension))
}

func handleRedisAreaZoomedEvent(conn *websocket.Conn, clientSession *clientSession, event redissubscriber.RedisAreaZoomedIntegrationEvent) {
	sendJSONMessageToClient(conn, clientSession, gameHandlerPresenter.CreateRedisAreaZoomedEvent(event.Area, event.UnitBlock))
}

func handleRedisZoomedAreaUpdatedEvent(conn *websocket.Conn, clientSession *clientSession, event redissubscriber.RedisZoomedAreaUpdatedIntegrationEvent) {
	sendJSONMessageToClient(conn, clientSession, gameHandlerPresenter.CreateRedisZoomedAreaUpdatedEvent(event.Area, event.UnitBlock))
}
