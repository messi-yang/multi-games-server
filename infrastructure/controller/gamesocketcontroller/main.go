package gamesocketcontroller

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/DumDumGeniuss/game-of-liberty-computer/application/usecase/getgameroomusecase"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/repository/gameroomrepository"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/service/gameroomservice"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/valueobject"
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/config"
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/dto"
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/eventbus/gamecomputedeventbus"
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/eventbus/gameunitsupdatedeventbus"
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/memory/gameroommemory"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

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
	gameRoomMemory := gameroommemory.GetGameRoomMemory()
	gameRoomService := gameroomservice.NewGameRoomService(gameRoomMemory)

	playersCount += 1
	session := &session{
		gameAreaToWatch: nil,
		socketLocker:    sync.RWMutex{},
	}

	emitGameInfoUpdatedEvent(conn, session, gameId, gameRoomMemory)

	gameComputeEventBus := gamecomputedeventbus.GetGameComputedEventBus()
	handleGameComputedEvent := func() {
		emitAreaUpdatedEvent(conn, session, gameId, gameRoomMemory)
	}
	gameComputeEventBus.Subscribe(gameId, handleGameComputedEvent)
	defer gameComputeEventBus.Unsubscribe(gameId, handleGameComputedEvent)

	gameUnitsUpdatedEvent := gameunitsupdatedeventbus.GetGameUnitsUpdatedEventBus()
	handleGameUnitsUpdatedEvent := func(coordinates []valueobject.Coordinate) {
		gameRoom, _ := getgameroomusecase.NewUseCase(gameRoomMemory).Execute(gameId)
		unitsUpdatedEventPayloadItems := []unitsUpdatedEventPayloadItem{}
		for _, coord := range coordinates {
			unit := gameRoom.GetGameUnit(coord)

			unitsUpdatedEventPayloadItems = append(
				unitsUpdatedEventPayloadItems,
				unitsUpdatedEventPayloadItem{
					Coordinate: dto.CoordinateDTO{X: coord.GetX(), Y: coord.GetY()},
					Unit:       dto.GameUnitDTO{Alive: unit.GetAlive(), Age: unit.GetAge()},
				},
			)
		}
		emitUnitsUpdatedEvent(conn, session, &unitsUpdatedEventPayloadItems)
	}
	gameUnitsUpdatedEvent.Subscribe(gameId, handleGameUnitsUpdatedEvent)
	defer gameUnitsUpdatedEvent.Unsubscribe(gameId, handleGameUnitsUpdatedEvent)

	conn.SetCloseHandler(func(code int, text string) error {
		playersCount -= 1
		return nil
	})

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
			}

			switch *actionType {
			case watchAreaActionType:
				watchAreaAction, err := extractWatchAreaActionFromMessage(message)
				if err != nil {
					emitErrorEvent(conn, session, err)
				}
				area := valueobject.NewArea(
					valueobject.NewCoordinate(
						watchAreaAction.Payload.Area.From.X,
						watchAreaAction.Payload.Area.From.Y,
					),
					valueobject.NewCoordinate(
						watchAreaAction.Payload.Area.To.X,
						watchAreaAction.Payload.Area.To.Y,
					),
				)
				session.gameAreaToWatch = &area
			case reviveUnitsActionType:
				reviveUnitsAction, err := extractReviveUnitsActionFromMessage(message)
				if err != nil {
					emitErrorEvent(conn, session, err)
				}

				coordinates := make([]valueobject.Coordinate, 0)

				for _, coord := range reviveUnitsAction.Payload.Coordinates {
					coordinate := valueobject.NewCoordinate(coord.X, coord.Y)
					coordinates = append(coordinates, coordinate)
					gameRoomService.ReviveGameUnit(gameId, coordinate)
				}

				gameUnitsUpdatedEvent.Publish(gameId, coordinates)
			default:
			}
		}
	}()

	for {
		select {
		case <-closeConnFlag:
			fmt.Println("Player left")
			return
		}
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

func emitGameInfoUpdatedEvent(conn *websocket.Conn, session *session, gameId uuid.UUID, gameRoomRepository gameroomrepository.GameRoomRepository) {
	gameRoom, _ := getgameroomusecase.NewUseCase(gameRoomRepository).Execute(gameId)
	gameMapSize := gameRoom.GetGameMapSize()
	informationUpdatedEvent := constructInformationUpdatedEvent(&gameMapSize, playersCount)

	sendJSONMessageToClient(conn, session, informationUpdatedEvent)
}

func emitUnitsUpdatedEvent(conn *websocket.Conn, session *session, updateUnitItems *[]unitsUpdatedEventPayloadItem) {
	unitsUpdatedEvent := constructUnitsUpdatedEvent(updateUnitItems)

	sendJSONMessageToClient(conn, session, unitsUpdatedEvent)
}

func emitAreaUpdatedEvent(conn *websocket.Conn, session *session, gameId uuid.UUID, gameRoomRepository gameroomrepository.GameRoomRepository) {
	if session.gameAreaToWatch == nil {
		return
	}

	gameRoom, _ := getgameroomusecase.NewUseCase(gameRoomRepository).Execute(gameId)
	gameUnits, err := gameRoom.GetGameUnitMatrixWithArea(*session.gameAreaToWatch)
	if err != nil {
		emitErrorEvent(conn, session, err)
		return
	}

	gameUnitsDTO := make([][]dto.GameUnitDTO, 0)

	for i := 0; i < len(gameUnits); i += 1 {
		gameUnitsDTO = append(gameUnitsDTO, make([]dto.GameUnitDTO, 0))
		for j := 0; j < len(gameUnits[i]); j += 1 {
			gameUnitsDTO[i] = append(gameUnitsDTO[i], dto.GameUnitDTO{
				Alive: gameUnits[i][j].GetAlive(),
				Age:   gameUnits[i][j].GetAge(),
			})
		}
	}

	areaUpdatedEvent := constructAreaUpdatedEvent(session.gameAreaToWatch, &gameUnitsDTO)

	sendJSONMessageToClient(conn, session, areaUpdatedEvent)
}
