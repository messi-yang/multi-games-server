package gamesocketcontroller

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/areadto"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/coordinatedto"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/usecase/getcoordinatesinareausecase"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/usecase/getunitmapsizeusecase"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/usecase/getunitmapusecase"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/usecase/getunitsusecase"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/usecase/reviveunitsusecase"
	"github.com/dum-dum-genius/game-of-liberty-computer/infrastructure/config"
	"github.com/dum-dum-genius/game-of-liberty-computer/infrastructure/eventbus/coordinatesupdatedeventbus"
	"github.com/dum-dum-genius/game-of-liberty-computer/infrastructure/eventbus/gamecomputedeventbus"
	"github.com/dum-dum-genius/game-of-liberty-computer/infrastructure/memory/gameroommemory"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func ungzipData(compressedData []byte) (data []byte) {
	gunzip, _ := gzip.NewReader(bytes.NewBuffer(compressedData))
	defer gunzip.Close()
	data, _ = ioutil.ReadAll(gunzip)
	return data
}

func gzipData(data []byte) (compressedData []byte) {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	gz.Write(data)
	gz.Flush()
	gz.Close()

	return b.Bytes()
}

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

	gameComputeEventBus := gamecomputedeventbus.GetGameComputedEventBus()
	gameComputedEventUnsubscriber := gameComputeEventBus.Subscribe(gameId, func() {
		emitAreaUpdatedEvent(conn, session, gameId)
	})
	defer gameComputedEventUnsubscriber()

	coordinatesUpdatedEventBus := coordinatesupdatedeventbus.GetCoordinatesUpdatedEventBus()
	coordinatesUpdatedEventUnsubscriber := coordinatesUpdatedEventBus.Subscribe(gameId, func(coordinateDTOs []coordinatedto.CoordinateDTO) {
		emitCoordinatesUpdatedEvent(conn, session, gameId, coordinateDTOs)
	})
	defer coordinatesUpdatedEventUnsubscriber()

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

			message := ungzipData(compressedMessage)
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

	conn.WriteMessage(2, gzipData(messageJsonInBytes))
}

func emitErrorEvent(conn *websocket.Conn, session *session, err error) {
	errorEvent := constructErrorHappenedEvent(err.Error())

	sendJSONMessageToClient(conn, session, errorEvent)
}

func emitGameInfoUpdatedEvent(conn *websocket.Conn, session *session, gameId uuid.UUID) {
	gameRoomMemory := gameroommemory.GetGameRoomMemory()
	unitMapSize, _ := getunitmapsizeusecase.New(gameRoomMemory).Execute(gameId)
	informationUpdatedEvent := constructInformationUpdatedEvent(unitMapSize)

	sendJSONMessageToClient(conn, session, informationUpdatedEvent)
}

func emitCoordinatesUpdatedEvent(conn *websocket.Conn, session *session, gameId uuid.UUID, coordinateDTOs []coordinatedto.CoordinateDTO) {
	if session.gameAreaToWatch == nil {
		return
	}

	gameRoomMemory := gameroommemory.GetGameRoomMemory()
	coordinateDTOsInAreaDTO, err := getcoordinatesinareausecase.New(gameRoomMemory).Execute(gameId, coordinateDTOs, *session.gameAreaToWatch)
	if err != nil {
		emitErrorEvent(conn, session, err)
		return
	}
	unitDTOs, _ := getunitsusecase.New(gameRoomMemory).Execute(gameId, coordinateDTOsInAreaDTO)
	coordinatesUpdatedEvent := constructCoordinatesUpdatedEvent(coordinateDTOsInAreaDTO, unitDTOs)
	sendJSONMessageToClient(conn, session, coordinatesUpdatedEvent)
}

func emitAreaUpdatedEvent(conn *websocket.Conn, session *session, gameId uuid.UUID) {
	if session.gameAreaToWatch == nil {
		return
	}

	gameRoomMemory := gameroommemory.GetGameRoomMemory()
	unitDTOMap, err := getunitmapusecase.New(gameRoomMemory).Execute(gameId, *session.gameAreaToWatch)
	if err != nil {
		emitErrorEvent(conn, session, err)
		return
	}

	areaUpdatedEvent := constructAreaUpdatedEvent(*session.gameAreaToWatch, unitDTOMap)
	sendJSONMessageToClient(conn, session, areaUpdatedEvent)
}

func handleWatchAreaAction(conn *websocket.Conn, session *session, message []byte, gameId uuid.UUID) {
	watchAreaAction, err := extractWatchAreaActionFromMessage(message)
	if err != nil {
		emitErrorEvent(conn, session, err)
	}
	session.gameAreaToWatch = &watchAreaAction.Payload.Area

	emitAreaUpdatedEvent(conn, session, gameId)
}

func handleReviveUnitsAction(conn *websocket.Conn, session *session, message []byte, gameId uuid.UUID) {
	reviveUnitsAction, err := extractReviveUnitsActionFromMessage(message)
	if err != nil {
		emitErrorEvent(conn, session, err)
	}

	gameRoomMemory := gameroommemory.GetGameRoomMemory()
	coordinatesUpdatedEventBus := coordinatesupdatedeventbus.GetCoordinatesUpdatedEventBus()
	err = reviveunitsusecase.New(gameRoomMemory, coordinatesUpdatedEventBus).Execute(gameId, reviveUnitsAction.Payload.Coordinates)
	if err != nil {
		emitErrorEvent(conn, session, err)
	}
}
