package gamesocketcontroller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/areadto"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/coordinatedto"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/service/compressionservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/service/gameroomservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/infrastructure/config"
	"github.com/dum-dum-genius/game-of-liberty-computer/infrastructure/eventbus/gameunitmapupdatedeventbus"
	"github.com/dum-dum-genius/game-of-liberty-computer/infrastructure/eventbus/gameunitsupdatedeventbus"
	"github.com/dum-dum-genius/game-of-liberty-computer/infrastructure/memory/gameroommemory"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/ztrue/tracerr"
)

type session struct {
	gameAreaToWatch *areadto.AreaDTO
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

	gameUnitMapUpdatedEventBus := gameunitmapupdatedeventbus.GetGameUnitMapUpdatedEventBus()
	gameUnitMapUpdateddEventUnsubscriber := gameUnitMapUpdatedEventBus.Subscribe(gameId, func() {
		emitUnitMapFetched(conn, session, gameId)
	})
	defer gameUnitMapUpdateddEventUnsubscriber()

	gameUnitsUpdatedEventBus := gameunitsupdatedeventbus.GetUnitsUpdatedEventBus()
	gameUnitsUpdatedEventUnsubscriber := gameUnitsUpdatedEventBus.Subscribe(gameId, func(coordinateDTOs []coordinatedto.CoordinateDTO) {
		emitUnitsUpdatedEvent(conn, session, gameId, coordinateDTOs)
	})
	defer gameUnitsUpdatedEventUnsubscriber()

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

			compressionService := compressionservice.NewCompressionService()
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

	compressionService := compressionservice.NewCompressionService()
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
	gameRoomService := gameroomservice.NewGameRoomService(
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

func emitUnitsUpdatedEvent(conn *websocket.Conn, session *session, gameId uuid.UUID, coordinateDTOs []coordinatedto.CoordinateDTO) {
	if session.gameAreaToWatch == nil {
		return
	}

	gameRoomMemory := gameroommemory.GetGameRoomMemory()
	gameRoomService := gameroomservice.NewGameRoomService(
		gameroomservice.Configuration{GameRoomRepository: gameRoomMemory},
	)
	coordinateDTOsOfUnits, unitDTOs, _ := gameRoomService.GetUnitsByCoordinatesInArea(gameId, coordinateDTOs, *session.gameAreaToWatch)
	gameUnitsUpdatedEvent := constructUnitsUpdatedEvent(coordinateDTOsOfUnits, unitDTOs)
	sendJSONMessageToClient(conn, session, gameUnitsUpdatedEvent)
}

func emitUnitMapFetched(conn *websocket.Conn, session *session, gameId uuid.UUID) {
	if session.gameAreaToWatch == nil {
		return
	}

	gameRoomMemory := gameroommemory.GetGameRoomMemory()
	gameRoomService := gameroomservice.NewGameRoomService(
		gameroomservice.Configuration{GameRoomRepository: gameRoomMemory},
	)
	unitDTOMap, err := gameRoomService.GetUnitMapByArea(gameId, *session.gameAreaToWatch)
	if err != nil {
		emitErrorEvent(conn, session, err)
		return
	}

	unitMapFetchedEvent := constructUnitMapFetched(*session.gameAreaToWatch, unitDTOMap)
	sendJSONMessageToClient(conn, session, unitMapFetchedEvent)
}

func handleWatchAreaAction(conn *websocket.Conn, session *session, message []byte, gameId uuid.UUID) {
	watchAreaAction, err := extractWatchAreaActionFromMessage(message)
	if err != nil {
		emitErrorEvent(conn, session, err)
		return
	}
	session.gameAreaToWatch = &watchAreaAction.Payload.Area

	emitUnitMapFetched(conn, session, gameId)
}

func handleReviveUnitsAction(conn *websocket.Conn, session *session, message []byte, gameId uuid.UUID) {
	reviveUnitsAction, err := extractReviveUnitsActionFromMessage(message)
	if err != nil {
		emitErrorEvent(conn, session, err)
		return
	}

	gameRoomMemory := gameroommemory.GetGameRoomMemory()
	gameUnitsUpdatedEventBus := gameunitsupdatedeventbus.GetUnitsUpdatedEventBus()
	gameRoomService := gameroomservice.NewGameRoomService(
		gameroomservice.Configuration{GameRoomRepository: gameRoomMemory, UnitsUpdatedEvent: gameUnitsUpdatedEventBus},
	)
	err = gameRoomService.ReviveUnits(gameId, reviveUnitsAction.Payload.Coordinates)
	if err != nil {
		emitErrorEvent(conn, session, err)
	}
}
