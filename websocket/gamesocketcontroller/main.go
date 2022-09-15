package gamesocketcontroller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/application/service/compressionservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/service/gameroomservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/config"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/dto/areadto"
	"github.com/dum-dum-genius/game-of-liberty-computer/dto/coordinatedto"
	"github.com/dum-dum-genius/game-of-liberty-computer/dto/mapsizedto"
	"github.com/dum-dum-genius/game-of-liberty-computer/dto/unitmapdto"
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
	gameUnitMapTickeddEventUnsubscriber := gameUnitMapTickedEventBus.Subscribe(gameId, func() {
		emitZoomedAreaUpdatedEvent(conn, clientSession, gameId)
	})
	defer gameUnitMapTickeddEventUnsubscriber()

	gameUnitsRevivedEventBus := gameunitsrevivedeventbus.GetEventBus()
	gameUnitsRevivedEventUnsubscriber := gameUnitsRevivedEventBus.Subscribe(gameId, func(coordinates []valueobject.Coordinate) {
		coordinateDTOs := coordinatedto.ToDtoList(coordinates)
		handleUnitsRevivedEvent(conn, clientSession, gameId, coordinateDTOs)
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
	unitMapSizeDTO := mapsizedto.ToDto(unitMapSize)
	informationUpdatedEvent := constructInformationUpdatedEvent(unitMapSizeDTO)

	sendJSONMessageToClient(conn, clientSession, informationUpdatedEvent)
}

func handleUnitsRevivedEvent(conn *websocket.Conn, clientSession *clientSession, gameId uuid.UUID, coordinateDtos []coordinatedto.Dto) {
	if clientSession.watchedArea == nil {
		return
	}

	gameRoomRepository := gameroommemory.GetRepository()
	gameRoomService := gameroomservice.NewService(
		gameroomservice.Configuration{GameRoomRepository: gameRoomRepository},
	)
	coordinates, err := coordinatedto.FromDtoList(coordinateDtos)
	if err != nil {
		emitErrorEvent(conn, clientSession, err)
		return
	}

	area, err := areadto.FromDto(*clientSession.watchedArea)
	if err != nil {
		emitErrorEvent(conn, clientSession, err)
		return
	}
	includes, err := gameRoomService.AreaIncludesAnyCoordinates(area, coordinates)
	if err != nil {
		emitErrorEvent(conn, clientSession, err)
		return
	}

	if !includes {
		return
	}

	unitMap, err := gameRoomService.GetUnitMapByArea(gameId, area)
	if err != nil {
		emitErrorEvent(conn, clientSession, err)
		return
	}
	unitMapDTO := unitmapdto.ToDto(&unitMap)
	zoomedAreaUpdatedEvent := constructZoomedAreaUpdatedEvent(*clientSession.watchedArea, unitMapDTO)
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
	area, err := areadto.FromDto(*clientSession.watchedArea)
	if err != nil {
		emitErrorEvent(conn, clientSession, err)
		return
	}
	unitMap, err := gameRoomService.GetUnitMapByArea(gameId, area)
	if err != nil {
		emitErrorEvent(conn, clientSession, err)
		return
	}

	unitMapDTO := unitmapdto.ToDto(&unitMap)
	areaZoomedEvent := constructAreaZoomedEvent(*clientSession.watchedArea, unitMapDTO)
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

	area, err := areadto.FromDto(*clientSession.watchedArea)
	if err != nil {
		emitErrorEvent(conn, clientSession, err)
		return
	}
	unitMap, err := gameRoomService.GetUnitMapByArea(gameId, area)
	if err != nil {
		emitErrorEvent(conn, clientSession, err)
		return
	}

	unitMapDTO := unitmapdto.ToDto(&unitMap)
	zoomedAreaUpdatedEvent := constructZoomedAreaUpdatedEvent(*clientSession.watchedArea, unitMapDTO)
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
	coordinates, err := coordinatedto.FromDtoList(reviveUnitsAction.Payload.Coordinates)
	if err != nil {
		emitErrorEvent(conn, clientSession, err)
	}
	err = gameRoomService.ReviveUnits(gameId, coordinates)
	if err != nil {
		emitErrorEvent(conn, clientSession, err)
	}
}
