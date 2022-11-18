package gamehandler

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/application/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/common/infrastructure/rediseventbus"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/gamecommonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/mainserver/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/messaging/redis"
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

		redisGameInfoUpdatedListener, _ := redis.NewRedisGameInfoUpdatedListener()
		redisGameInfoUpdatedListenerUnsubscriber := redisGameInfoUpdatedListener.Subscribe(
			clientSession.liveGameId,
			clientSession.playerId,
			func(event redis.RedisGameInfoUpdatedIntegrationEvent) {
				handleRedisGameInfoUpdatedEvent(conn, clientSession, event)
			},
		)
		defer redisGameInfoUpdatedListenerUnsubscriber()

		redisAreaZoomedListener, _ := redis.NewRedisAreaZoomedListener()
		redisAreaZoomedListenerUnsubscriber := redisAreaZoomedListener.Subscribe(clientSession.liveGameId, clientSession.playerId, func(event redis.RedisAreaZoomedIntegrationEvent) {
			handleRedisAreaZoomedEvent(conn, clientSession, event)
		})
		defer redisAreaZoomedListenerUnsubscriber()

		redisZoomedAreaUpdatedListener, _ := redis.NewRedisZoomedAreaUpdatedListener()
		redisZoomedAreaUpdatedListenerUnsubscriber := redisZoomedAreaUpdatedListener.Subscribe(clientSession.liveGameId, clientSession.playerId, func(event redis.RedisZoomedAreaUpdatedIntegrationEvent) {
			handleRedisZoomedAreaUpdatedEvent(conn, clientSession, event)
		})
		defer redisZoomedAreaUpdatedListenerUnsubscriber()

		rediseventbus.NewRedisIntegrationEventBus(
			rediseventbus.WithRedisInfrastructureService[redis.RedisAddPlayerRequestedIntegrationEvent](),
		).Publish(
			redis.RedisAddPlayerRequestedListenerChannel,
			redis.NewRedisAddPlayerRequestedIntegrationEvent(clientSession.liveGameId, clientSession.playerId),
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

			rediseventbus.NewRedisIntegrationEventBus(
				rediseventbus.WithRedisInfrastructureService[redis.RedisRemovePlayerRequestedIntegrationEvent](),
			).Publish(
				redis.RedisRemovePlayerRequestedListenerChannel,
				redis.NewRedisRemovePlayerRequestedIntegrationEvent(clientSession.liveGameId, clientSession.playerId),
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

func handleRedisGameInfoUpdatedEvent(conn *websocket.Conn, clientSession *clientSession, event redis.RedisGameInfoUpdatedIntegrationEvent) {
	sendJSONMessageToClient(conn, clientSession, gameHandlerPresenter.CreateInformationUpdatedEvent(event.Dimension))
}

func handleRedisAreaZoomedEvent(conn *websocket.Conn, clientSession *clientSession, event redis.RedisAreaZoomedIntegrationEvent) {
	sendJSONMessageToClient(conn, clientSession, gameHandlerPresenter.CreateRedisAreaZoomedEvent(event.Area, event.UnitBlock))
}

func handleRedisZoomedAreaUpdatedEvent(conn *websocket.Conn, clientSession *clientSession, event redis.RedisZoomedAreaUpdatedIntegrationEvent) {
	sendJSONMessageToClient(conn, clientSession, gameHandlerPresenter.CreateRedisZoomedAreaUpdatedEvent(event.Area, event.UnitBlock))
}

func handleRedisZoomAreaRequestedEvent(conn *websocket.Conn, clientSession *clientSession, event []byte) {
	areaPresenterDto, err := gameHandlerPresenter.ExtractRedisZoomAreaRequestedEvent(event)
	if err != nil {
		emitErrorEvent(conn, clientSession, err)
		return
	}

	rediseventbus.NewRedisIntegrationEventBus(
		rediseventbus.WithRedisInfrastructureService[redis.RedisZoomAreaRequestedIntegrationEvent](),
	).Publish(
		redis.RedisZoomAreaRequestedListenerChannel,
		redis.NewRedisZoomAreaRequestedIntegrationEvent(clientSession.liveGameId, clientSession.playerId, areaPresenterDto),
	)
}

func handleRedisReviveUnitsRequestedEvent(conn *websocket.Conn, clientSession *clientSession, event []byte) {
	coordinatePresenterDtos, err := gameHandlerPresenter.ExtractRedisReviveUnitsRequestedEvent(event)
	if err != nil {
		emitErrorEvent(conn, clientSession, err)
		return
	}

	rediseventbus.NewRedisIntegrationEventBus(
		rediseventbus.WithRedisInfrastructureService[redis.RedisReviveUnitsRequestedIntegrationEvent](),
	).Publish(
		redis.RedisReviveUnitsRequestedListenerChannel,
		redis.NewRedisReviveUnitsRequestedIntegrationEvent(clientSession.liveGameId, coordinatePresenterDtos),
	)
}
