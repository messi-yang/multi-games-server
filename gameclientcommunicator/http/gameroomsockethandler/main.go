package gameroomsockethandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/gameclientcommunicator/application/service/compressionservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/application/service/gameroomservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/infrastructure/memory/gameroommemory"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/config"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/eventbus"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/eventbus/gameunitmaptickedeventbus"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/event/reviveunitsrequestedevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/event/unitsrevivedevent"
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

	gameId := config.GetConfig().GetGameId()

	clientSession := &clientSession{
		watchedArea:           nil,
		socketSendMessageLock: sync.RWMutex{},
	}

	emitGameInfoUpdatedEvent(conn, clientSession, gameId)
	eventBus := eventbus.GetEventBus()

	gameUnitMapTickedEventBus := gameunitmaptickedeventbus.GetEventBus()
	gameUnitMapTickeddEventUnsubscriber := gameUnitMapTickedEventBus.Subscribe(gameId, func() {
		emitZoomedAreaUpdatedEvent(conn, clientSession, gameId)
	})
	defer gameUnitMapTickeddEventUnsubscriber()

	unitsRevivedEventInsubscriber := eventBus.Subscribe(unitsrevivedevent.NewEventTopic(gameId), func(event []byte) {
		handleUnitsRevivedEvent(conn, clientSession, gameId, event)
	})
	defer unitsRevivedEventInsubscriber()

	// conn.SetCloseHandler(func(code int, text string) error {
	// 	return nil
	// })

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
				handleZoomAreaAction(conn, clientSession, message, gameId)
			case reviveUnitsActionType:
				handleReviveUnitsAction(conn, clientSession, message, gameId)
			default:
			}
		}
	}()

	for {
		<-closeConnFlag
		fmt.Println("Player left")
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

func handleUnitsRevivedEvent(conn *websocket.Conn, clientSession *clientSession, gameId uuid.UUID, event []byte) {
	if clientSession.watchedArea == nil {
		return
	}

	var unitsRevivedEvent unitsrevivedevent.Event
	json.Unmarshal(event, &unitsRevivedEvent)

	coordinates := make([]valueobject.Coordinate, 0)

	for _, coordInPayload := range unitsRevivedEvent.Payload.Coordinates {
		coordinate, _ := valueobject.NewCoordinate(coordInPayload.X, coordInPayload.Y)
		coordinates = append(coordinates, coordinate)
	}

	gameRoomRepository := gameroommemory.GetRepository()
	gameRoomService := gameroomservice.NewService(
		gameroomservice.Configuration{GameRoomRepository: gameRoomRepository},
	)

	includes, err := gameRoomService.AreaIncludesAnyCoordinates(*clientSession.watchedArea, coordinates)
	if err != nil {
		emitErrorEvent(conn, clientSession, err)
		return
	}

	if !includes {
		return
	}

	unitMap, err := gameRoomService.GetUnitMapByArea(gameId, *clientSession.watchedArea)
	if err != nil {
		emitErrorEvent(conn, clientSession, err)
		return
	}
	zoomedAreaUpdatedEvent := constructZoomedAreaUpdatedEvent(*clientSession.watchedArea, unitMap)
	sendJSONMessageToClient(conn, clientSession, zoomedAreaUpdatedEvent)
}

func emitAreaZoomedEvent(conn *websocket.Conn, clientSession *clientSession, gameId uuid.UUID) {
	if clientSession.watchedArea == nil {
		return
	}

	gameRoomRepository := gameroommemory.GetRepository()
	gameRoomService := gameroomservice.NewService(
		gameroomservice.Configuration{GameRoomRepository: gameRoomRepository},
	)
	unitMap, err := gameRoomService.GetUnitMapByArea(gameId, *clientSession.watchedArea)
	if err != nil {
		emitErrorEvent(conn, clientSession, err)
		return
	}

	areaZoomedEvent := constructAreaZoomedEvent(*clientSession.watchedArea, unitMap)
	sendJSONMessageToClient(conn, clientSession, areaZoomedEvent)
}

func emitZoomedAreaUpdatedEvent(conn *websocket.Conn, clientSession *clientSession, gameId uuid.UUID) {
	if clientSession.watchedArea == nil {
		return
	}

	gameRoomRepository := gameroommemory.GetRepository()
	gameRoomService := gameroomservice.NewService(
		gameroomservice.Configuration{GameRoomRepository: gameRoomRepository},
	)

	unitMap, err := gameRoomService.GetUnitMapByArea(gameId, *clientSession.watchedArea)
	if err != nil {
		emitErrorEvent(conn, clientSession, err)
		return
	}

	zoomedAreaUpdatedEvent := constructZoomedAreaUpdatedEvent(*clientSession.watchedArea, unitMap)
	sendJSONMessageToClient(conn, clientSession, zoomedAreaUpdatedEvent)
}

func handleZoomAreaAction(conn *websocket.Conn, clientSession *clientSession, message []byte, gameId uuid.UUID) {
	area, err := extractInformationFromZoomAreaAction(message)
	if err != nil {
		emitErrorEvent(conn, clientSession, err)
		return
	}

	clientSession.watchedArea = &area

	emitAreaZoomedEvent(conn, clientSession, gameId)
}

func handleReviveUnitsAction(conn *websocket.Conn, clientSession *clientSession, message []byte, gameId uuid.UUID) {
	coordinates, err := extractInformationFromReviveUnitsAction(message)
	if err != nil {
		emitErrorEvent(conn, clientSession, err)
		return
	}

	eventBus := eventbus.GetEventBus()
	eventBus.Publish(
		reviveunitsrequestedevent.NewEventTopic(),
		reviveunitsrequestedevent.NewEvent(gameId, coordinates),
	)
}
