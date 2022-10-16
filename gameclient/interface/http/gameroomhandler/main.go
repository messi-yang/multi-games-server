package gameroomhandler

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/gameclient/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/gameclient/presenter/gameroomhandlerpresenter"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/application/eventbus"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/application/sharedapplicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/entity"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/integrationevent"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/ztrue/tracerr"
)

type clientSession struct {
	gameId                   uuid.UUID
	player                   entity.Player
	integrationEventBus      eventbus.IntegrationEventBus
	gameRoomHandlerPresenter gameroomhandlerpresenter.GameRoomHandlerPresenter
	socketSendMessageLock    sync.RWMutex
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type HandlerConfiguration struct {
	IntegrationEventBus      eventbus.IntegrationEventBus
	GameApplicationService   applicationservice.GameApplicationService
	GameRoomHandlerPresenter gameroomhandlerpresenter.GameRoomHandlerPresenter
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
			c.Error(err)
			return
		}

		clientSession := &clientSession{
			gameId:                   gameId,
			player:                   entity.NewPlayer(),
			integrationEventBus:      configuration.IntegrationEventBus,
			gameRoomHandlerPresenter: configuration.GameRoomHandlerPresenter,
			socketSendMessageLock:    sync.RWMutex{},
		}

		gameInfoUpdatedIntegrationEventSubscriber := clientSession.integrationEventBus.Subscribe(
			integrationevent.NewGameInfoUpdatedIntegrationEventTopic(clientSession.gameId, clientSession.player.GetId()),
			func(event []byte) {
				handleGameInfoUpdatedEvent(conn, clientSession, event)
			},
		)
		defer gameInfoUpdatedIntegrationEventSubscriber()

		areaZoomedIntegrationEventUnsubscriber := clientSession.integrationEventBus.Subscribe(
			integrationevent.NewAreaZoomedIntegrationEventTopic(clientSession.gameId, clientSession.player.GetId()),
			func(event []byte) {
				handleAreaZoomedEvent(conn, clientSession, event)
			},
		)
		defer areaZoomedIntegrationEventUnsubscriber()

		zoomedAreaUpdatedIntegrationEventUnsubscriber := clientSession.integrationEventBus.Subscribe(
			integrationevent.NewZoomedAreaUpdatedIntegrationEventTopic(clientSession.gameId, clientSession.player.GetId()),
			func(event []byte) {
				handleZoomedAreaUpdatedEvent(conn, clientSession, event)
			},
		)
		defer zoomedAreaUpdatedIntegrationEventUnsubscriber()

		clientSession.integrationEventBus.Publish(
			integrationevent.NewAddPlayerRequestedIntegrationEventTopic(clientSession.gameId),
			integrationevent.NewAddPlayerRequestedIntegrationEvent(clientSession.player),
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

				compressionApplicationService := sharedapplicationservice.NewCompressionApplicationService()
				message, err := compressionApplicationService.Ungzip(compressedMessage)
				if err != nil {
					emitErrorEvent(conn, clientSession, err)
					break
				}

				eventType, err := clientSession.gameRoomHandlerPresenter.ExtractEventType(message)
				if err != nil {
					emitErrorEvent(conn, clientSession, err)
					break
				}

				switch eventType {
				case gameroomhandlerpresenter.ZoomAreaRequestedEventType:
					handleZoomAreaRequestedEvent(conn, clientSession, message)
				case gameroomhandlerpresenter.ReviveUnitsRequestedEventType:
					handleReviveUnitsRequestedEvent(conn, clientSession, message)
				default:
				}
			}
		}()

		for {
			<-closeConnFlag

			clientSession.integrationEventBus.Publish(
				integrationevent.NewRemovePlayerRequestedIntegrationEventTopic(clientSession.gameId),
				integrationevent.NewRemovePlayerRequestedIntegrationEvent(clientSession.player.GetId()),
			)

			return
		}
	}
}

func sendJSONMessageToClient(conn *websocket.Conn, clientSession *clientSession, message any) {
	clientSession.socketSendMessageLock.Lock()
	defer clientSession.socketSendMessageLock.Unlock()

	messageJsonInBytes, _ := json.Marshal(message)

	compressionApplicationService := sharedapplicationservice.NewCompressionApplicationService()
	compressedMessage, err := compressionApplicationService.Gzip(messageJsonInBytes)
	if err != nil {
		emitErrorEvent(conn, clientSession, err)
		return
	}

	conn.WriteMessage(2, compressedMessage)
}

func emitErrorEvent(conn *websocket.Conn, clientSession *clientSession, err error) {
	tracerr.Print(tracerr.Wrap(err))
	sendJSONMessageToClient(conn, clientSession, clientSession.gameRoomHandlerPresenter.CreateErroredEvent(err.Error()))
}

func handleGameInfoUpdatedEvent(conn *websocket.Conn, clientSession *clientSession, event []byte) {
	var gameInfoUpdatedIntegrationEvent integrationevent.GameInfoUpdatedIntegrationEvent
	json.Unmarshal(event, &gameInfoUpdatedIntegrationEvent)

	sendJSONMessageToClient(conn, clientSession, clientSession.gameRoomHandlerPresenter.CreateInformationUpdatedEvent(gameInfoUpdatedIntegrationEvent.Payload.MapSize))
}

func handleAreaZoomedEvent(conn *websocket.Conn, clientSession *clientSession, event []byte) {
	var areaZoomedIntegrationEvent integrationevent.AreaZoomedIntegrationEvent
	json.Unmarshal(event, &areaZoomedIntegrationEvent)

	sendJSONMessageToClient(conn, clientSession, clientSession.gameRoomHandlerPresenter.CreateAreaZoomedEvent(areaZoomedIntegrationEvent.Payload.Area, areaZoomedIntegrationEvent.Payload.UnitMap))
}

func handleZoomedAreaUpdatedEvent(conn *websocket.Conn, clientSession *clientSession, event []byte) {
	var zoomedAreaUpdatedIntegrationEvent integrationevent.ZoomedAreaUpdatedIntegrationEvent
	json.Unmarshal(event, &zoomedAreaUpdatedIntegrationEvent)

	sendJSONMessageToClient(conn, clientSession, clientSession.gameRoomHandlerPresenter.CreateZoomedAreaUpdatedEvent(zoomedAreaUpdatedIntegrationEvent.Payload.Area, zoomedAreaUpdatedIntegrationEvent.Payload.UnitMap))
}

func handleZoomAreaRequestedEvent(conn *websocket.Conn, clientSession *clientSession, event []byte) {
	area, err := clientSession.gameRoomHandlerPresenter.ExtractZoomAreaRequestedEvent(event)
	if err != nil {
		emitErrorEvent(conn, clientSession, err)
		return
	}

	clientSession.integrationEventBus.Publish(
		integrationevent.NewZoomAreaRequestedIntegrationEventTopic(clientSession.gameId),
		integrationevent.NewZoomAreaRequestedIntegrationEvent(clientSession.player.GetId(), area),
	)
}

func handleReviveUnitsRequestedEvent(conn *websocket.Conn, clientSession *clientSession, event []byte) {
	coordinates, err := clientSession.gameRoomHandlerPresenter.ExtractReviveUnitsRequestedEvent(event)
	if err != nil {
		emitErrorEvent(conn, clientSession, err)
		return
	}

	clientSession.integrationEventBus.Publish(
		integrationevent.NewReviveUnitsRequestedIntegrationEventTopic(clientSession.gameId),
		integrationevent.NewReviveUnitsRequestedIntegrationEvent(coordinates),
	)
}
