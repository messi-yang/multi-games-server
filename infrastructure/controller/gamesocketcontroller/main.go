package gamesocketcontroller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/service/gameroomservice"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/valueobject"
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/config"
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/dto"
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/memory/gameroommemory"
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/service/messageservice"
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/service/messageservicetopic"
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
	messageService := messageservice.GetMessageService()
	gameRoomMemoryRepository := gameroommemory.NewGameRoomMemoryRepository()
	gameRoomService := gameroomservice.NewGameRoomService(gameRoomMemoryRepository)

	playersCount += 1
	session := &session{
		gameAreaToWatch: nil,
		socketLocker:    sync.RWMutex{},
	}

	emitGameInfoUpdatedEvent(conn, session, gameId, gameRoomService)
	messageService.Publish(messageservicetopic.GamePlayerJoinedMessageTopic, nil)

	areaUpdatedSubscriptionToken := messageService.Subscribe(messageservicetopic.GameRoomJobTickedMessageTopic, func(_ []byte) {
		emitAreaUpdatedEvent(conn, session, gameId, gameRoomService)
	})
	defer messageService.Unsubscribe(messageservicetopic.GameRoomJobTickedMessageTopic, areaUpdatedSubscriptionToken)

	unitsUpdatedSubscriptionToken := messageService.Subscribe(messageservicetopic.GameUnitsUpdatedMessageTopic, func(message []byte) {
		var messagePayload messageservicetopic.GameUnitsUpdatedMessageTopicPayload
		json.Unmarshal(message, &messagePayload)

		unitsUpdatedEventPayloadItems := []unitsUpdatedEventPayloadItem{}
		for _, messagePayloadUnit := range messagePayload {

			unitsUpdatedEventPayloadItems = append(
				unitsUpdatedEventPayloadItems,
				unitsUpdatedEventPayloadItem{
					Coordinate: messagePayloadUnit.Coordinate,
					Unit:       messagePayloadUnit.Unit,
				},
			)
		}
		emitUnitsUpdatedEvent(conn, session, &unitsUpdatedEventPayloadItems)
	})
	defer messageService.Unsubscribe(messageservicetopic.GameUnitsUpdatedMessageTopic, unitsUpdatedSubscriptionToken)

	playerJoinedSubscriptionToken := messageService.Subscribe(messageservicetopic.GamePlayerJoinedMessageTopic, func(_ []byte) {
		emitPlayerJoinedEvent(conn, session)
	})
	defer messageService.Unsubscribe(messageservicetopic.GamePlayerJoinedMessageTopic, playerJoinedSubscriptionToken)

	playerLeftSubscriptionToken := messageService.Subscribe(messageservicetopic.GamePlayerLeftMessageTopic, func(_ []byte) {
		emitPlayerLeftEvent(conn, session)
	})
	defer messageService.Unsubscribe(messageservicetopic.GamePlayerLeftMessageTopic, playerLeftSubscriptionToken)

	conn.SetCloseHandler(func(code int, text string) error {
		playersCount -= 1
		messageService.Publish(messageservicetopic.GamePlayerLeftMessageTopic, nil)
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
				break
			case reviveUnitsActionType:
				reviveUnitsAction, err := extractReviveUnitsActionFromMessage(message)
				if err != nil {
					emitErrorEvent(conn, session, err)
				}

				for _, coord := range reviveUnitsAction.Payload.Coordinates {
					coordinate := valueobject.NewCoordinate(coord.X, coord.Y)
					gameRoomService.ReviveGameUnit(gameId, coordinate)
				}

				gameRoom, _ := gameRoomService.GetGameRoom(gameId)
				payload := messageservicetopic.GameUnitsUpdatedMessageTopicPayload{}
				for _, coord := range reviveUnitsAction.Payload.Coordinates {
					coordinate := valueobject.NewCoordinate(coord.X, coord.Y)
					newGameUnit := gameRoom.GetGameUnit(coordinate)
					payloadUnit := messageservicetopic.GameUnitsUpdatedMessageTopicPayloadUnit{
						Coordinate: coord,
						Unit: dto.GameUnitDTO{
							Alive: newGameUnit.GetAlive(),
							Age:   newGameUnit.GetAge(),
						},
					}

					payload = append(payload, payloadUnit)
				}

				message, err := json.Marshal(payload)
				messageService.Publish(messageservicetopic.GameUnitsUpdatedMessageTopic, message)
				break
			default:
				break
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

func emitGameInfoUpdatedEvent(conn *websocket.Conn, session *session, gameId uuid.UUID, gameRoomService gameroomservice.GameRoomService) {
	gameRoom, _ := gameRoomService.GetGameRoom(gameId)
	gameMapSize := gameRoom.GetGameMapSize()
	informationUpdatedEvent := constructInformationUpdatedEvent(&gameMapSize, playersCount)

	sendJSONMessageToClient(conn, session, informationUpdatedEvent)
}

func emitUnitsUpdatedEvent(conn *websocket.Conn, session *session, updateUnitItems *[]unitsUpdatedEventPayloadItem) {
	unitsUpdatedEvent := constructUnitsUpdatedEvent(updateUnitItems)

	sendJSONMessageToClient(conn, session, unitsUpdatedEvent)
}

func emitAreaUpdatedEvent(conn *websocket.Conn, session *session, gameId uuid.UUID, gameRoomService gameroomservice.GameRoomService) {
	if session.gameAreaToWatch == nil {
		return
	}

	gameRoom, _ := gameRoomService.GetGameRoom(gameId)

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

func emitPlayerJoinedEvent(conn *websocket.Conn, session *session) {
	playerJoinedEvent := constructPlayerJoinedEvent()

	sendJSONMessageToClient(conn, session, playerJoinedEvent)
}

func emitPlayerLeftEvent(conn *websocket.Conn, session *session) {
	playerLeftEvent := constructPlayerLeftEvent()

	sendJSONMessageToClient(conn, session, playerLeftEvent)
}
