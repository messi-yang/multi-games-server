package gameroomsockethandler

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/gameclientcommunicator/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/config"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/entity"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/eventbusredis"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/integrationevent"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/ztrue/tracerr"
)

type clientSession struct {
	watchedArea           *valueobject.Area
	socketSendMessageLock sync.RWMutex
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Handler(c *gin.Context) {
	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.Error(err)
		return
	}
	defer conn.Close()
	closeConnFlag := make(chan bool)

	integrationEventBusRedis := eventbusredis.GetIntegrationEventBusRedis()

	gameId := config.GetConfig().GetGameId()

	player := entity.NewPlayer()

	clientSession := &clientSession{
		watchedArea:           nil,
		socketSendMessageLock: sync.RWMutex{},
	}

	gameInfoUpdatedIntegrationEventSubscriber := integrationEventBusRedis.Subscribe(
		integrationevent.NewGameInfoUpdatedIntegrationEventTopic(gameId, player.GetId()),
		func(event []byte) {
			handleGameInfoUpdatedEvent(conn, clientSession, gameId, player.GetId(), event)
		},
	)
	defer gameInfoUpdatedIntegrationEventSubscriber()

	areaZoomedIntegrationEventUnsubscriber := integrationEventBusRedis.Subscribe(
		integrationevent.NewAreaZoomedIntegrationEventTopic(gameId, player.GetId()),
		func(event []byte) {
			handleAreaZoomedEvent(conn, clientSession, gameId, player.GetId(), event)
		},
	)
	defer areaZoomedIntegrationEventUnsubscriber()

	zoomedAreaUpdatedIntegrationEventUnsubscriber := integrationEventBusRedis.Subscribe(
		integrationevent.NewZoomedAreaUpdatedIntegrationEventTopic(gameId, player.GetId()),
		func(event []byte) {
			handleZoomedAreaUpdatedEvent(conn, clientSession, gameId, player.GetId(), event)
		},
	)
	defer zoomedAreaUpdatedIntegrationEventUnsubscriber()

	integrationEventBusRedis.Publish(
		integrationevent.NewAddPlayerRequestedIntegrationEventTopic(gameId),
		integrationevent.NewAddPlayerRequestedIntegrationEvent(player),
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

			compressionApplicationService := applicationservice.NewCompressionApplicationService()
			message, err := compressionApplicationService.Ungzip(compressedMessage)
			if err != nil {
				emitErrorEvent(conn, clientSession, err)
				break
			}

			actionType, err := getActionTypeFromMessage(message)
			if err != nil {
				emitErrorEvent(conn, clientSession, err)
				break
			}

			switch *actionType {
			case zoomAreaActionType:
				handleZoomAreaAction(conn, clientSession, message, gameId, player.GetId())
			case reviveUnitsActionType:
				handleReviveUnitsAction(conn, clientSession, message, gameId)
			default:
			}
		}
	}()

	for {
		<-closeConnFlag

		integrationEventBusRedis.Publish(
			integrationevent.NewRemovePlayerRequestedIntegrationEventTopic(gameId),
			integrationevent.NewRemovePlayerRequestedIntegrationEvent(player.GetId()),
		)

		return
	}
}

func sendJSONMessageToClient(conn *websocket.Conn, clientSession *clientSession, message any) {
	clientSession.socketSendMessageLock.Lock()
	defer clientSession.socketSendMessageLock.Unlock()

	messageJsonInBytes, _ := json.Marshal(message)

	compressionApplicationService := applicationservice.NewCompressionApplicationService()
	compressedMessage, err := compressionApplicationService.Gzip(messageJsonInBytes)
	if err != nil {
		emitErrorEvent(conn, clientSession, err)
		return
	}

	conn.WriteMessage(2, compressedMessage)
}

func emitErrorEvent(conn *websocket.Conn, clientSession *clientSession, err error) {
	errorEvent := constructErrorHappenedEvent(err.Error())

	tracerr.Print(tracerr.Wrap(err))

	sendJSONMessageToClient(conn, clientSession, errorEvent)
}

func handleGameInfoUpdatedEvent(conn *websocket.Conn, clientSession *clientSession, gameId uuid.UUID, playerId uuid.UUID, event []byte) {
	var gameInfoUpdatedIntegrationEvent integrationevent.GameInfoUpdatedIntegrationEvent
	json.Unmarshal(event, &gameInfoUpdatedIntegrationEvent)

	informationUpdatedEvent := constructInformationUpdatedEvent(gameInfoUpdatedIntegrationEvent.Payload.MapSize)
	sendJSONMessageToClient(conn, clientSession, informationUpdatedEvent)
}

func handleAreaZoomedEvent(conn *websocket.Conn, clientSession *clientSession, gameId uuid.UUID, playerId uuid.UUID, event []byte) {
	var areaZoomedIntegrationEvent integrationevent.AreaZoomedIntegrationEvent
	json.Unmarshal(event, &areaZoomedIntegrationEvent)

	sendJSONMessageToClient(conn, clientSession, constructAreaZoomedEvent(areaZoomedIntegrationEvent.Payload.Area, areaZoomedIntegrationEvent.Payload.UnitMap))
}

func handleZoomedAreaUpdatedEvent(conn *websocket.Conn, clientSession *clientSession, gameId uuid.UUID, playerId uuid.UUID, event []byte) {
	var zoomedAreaUpdatedIntegrationEvent integrationevent.ZoomedAreaUpdatedIntegrationEvent
	json.Unmarshal(event, &zoomedAreaUpdatedIntegrationEvent)

	sendJSONMessageToClient(conn, clientSession, constructZoomedAreaUpdatedEvent(zoomedAreaUpdatedIntegrationEvent.Payload.Area, zoomedAreaUpdatedIntegrationEvent.Payload.UnitMap))
}

func handleZoomAreaAction(conn *websocket.Conn, clientSession *clientSession, message []byte, gameId uuid.UUID, playerId uuid.UUID) {
	area, err := extractInformationFromZoomAreaAction(message)
	if err != nil {
		emitErrorEvent(conn, clientSession, err)
		return
	}

	integrationEventBusRedis := eventbusredis.GetIntegrationEventBusRedis()
	integrationEventBusRedis.Publish(
		integrationevent.NewZoomAreaRequestedIntegrationEventTopic(gameId),
		integrationevent.NewZoomAreaRequestedIntegrationEvent(playerId, area),
	)
}

func handleReviveUnitsAction(conn *websocket.Conn, clientSession *clientSession, message []byte, gameId uuid.UUID) {
	coordinates, err := extractInformationFromReviveUnitsAction(message)
	if err != nil {
		emitErrorEvent(conn, clientSession, err)
		return
	}

	integrationEventBusRedis := eventbusredis.GetIntegrationEventBusRedis()
	integrationEventBusRedis.Publish(
		integrationevent.NewReviveUnitsRequestedIntegrationEventTopic(gameId),
		integrationevent.NewReviveUnitsRequestedIntegrationEvent(coordinates),
	)
}
