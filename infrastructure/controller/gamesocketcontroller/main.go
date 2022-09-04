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

type clientSession struct {
	watchedArea           *areadto.Dto
	socketSendMessageLock sync.RWMutex
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

	clientSession := &clientSession{
		watchedArea:           nil,
		socketSendMessageLock: sync.RWMutex{},
	}

	emitGameInfoUpdatedEvent(conn, clientSession, gameId)

	gameUnitMapTickedEventBus := gameunitmaptickedeventbus.GetEventBus()
	gameUnitMapTickeddEventUnsubscriber := gameUnitMapTickedEventBus.Subscribe(gameId, func(updatedAt time.Time) {
		emitZoomedAreaUpdatedEvent(conn, clientSession, gameId, updatedAt)
	})
	defer gameUnitMapTickeddEventUnsubscriber()

	gameUnitsRevivedEventBus := gameunitsrevivedeventbus.GetEventBus()
	gameUnitsRevivedEventUnsubscriber := gameUnitsRevivedEventBus.Subscribe(gameId, func(coordinateDtos []coordinatedto.Dto, updatedAt time.Time) {
		emitUnitsRevivedEvent(conn, clientSession, gameId, coordinateDtos, updatedAt)
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

func emitUnitsRevivedEvent(conn *websocket.Conn, clientSession *clientSession, gameId uuid.UUID, coordinateDtos []coordinatedto.Dto, updatedAt time.Time) {
	if clientSession.watchedArea == nil {
		return
	}

	gameRoomRepository := gameroommemory.GetRepository()
	gameRoomService := gameroomservice.NewService(
		gameroomservice.Configuration{GameRoomRepository: gameRoomRepository},
	)
	coordinateDtosOfUnits, unitDtos, _ := gameRoomService.GetUnitsByCoordinatesInArea(gameId, coordinateDtos, *clientSession.watchedArea)
	gameUnitsRevivedEvent := constructUnitsRevivedEvent(coordinateDtosOfUnits, unitDtos, updatedAt)
	sendJSONMessageToClient(conn, clientSession, gameUnitsRevivedEvent)
}

func emitAreaZoomedEvent(conn *websocket.Conn, clientSession *clientSession, gameId uuid.UUID) {
	if clientSession.watchedArea == nil {
		return
	}

	gameRoomRepository := gameroommemory.GetRepository()
	gameRoomService := gameroomservice.NewService(
		gameroomservice.Configuration{GameRoomRepository: gameRoomRepository},
	)
	unitDtoMap, err := gameRoomService.GetUnitMapByArea(gameId, *clientSession.watchedArea)
	if err != nil {
		emitErrorEvent(conn, clientSession, err)
		return
	}

	areaZoomedEvent := constructAreaZoomedEvent(*clientSession.watchedArea, unitDtoMap)
	sendJSONMessageToClient(conn, clientSession, areaZoomedEvent)
}

func emitZoomedAreaUpdatedEvent(conn *websocket.Conn, clientSession *clientSession, gameId uuid.UUID, updatedAt time.Time) {
	if clientSession.watchedArea == nil {
		return
	}

	gameRoomRepository := gameroommemory.GetRepository()
	gameRoomService := gameroomservice.NewService(
		gameroomservice.Configuration{GameRoomRepository: gameRoomRepository},
	)
	unitDtoMap, err := gameRoomService.GetUnitMapByArea(gameId, *clientSession.watchedArea)
	if err != nil {
		emitErrorEvent(conn, clientSession, err)
		return
	}

	zoomedAreaUpdatedEvent := constructZoomedAreaUpdatedEvent(*clientSession.watchedArea, unitDtoMap, updatedAt)
	sendJSONMessageToClient(conn, clientSession, zoomedAreaUpdatedEvent)
}

func handleZoomAreaAction(conn *websocket.Conn, clientSession *clientSession, message []byte, gameId uuid.UUID) {
	zoomAreaAction, err := extractZoomAreaActionFromMessage(message)
	if err != nil {
		emitErrorEvent(conn, clientSession, err)
		return
	}

	clientSession.watchedArea = &zoomAreaAction.Payload.Area

	emitAreaZoomedEvent(conn, clientSession, gameId)
}

func handleReviveUnitsAction(conn *websocket.Conn, clientSession *clientSession, message []byte, gameId uuid.UUID) {
	reviveUnitsAction, err := extractReviveUnitsActionFromMessage(message)
	if err != nil {
		emitErrorEvent(conn, clientSession, err)
		return
	}

	gameRoomRepository := gameroommemory.GetRepository()
	gameUnitsRevivedEventBus := gameunitsrevivedeventbus.GetEventBus()
	gameRoomService := gameroomservice.NewService(
		gameroomservice.Configuration{GameRoomRepository: gameRoomRepository, UnitsRevivedEvent: gameUnitsRevivedEventBus},
	)
	err = gameRoomService.ReviveUnits(gameId, reviveUnitsAction.Payload.Coordinates)
	if err != nil {
		emitErrorEvent(conn, clientSession, err)
	}
}
