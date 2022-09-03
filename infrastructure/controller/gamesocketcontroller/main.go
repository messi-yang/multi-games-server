package gamesocketcontroller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/areadto"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/coordinatedto"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/service/compressionservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/service/gameroomservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/infrastructure/config"
	"github.com/dum-dum-genius/game-of-liberty-computer/infrastructure/eventbus/gameunitmaptickedeventbus"
	"github.com/dum-dum-genius/game-of-liberty-computer/infrastructure/eventbus/gameunitsrevivedeventbus"
	"github.com/dum-dum-genius/game-of-liberty-computer/infrastructure/memory/gameroommemory"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/ztrue/tracerr"
)

type session struct {
	gameAreaToWatch *areadto.DTO
	socketLocker    sync.RWMutex
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Controller(c *gin.Context) {
	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.Error(err)
		return
	}
	defer conn.Close()
	closeConnFlag := make(chan bool)

	gameId := config.GetConfig().GetGameId()

	session := &session{
		gameAreaToWatch: nil,
		socketLocker:    sync.RWMutex{},
	}

	emitGameInfoUpdatedEvent(conn, session, gameId)

	gameUnitMapTickedEventBus := gameunitmaptickedeventbus.GetGameUnitMapTickedEventBus()
	gameUnitMapTickeddEventUnsubscriber := gameUnitMapTickedEventBus.Subscribe(gameId, func(updatedAt time.Time) {
		emitUnitMapTickedEvent(conn, session, gameId, updatedAt)
	})
	defer gameUnitMapTickeddEventUnsubscriber()

	gameUnitsRevivedEventBus := gameunitsrevivedeventbus.GetUnitsRevivedEventBus()
	gameUnitsRevivedEventUnsubscriber := gameUnitsRevivedEventBus.Subscribe(gameId, func(coordinateDTOs []coordinatedto.DTO, updatedAt time.Time) {
		emitUnitsRevivedEvent(conn, session, gameId, coordinateDTOs, updatedAt)
	})
	defer gameUnitsRevivedEventUnsubscriber()

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
				emitErrorEvent(conn, session, err)
				break
			}

			compressionService := compressionservice.NewService()
			message, err := compressionService.Ungzip(compressedMessage)
			if err != nil {
				emitErrorEvent(conn, session, err)
				break
			}

			actionType, err := getActionTypeFromMessage(message)
			if err != nil {
				emitErrorEvent(conn, session, err)
				break
			}

			switch *actionType {
			case watchAreaActionType:
				handleWatchAreaAction(conn, session, message, gameId)
			case reviveUnitsActionType:
				handleReviveUnitsAction(conn, session, message, gameId)
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

func sendJSONMessageToClient(conn *websocket.Conn, session *session, message any) {
	session.socketLocker.Lock()
	defer session.socketLocker.Unlock()

	messageJsonInBytes, _ := json.Marshal(message)

	compressionService := compressionservice.NewService()
	compressedMessage, err := compressionService.Gzip(messageJsonInBytes)
	if err != nil {
		emitErrorEvent(conn, session, err)
		return
	}

	conn.WriteMessage(2, compressedMessage)
}

func emitErrorEvent(conn *websocket.Conn, session *session, err error) {
	errorEvent := constructErrorHappenedEvent(err.Error())

	tracerr.Print(tracerr.Wrap(err))

	sendJSONMessageToClient(conn, session, errorEvent)
}

func emitGameInfoUpdatedEvent(conn *websocket.Conn, session *session, gameId uuid.UUID) {
	gameRoomMemory := gameroommemory.GetGameRoomMemory()
	gameRoomService := gameroomservice.NewService(
		gameroomservice.Configuration{GameRoomRepository: gameRoomMemory},
	)
	unitMapSize, err := gameRoomService.GetUnitMapSize(gameId)
	if err != nil {
		emitErrorEvent(conn, session, err)
		return
	}
	informationUpdatedEvent := constructInformationUpdatedEvent(unitMapSize)

	sendJSONMessageToClient(conn, session, informationUpdatedEvent)
}

func emitUnitsRevivedEvent(conn *websocket.Conn, session *session, gameId uuid.UUID, coordinateDTOs []coordinatedto.DTO, updatedAt time.Time) {
	if session.gameAreaToWatch == nil {
		return
	}

	gameRoomMemory := gameroommemory.GetGameRoomMemory()
	gameRoomService := gameroomservice.NewService(
		gameroomservice.Configuration{GameRoomRepository: gameRoomMemory},
	)
	coordinateDTOsOfUnits, unitDTOs, _ := gameRoomService.GetUnitsByCoordinatesInArea(gameId, coordinateDTOs, *session.gameAreaToWatch)
	gameUnitsRevivedEvent := constructUnitsRevivedEvent(coordinateDTOsOfUnits, unitDTOs, updatedAt)
	sendJSONMessageToClient(conn, session, gameUnitsRevivedEvent)
}

func emitUnitMapReceivedEvent(conn *websocket.Conn, session *session, gameId uuid.UUID) {
	if session.gameAreaToWatch == nil {
		return
	}

	gameRoomMemory := gameroommemory.GetGameRoomMemory()
	gameRoomService := gameroomservice.NewService(
		gameroomservice.Configuration{GameRoomRepository: gameRoomMemory},
	)
	unitDTOMap, receivedAt, err := gameRoomService.GetUnitMapByArea(gameId, *session.gameAreaToWatch)
	if err != nil {
		emitErrorEvent(conn, session, err)
		return
	}

	unitMapReceivedEvent := constructUnitMapReceived(*session.gameAreaToWatch, unitDTOMap, receivedAt)
	sendJSONMessageToClient(conn, session, unitMapReceivedEvent)
}

func emitUnitMapTickedEvent(conn *websocket.Conn, session *session, gameId uuid.UUID, updatedAt time.Time) {
	if session.gameAreaToWatch == nil {
		return
	}

	gameRoomMemory := gameroommemory.GetGameRoomMemory()
	gameRoomService := gameroomservice.NewService(
		gameroomservice.Configuration{GameRoomRepository: gameRoomMemory},
	)
	unitDTOMap, _, err := gameRoomService.GetUnitMapByArea(gameId, *session.gameAreaToWatch)
	if err != nil {
		emitErrorEvent(conn, session, err)
		return
	}

	unitMapTickedEvent := constructUnitMapTicked(*session.gameAreaToWatch, unitDTOMap, updatedAt)
	sendJSONMessageToClient(conn, session, unitMapTickedEvent)
}

func handleWatchAreaAction(conn *websocket.Conn, session *session, message []byte, gameId uuid.UUID) {
	watchAreaAction, err := extractWatchAreaActionFromMessage(message)
	if err != nil {
		emitErrorEvent(conn, session, err)
		return
	}

	session.gameAreaToWatch = &watchAreaAction.Payload.Area

	emitUnitMapReceivedEvent(conn, session, gameId)
}

func handleReviveUnitsAction(conn *websocket.Conn, session *session, message []byte, gameId uuid.UUID) {
	reviveUnitsAction, err := extractReviveUnitsActionFromMessage(message)
	if err != nil {
		emitErrorEvent(conn, session, err)
		return
	}

	gameRoomMemory := gameroommemory.GetGameRoomMemory()
	gameUnitsRevivedEventBus := gameunitsrevivedeventbus.GetUnitsRevivedEventBus()
	gameRoomService := gameroomservice.NewService(
		gameroomservice.Configuration{GameRoomRepository: gameRoomMemory, UnitsRevivedEvent: gameUnitsRevivedEventBus},
	)
	err = gameRoomService.ReviveUnits(gameId, reviveUnitsAction.Payload.Coordinates)
	if err != nil {
		emitErrorEvent(conn, session, err)
	}
}
