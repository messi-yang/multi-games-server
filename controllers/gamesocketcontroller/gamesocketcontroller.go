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
	session := session{
		gameAreaToWatch: &gameservice.GameArea{
			From: gameservice.GameCoordinate{X: 0, Y: 0},
			To:   gameservice.GameCoordinate{X: 3, Y: 3},
		},
	}

	gameSize, _ := gameService.GetGameSize()
	gameInfoUpdatedEvent := constructGameInfoUpdatedEvent(gameSize, playersCount)
	conn.WriteJSON(gameInfoUpdatedEvent)

	messageService.Publish("PLAYER_JOINED", nil)

	unitsUpdatedSubscriptionToken := messageService.Subscribe("UNITS_UPDATED", func(_ []byte) {
		if session.gameAreaToWatch == nil {
			return
		}
		gameUnits, err := gameService.GetGameUnitsInArea(
			session.gameAreaToWatch,
		)
		if err != nil {
			errorEvent := constructErrorHappenedEvent(err.Error())
			conn.WriteJSON(errorEvent)
			return
		}

		unitsUpdatedEvent := constructUnitsUpdatedEvent(session.gameAreaToWatch, gameUnits)
		conn.WriteJSON(unitsUpdatedEvent)
	})
	defer messageService.Unsubscribe("UNITS_UPDATED", unitsUpdatedSubscriptionToken)

	playerJoinedSubscriptionToken := messageService.Subscribe("PLAYER_JOINED", func(_ []byte) {
		playerJoinedEvent := constructPlayerJoinedEvent()
		conn.WriteJSON(playerJoinedEvent)
	})
	defer messageService.Unsubscribe("PLAYER_JOINED", playerJoinedSubscriptionToken)

	playerLeftSubscriptionToken := messageService.Subscribe("PLAYER_LEFT", func(_ []byte) {
		playerLeftEvent := constructPlayerLeftEvent()
		conn.WriteJSON(playerLeftEvent)
	})
	defer messageService.Unsubscribe("PLAYER_LEFT", playerLeftSubscriptionToken)

	conn.SetCloseHandler(func(code int, text string) error {
		playersCount -= 1
		messageService.Publish("PLAYER_LEFT", nil)
		return nil
	})

	go func() {
		defer func() {
			closeConnFlag <- true
		}()

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				errorEvent := constructErrorHappenedEvent(err.Error())
				conn.WriteJSON(errorEvent)
				break
			}

			actionType, err := getActionTypeFromMessage(message)
			if err != nil {
				errorEvent := constructErrorHappenedEvent(err.Error())
				conn.WriteJSON(errorEvent)
			}

			switch *actionType {
			case watchUnitsActionType:
				watchUnitsAction, err := extractWatchUnitsActionFromMessage(message)
				if err != nil {
					errorEvent := constructErrorHappenedEvent(err.Error())
					conn.WriteJSON(errorEvent)
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
