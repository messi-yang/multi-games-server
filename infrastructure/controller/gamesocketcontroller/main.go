package gamesocketcontroller

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/DumDumGeniuss/game-of-liberty-computer/application/dto/areadto"
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/dto/coordinatedto"
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/usecase/getgameroomusecase"
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/usecase/getunitmatrixusecase"
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/usecase/getunitsusecase"
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/usecase/reviveunitsusecase"
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/config"
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/eventbus/gamecomputedeventbus"
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/eventbus/unitsupdatedeventbus"
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/memory/gameroommemory"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
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

var playersCount int = 0

func Controller(c *gin.Context) {
	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.Error(err)
		return
	}
	defer conn.Close()
	closeConnFlag := make(chan bool)

	gameId := config.GetConfig().GetGameId()

	playersCount += 1
	session := &session{
		gameAreaToWatch: nil,
		socketLocker:    sync.RWMutex{},
	}

	emitGameInfoUpdatedEvent(conn, session, gameId)

	gameComputeEventBus := gamecomputedeventbus.GetGameComputedEventBus()
	gameComputedEventUnsubscriber := gameComputeEventBus.Subscribe(gameId, func() {
		emitAreaUpdatedEvent(conn, session, gameId)
	})
	defer gameComputedEventUnsubscriber()

	unitsUpdatedEventBus := unitsupdatedeventbus.GetUnitsUpdatedEventBus()
	unitsUpdatedEventUnsubscriber := unitsUpdatedEventBus.Subscribe(gameId, func(coordinateDTOs []coordinatedto.CoordinateDTO) {
		emitUnitsUpdatedEvent(conn, session, gameId, coordinateDTOs)
	})
	defer unitsUpdatedEventUnsubscriber()

	// conn.SetCloseHandler(func(code int, text string) error {
	// 	playersCount -= 1
	// 	return nil
	// })

	go func() {
		defer func() {
			closeConnFlag <- true
		}()

		for {
			_, message, err := conn.ReadMessage()
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
				handleWatchAreaAction(conn, session, message)
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

	conn.WriteJSON(message)
}

func emitErrorEvent(conn *websocket.Conn, session *session, err error) {
	errorEvent := constructErrorHappenedEvent(err.Error())

	sendJSONMessageToClient(conn, session, errorEvent)
}

func emitGameInfoUpdatedEvent(conn *websocket.Conn, session *session, gameId uuid.UUID) {
	gameRoomMemory := gameroommemory.GetGameRoomMemory()
	gameRoomDTO, _ := getgameroomusecase.New(gameRoomMemory).Execute(gameId)
	informationUpdatedEvent := constructInformationUpdatedEvent(gameRoomDTO.Game.MapSize, playersCount)

	sendJSONMessageToClient(conn, session, informationUpdatedEvent)
}

func emitUnitsUpdatedEvent(conn *websocket.Conn, session *session, gameId uuid.UUID, coordinateDTOs []coordinatedto.CoordinateDTO) {
	gameRoomMemory := gameroommemory.GetGameRoomMemory()
	unitDTOs, _ := getunitsusecase.New(gameRoomMemory).Execute(gameId, coordinateDTOs)
	unitsUpdatedEvent := constructUnitsUpdatedEvent(coordinateDTOs, unitDTOs)
	sendJSONMessageToClient(conn, session, unitsUpdatedEvent)
}

func emitAreaUpdatedEvent(conn *websocket.Conn, session *session, gameId uuid.UUID) {
	if session.gameAreaToWatch == nil {
		return
	}

	gameRoomMemory := gameroommemory.GetGameRoomMemory()
	unitDTOMatrix, err := getunitmatrixusecase.New(gameRoomMemory).Execute(gameId, *session.gameAreaToWatch)
	if err != nil {
		emitErrorEvent(conn, session, err)
		return
	}

	areaUpdatedEvent := constructAreaUpdatedEvent(*session.gameAreaToWatch, unitDTOMatrix)
	sendJSONMessageToClient(conn, session, areaUpdatedEvent)
}

func handleWatchAreaAction(conn *websocket.Conn, session *session, message []byte) {
	watchAreaAction, err := extractWatchAreaActionFromMessage(message)
	if err != nil {
		emitErrorEvent(conn, session, err)
	}
	session.gameAreaToWatch = &watchAreaAction.Payload.Area
}

func handleReviveUnitsAction(conn *websocket.Conn, session *session, message []byte, gameId uuid.UUID) {
	reviveUnitsAction, err := extractReviveUnitsActionFromMessage(message)
	if err != nil {
		emitErrorEvent(conn, session, err)
	}

	gameRoomMemory := gameroommemory.GetGameRoomMemory()
	unitsUpdatedEventBus := unitsupdatedeventbus.GetUnitsUpdatedEventBus()
	reviveunitsusecase.New(gameRoomMemory, unitsUpdatedEventBus).Execute(gameId, reviveUnitsAction.Payload.Coordinates)
}
