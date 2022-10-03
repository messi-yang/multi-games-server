package gameroomsockethandler

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/gameclientcommunicator/application/service/compressionservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/application/service/gameroomservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/infrastructure/memory/gameroommemory"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/config"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/entity"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/memoryeventbus"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/event/addplayerrequestedevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/event/areazoomedevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/event/removeplayerrequestedevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/event/reviveunitsrequestedevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/event/zoomarearequestedevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/event/zoomedareaupdatedevent"
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

	eventBus := memoryeventbus.GetEventBus()

	gameId := config.GetConfig().GetGameId()

	player := entity.NewPlayer()

	eventBus.Publish(
		addplayerrequestedevent.NewEventTopic(gameId),
		addplayerrequestedevent.NewEvent(player),
	)

	clientSession := &clientSession{
		watchedArea:           nil,
		socketSendMessageLock: sync.RWMutex{},
	}

	emitGameInfoUpdatedEvent(conn, clientSession, gameId)

	areaZoomedEventUnsubscriber := eventBus.Subscribe(areazoomedevent.NewEventTopic(gameId, player.GetId()), func(event []byte) {
		handleAreaZoomedEvent(conn, clientSession, gameId, player.GetId(), event)
	})
	defer areaZoomedEventUnsubscriber()

	zoomedAreaUpdatedEventUnsubscriber := eventBus.Subscribe(zoomedareaupdatedevent.NewEventTopic(gameId, player.GetId()), func(event []byte) {
		handleZoomedAreaUpdatedEvent(conn, clientSession, gameId, player.GetId(), event)
	})
	defer zoomedAreaUpdatedEventUnsubscriber()

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

			compressionService := compressionservice.NewService()
			message, err := compressionService.Ungzip(compressedMessage)
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

		eventBus.Publish(
			removeplayerrequestedevent.NewEventTopic(gameId),
			removeplayerrequestedevent.NewEvent(player.GetId()),
		)

		return
	}
}

func sendJSONMessageToClient(conn *websocket.Conn, clientSession *clientSession, message any) {
	clientSession.socketSendMessageLock.Lock()
	defer clientSession.socketSendMessageLock.Unlock()

	messageJsonInBytes, _ := json.Marshal(message)

	compressionService := compressionservice.NewService()
	compressedMessage, err := compressionService.Gzip(messageJsonInBytes)
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

func emitGameInfoUpdatedEvent(conn *websocket.Conn, clientSession *clientSession, gameId uuid.UUID) {
	gameRoomRepository := gameroommemory.GetRepository()
	gameRoomService := gameroomservice.NewService(
		gameroomservice.Configuration{GameRoomRepository: gameRoomRepository},
	)
	unitMapSize, err := gameRoomService.GetUnitMapSize(gameId)
	if err != nil {
		emitErrorEvent(conn, clientSession, err)
		return
	}
	informationUpdatedEvent := constructInformationUpdatedEvent(unitMapSize)

	sendJSONMessageToClient(conn, clientSession, informationUpdatedEvent)
}

func handleAreaZoomedEvent(conn *websocket.Conn, clientSession *clientSession, gameId uuid.UUID, playerId uuid.UUID, event []byte) {
	var areaZoomedEvent areazoomedevent.Event
	json.Unmarshal(event, &areaZoomedEvent)

	sendJSONMessageToClient(conn, clientSession, constructAreaZoomedEvent(areaZoomedEvent.Payload.Area, areaZoomedEvent.Payload.UnitMap))
}

func handleZoomedAreaUpdatedEvent(conn *websocket.Conn, clientSession *clientSession, gameId uuid.UUID, playerId uuid.UUID, event []byte) {
	var zoomedAreaUpdatedEvent zoomedareaupdatedevent.Event
	json.Unmarshal(event, &zoomedAreaUpdatedEvent)

	sendJSONMessageToClient(conn, clientSession, constructZoomedAreaUpdatedEvent(zoomedAreaUpdatedEvent.Payload.Area, zoomedAreaUpdatedEvent.Payload.UnitMap))
}

func handleZoomAreaAction(conn *websocket.Conn, clientSession *clientSession, message []byte, gameId uuid.UUID, playerId uuid.UUID) {
	area, err := extractInformationFromZoomAreaAction(message)
	if err != nil {
		emitErrorEvent(conn, clientSession, err)
		return
	}

	eventBus := memoryeventbus.GetEventBus()
	eventBus.Publish(
		zoomarearequestedevent.NewEventTopic(gameId),
		zoomarearequestedevent.NewEvent(playerId, area),
	)
}

func handleReviveUnitsAction(conn *websocket.Conn, clientSession *clientSession, message []byte, gameId uuid.UUID) {
	coordinates, err := extractInformationFromReviveUnitsAction(message)
	if err != nil {
		emitErrorEvent(conn, clientSession, err)
		return
	}

	eventBus := memoryeventbus.GetEventBus()
	eventBus.Publish(
		reviveunitsrequestedevent.NewEventTopic(gameId),
		reviveunitsrequestedevent.NewEvent(coordinates),
	)
}
