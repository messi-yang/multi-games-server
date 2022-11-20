package gamehandler

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/application/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/common/port/adapter/notification/commonredisnotification"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/gamecommonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/mainserver/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/messaging/redislistener"
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

		redisGameInfoUpdatedListener, _ := redislistener.NewRedisGameInfoUpdatedListener(clientSession.liveGameId, clientSession.playerId)
		redisGameInfoUpdatedListenerUnsubscriber := redisGameInfoUpdatedListener.Subscribe(
			func(event redislistener.RedisGameInfoUpdatedIntegrationEvent) {
				handleRedisGameInfoUpdatedEvent(conn, clientSession, event)
			},
		)
		defer redisGameInfoUpdatedListenerUnsubscriber()

		redisAreaZoomedListener, _ := redislistener.NewRedisAreaZoomedListener(clientSession.liveGameId, clientSession.playerId)
		redisAreaZoomedListenerUnsubscriber := redisAreaZoomedListener.Subscribe(func(event redislistener.RedisAreaZoomedIntegrationEvent) {
			handleRedisAreaZoomedEvent(conn, clientSession, event)
		})
		defer redisAreaZoomedListenerUnsubscriber()

		redisZoomedAreaUpdatedListener, _ := redislistener.NewRedisZoomedAreaUpdatedListener(clientSession.liveGameId, clientSession.playerId)
		redisZoomedAreaUpdatedListenerUnsubscriber := redisZoomedAreaUpdatedListener.Subscribe(func(event redislistener.RedisZoomedAreaUpdatedIntegrationEvent) {
			handleRedisZoomedAreaUpdatedEvent(conn, clientSession, event)
		})
		defer redisZoomedAreaUpdatedListenerUnsubscriber()

		commonredisnotification.NewRedisNotificationPublisher().Publish(
			redislistener.RedisAddPlayerRequestedListenerChannel,
			redislistener.NewRedisAddPlayerRequestedIntegrationEvent(clientSession.liveGameId, clientSession.playerId),
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
					handleRedisZoomAreaRequestedEvent(conn, clientSession, message)
				case RedisReviveUnitsRequestedEventType:
					handleRedisReviveUnitsRequestedEvent(conn, clientSession, message)
				default:
				}
			}
		}()

		for {
			<-closeConnFlag

			commonredisnotification.NewRedisNotificationPublisher().Publish(
				redislistener.RedisRemovePlayerRequestedListenerChannel,
				redislistener.NewRedisRemovePlayerRequestedIntegrationEvent(clientSession.liveGameId, clientSession.playerId),
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

func handleRedisGameInfoUpdatedEvent(conn *websocket.Conn, clientSession *clientSession, event redislistener.RedisGameInfoUpdatedIntegrationEvent) {
	sendJSONMessageToClient(conn, clientSession, gameHandlerPresenter.CreateInformationUpdatedEvent(event.Dimension))
}

func handleRedisAreaZoomedEvent(conn *websocket.Conn, clientSession *clientSession, event redislistener.RedisAreaZoomedIntegrationEvent) {
	sendJSONMessageToClient(conn, clientSession, gameHandlerPresenter.CreateRedisAreaZoomedEvent(event.Area, event.UnitBlock))
}

func handleRedisZoomedAreaUpdatedEvent(conn *websocket.Conn, clientSession *clientSession, event redislistener.RedisZoomedAreaUpdatedIntegrationEvent) {
	sendJSONMessageToClient(conn, clientSession, gameHandlerPresenter.CreateRedisZoomedAreaUpdatedEvent(event.Area, event.UnitBlock))
}

func handleRedisZoomAreaRequestedEvent(conn *websocket.Conn, clientSession *clientSession, event []byte) {
	area, err := gameHandlerPresenter.ExtractRedisZoomAreaRequestedEvent(event)
	if err != nil {
		emitErrorEvent(conn, clientSession, err)
		return
	}

	commonredisnotification.NewRedisNotificationPublisher().Publish(
		redislistener.RedisZoomAreaRequestedListenerChannel,
		redislistener.NewRedisZoomAreaRequestedIntegrationEvent(clientSession.liveGameId, clientSession.playerId, area),
	)
}

func handleRedisReviveUnitsRequestedEvent(conn *websocket.Conn, clientSession *clientSession, event []byte) {
	coordinatePresenters, err := gameHandlerPresenter.ExtractRedisReviveUnitsRequestedEvent(event)
	if err != nil {
		emitErrorEvent(conn, clientSession, err)
		return
	}

	commonredisnotification.NewRedisNotificationPublisher().Publish(
		redislistener.RedisReviveUnitsRequestedListenerChannel,
		redislistener.NewRedisReviveUnitsRequestedIntegrationEvent(clientSession.liveGameId, coordinatePresenters),
	)
}
