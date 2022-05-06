package gamesocketcontroller

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/DumDumGeniuss/game-of-liberty-computer/services/gameservice"
	"github.com/DumDumGeniuss/game-of-liberty-computer/services/messageservice"
	"github.com/gin-gonic/gin"
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

	messageService := messageservice.GetMessageService()
	gameService := gameservice.GetGameService()

	playersCount += 1
	session := &session{
		gameAreaToWatch: &gameservice.GameArea{
			From: gameservice.GameCoordinate{X: 0, Y: 0},
			To:   gameservice.GameCoordinate{X: 3, Y: 3},
		},
		socketLocker: sync.RWMutex{},
	}

	emitGameInfoUpdatedEvent(conn, session, gameService)
	messageService.Publish(messageservice.GamePlayerJoined, nil)

	areaUpdatedSubscriptionToken := messageService.Subscribe(messageservice.GameAreaUpdated, func(_ []byte) {
		emitAreaUpdatedEvent(conn, session, gameService)
	})
	defer messageService.Unsubscribe(messageservice.GameAreaUpdated, areaUpdatedSubscriptionToken)

	playerJoinedSubscriptionToken := messageService.Subscribe(messageservice.GamePlayerJoined, func(_ []byte) {
		emitPlayerJoinedEvent(conn, session)
	})
	defer messageService.Unsubscribe(messageservice.GamePlayerJoined, playerJoinedSubscriptionToken)

	playerLeftSubscriptionToken := messageService.Subscribe(messageservice.GamePlayerLeft, func(_ []byte) {
		emitPlayerLeftEvent(conn, session)
	})
	defer messageService.Unsubscribe(messageservice.GamePlayerLeft, playerLeftSubscriptionToken)

	conn.SetCloseHandler(func(code int, text string) error {
		playersCount -= 1
		messageService.Publish(messageservice.GamePlayerLeft, nil)
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
				session.gameAreaToWatch = &watchAreaAction.Payload.Area
				break
			case reviveUnitsActionType:
				reviveUnitsAction, err := extractReviveUnitsActionFromMessage(message)
				if err != nil {
					emitErrorEvent(conn, session, err)
				}

				for _, coord := range reviveUnitsAction.Payload.Coordinates {
					gameService.ReviveGameUnit(&coord)
				}
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

func emitGameInfoUpdatedEvent(conn *websocket.Conn, session *session, gameService gameservice.GameService) {
	gameSize, _ := gameService.GetGameSize()
	informationUpdatedEvent := constructInformationUpdatedEvent(gameSize, playersCount)

	sendJSONMessageToClient(conn, session, informationUpdatedEvent)
}

func emitAreaUpdatedEvent(conn *websocket.Conn, session *session, gameService gameservice.GameService) {
	if session.gameAreaToWatch == nil {
		return
	}

	gameUnits, err := gameService.GetGameUnitsInArea(
		session.gameAreaToWatch,
	)
	if err != nil {
		emitErrorEvent(conn, session, err)
		return
	}

	areaUpdatedEvent := constructAreaUpdatedEvent(session.gameAreaToWatch, gameUnits)

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
