package gamesocketcontroller

import (
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

func Controller(c *gin.Context) {
	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.Error(err)
		return
	}
	defer conn.Close()

	session := session{
		gameAreaToWatch: &gameservice.GameArea{
			From: gameservice.GameCoordinate{X: 0, Y: 0},
			To:   gameservice.GameCoordinate{X: 3, Y: 3},
		},
	}

	closeConnFlag := make(chan bool)

	messageService := messageservice.GetMessageService()
	gameService := gameservice.GetGameService()
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
		conn.WriteJSON(gameUnits)
	})
	defer messageService.Unsubscribe("UNITS_UPDATED", unitsUpdatedSubscriptionToken)

	go func() {
		defer func() {
			closeConnFlag <- true
		}()

		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				break
			}
		}
	}()

	for {
		select {
		case <-closeConnFlag:
			return
		}
	}
}
