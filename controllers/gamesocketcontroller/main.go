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

func Controller(c *gin.Context) {
	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()

	session := session{
		watchArea: &gameservice.GameArea{
			From: gameservice.GameCoordinate{X: 0, Y: 0},
			To:   gameservice.GameCoordinate{X: 0, Y: 0},
		},
	}
	fmt.Println(session)

	closeConnFlag := make(chan bool)

	messageService := messageservice.GetMessageService()
	gameService := gameservice.GetGameService()
	unitsUpdatedSubscriptionToken := messageService.Subscribe("UNITS_UPDATED", func(_ []byte) {
		if session.watchArea == nil {
			return
		}
		gameUnits, err := gameService.GetGameUnitsInArea(
			session.watchArea,
		)
		if err != nil {
			conn.WriteJSON(err.Error())
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
			fmt.Println("Connection closed.")
			return
		}
	}
}
