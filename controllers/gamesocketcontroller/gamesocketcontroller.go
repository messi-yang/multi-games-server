package gamesocketcontroller

import (
	"fmt"
	"net/http"

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
	}

	emitGameInfoUpdatedEvent(conn, gameService)
	messageService.Publish(messageservice.GamePlayerJoined, nil)

	unitsUpdatedSubscriptionToken := messageService.Subscribe(messageservice.GameUnitsUpdated, func(_ []byte) {
		emitUnitsUpdatedEvent(conn, session, gameService)
	})
	defer messageService.Unsubscribe(messageservice.GameUnitsUpdated, unitsUpdatedSubscriptionToken)

	playerJoinedSubscriptionToken := messageService.Subscribe(messageservice.GamePlayerJoined, func(_ []byte) {
		emitPlayerJoinedEvent(conn)
	})
	defer messageService.Unsubscribe(messageservice.GamePlayerJoined, playerJoinedSubscriptionToken)

	playerLeftSubscriptionToken := messageService.Subscribe(messageservice.GamePlayerLeft, func(_ []byte) {
		emitPlayerLeftEvent(conn)
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
				emitErrorEvent(conn, err)
				break
			}

			actionType, err := getActionTypeFromMessage(message)
			if err != nil {
				emitErrorEvent(conn, err)
			}

			switch *actionType {
			case watchUnitsActionType:
				watchUnitsAction, err := extractWatchUnitsActionFromMessage(message)
				if err != nil {
					emitErrorEvent(conn, err)
				}
				session.gameAreaToWatch = &watchUnitsAction.Payload.Area
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

func emitErrorEvent(conn *websocket.Conn, err error) {
	errorEvent := constructErrorHappenedEvent(err.Error())
	conn.WriteJSON(errorEvent)
}

func emitGameInfoUpdatedEvent(conn *websocket.Conn, gameService gameservice.GameService) {
	gameSize, _ := gameService.GetGameSize()
	gameInfoUpdatedEvent := constructGameInfoUpdatedEvent(gameSize, playersCount)
	conn.WriteJSON(gameInfoUpdatedEvent)
}

func emitUnitsUpdatedEvent(conn *websocket.Conn, session *session, gameService gameservice.GameService) {
	if session.gameAreaToWatch == nil {
		return
	}
	gameUnits, err := gameService.GetGameUnitsInArea(
		session.gameAreaToWatch,
	)
	if err != nil {
		emitErrorEvent(conn, err)
		return
	}

	unitsUpdatedEvent := constructUnitsUpdatedEvent(session.gameAreaToWatch, gameUnits)
	conn.WriteJSON(unitsUpdatedEvent)
}

func emitPlayerJoinedEvent(conn *websocket.Conn) {
	playerJoinedEvent := constructPlayerJoinedEvent()
	conn.WriteJSON(playerJoinedEvent)
}

func emitPlayerLeftEvent(conn *websocket.Conn) {
	playerLeftEvent := constructPlayerLeftEvent()
	conn.WriteJSON(playerLeftEvent)
}
