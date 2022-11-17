package gamehandler

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/application/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/common/infrastructure/rediseventbus"
	"github.com/dum-dum-genius/game-of-liberty-computer/mainserver/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/integrationeventlistener"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/ztrue/tracerr"
)

type clientSession struct {
	liveGameId            uuid.UUID
	playerId              uuid.UUID
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

		gamesId, err := configuration.GameApplicationService.GetFirstGameId()
		if err != nil {
			return
		}
		liveGameId := gamesId

		clientSession := &clientSession{
			liveGameId:            liveGameId.GetId(),
			playerId:              uuid.New(),
			socketSendMessageLock: sync.RWMutex{},
		}

		gameInfoUpdatedListener, _ := integrationeventlistener.NewGameInfoUpdatedListener()
		gameInfoUpdatedListenerUnsubscriber := gameInfoUpdatedListener.Subscribe(
			clientSession.liveGameId,
			clientSession.playerId,
			func(event integrationeventlistener.GameInfoUpdatedIntegrationEvent) {
				handleGameInfoUpdatedEvent(conn, clientSession, event)
			},
		)
		defer gameInfoUpdatedListenerUnsubscriber()

		areaZoomedListener, _ := integrationeventlistener.NewAreaZoomedListener()
		areaZoomedListenerUnsubscriber := areaZoomedListener.Subscribe(clientSession.liveGameId, clientSession.playerId, func(event integrationeventlistener.AreaZoomedIntegrationEvent) {
			handleAreaZoomedEvent(conn, clientSession, event)
		})
		defer areaZoomedListenerUnsubscriber()

		zoomedAreaUpdatedListener, _ := integrationeventlistener.NewZoomedAreaUpdatedListener()
		zoomedAreaUpdatedListenerUnsubscriber := zoomedAreaUpdatedListener.Subscribe(clientSession.liveGameId, clientSession.playerId, func(event integrationeventlistener.ZoomedAreaUpdatedIntegrationEvent) {
			handleZoomedAreaUpdatedEvent(conn, clientSession, event)
		})
		defer zoomedAreaUpdatedListenerUnsubscriber()

		rediseventbus.NewRedisIntegrationEventBus(
			rediseventbus.WithRedisInfrastructureService[integrationeventlistener.AddPlayerRequestedIntegrationEvent](),
		).Publish(
			integrationeventlistener.AddPlayerRequestedListenerChannel,
			integrationeventlistener.NewAddPlayerRequestedIntegrationEvent(clientSession.liveGameId, clientSession.playerId),
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
				case ZoomAreaRequestedEventType:
					handleZoomAreaRequestedEvent(conn, clientSession, message)
				case ReviveUnitsRequestedEventType:
					handleReviveUnitsRequestedEvent(conn, clientSession, message)
				default:
				}
			}
		}()

		for {
			<-closeConnFlag

			rediseventbus.NewRedisIntegrationEventBus(
				rediseventbus.WithRedisInfrastructureService[integrationeventlistener.RemovePlayerRequestedIntegrationEvent](),
			).Publish(
				integrationeventlistener.RemovePlayerRequestedListenerChannel,
				integrationeventlistener.NewRemovePlayerRequestedIntegrationEvent(clientSession.liveGameId, clientSession.playerId),
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

func handleGameInfoUpdatedEvent(conn *websocket.Conn, clientSession *clientSession, event integrationeventlistener.GameInfoUpdatedIntegrationEvent) {
	sendJSONMessageToClient(conn, clientSession, gameHandlerPresenter.CreateInformationUpdatedEvent(event.Dimension))
}

func handleAreaZoomedEvent(conn *websocket.Conn, clientSession *clientSession, event integrationeventlistener.AreaZoomedIntegrationEvent) {
	sendJSONMessageToClient(conn, clientSession, gameHandlerPresenter.CreateAreaZoomedEvent(event.Area, event.UnitBlock))
}

func handleZoomedAreaUpdatedEvent(conn *websocket.Conn, clientSession *clientSession, event integrationeventlistener.ZoomedAreaUpdatedIntegrationEvent) {
	sendJSONMessageToClient(conn, clientSession, gameHandlerPresenter.CreateZoomedAreaUpdatedEvent(event.Area, event.UnitBlock))
}

func handleZoomAreaRequestedEvent(conn *websocket.Conn, clientSession *clientSession, event []byte) {
	areaPresenterDto, err := gameHandlerPresenter.ExtractZoomAreaRequestedEvent(event)
	if err != nil {
		emitErrorEvent(conn, clientSession, err)
		return
	}

	rediseventbus.NewRedisIntegrationEventBus(
		rediseventbus.WithRedisInfrastructureService[integrationeventlistener.ZoomAreaRequestedIntegrationEvent](),
	).Publish(
		integrationeventlistener.ZoomAreaRequestedListenerChannel,
		integrationeventlistener.NewZoomAreaRequestedIntegrationEvent(clientSession.liveGameId, clientSession.playerId, areaPresenterDto),
	)
}

func handleReviveUnitsRequestedEvent(conn *websocket.Conn, clientSession *clientSession, event []byte) {
	coordinatePresenterDtos, err := gameHandlerPresenter.ExtractReviveUnitsRequestedEvent(event)
	if err != nil {
		emitErrorEvent(conn, clientSession, err)
		return
	}

	rediseventbus.NewRedisIntegrationEventBus(
		rediseventbus.WithRedisInfrastructureService[integrationeventlistener.ReviveUnitsRequestedIntegrationEvent](),
	).Publish(
		integrationeventlistener.ReviveUnitsRequestedListenerChannel,
		integrationeventlistener.NewReviveUnitsRequestedIntegrationEvent(clientSession.liveGameId, coordinatePresenterDtos),
	)
}
