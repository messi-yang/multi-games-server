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

	closeConn := make(chan bool)

	conn.SetCloseHandler(func(code int, text string) error {
		return nil
	})

	messageService := messageservice.GetMessageService()
	gameService := gameservice.GetGameService()
	messageService.Subscribe("UNITS_UPDATED", func(_ []byte) {
		fmt.Println("HHHHHIIII")
		gameUnits, err := gameService.GetGameUnitsInArea(
			&gameservice.GameArea{
				From: gameservice.GameCoordinate{X: 0, Y: 0},
				To:   gameservice.GameCoordinate{X: 0, Y: 0},
			},
		)
		if err != nil {
			conn.WriteJSON(err.Error())
			return
		}
		conn.WriteJSON(gameUnits)
	})

	go func() {
		defer func() {
			closeConn <- true
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
		case <-closeConn:
			fmt.Println("Connection closed.")
			return
		}
	}
}
