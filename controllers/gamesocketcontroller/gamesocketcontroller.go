package gamesocketcontroller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DumDumGeniuss/game-of-liberty-computer/entities/gameentity"
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

func authenticate(token string) error {
	return nil
}

func Controller(c *gin.Context) {
	err := authenticate("hi")
	if err != nil {
		fmt.Println("Authentication failed")
		return
	}

	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	closeConn := make(chan bool)
	defer conn.Close()

	conn.SetCloseHandler(func(code int, text string) error {
		closeConn <- true
		return nil
	})

	var unsubscribeGameBlockChangeEventChannel chan bool = make(chan bool)
	defer func() {
		unsubscribeGameBlockChangeEventChannel <- true
	}()

	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				break
			}
			eventType, err := getEventTypeFromMessage(msg)
			if err != nil {
				break
			}

			if *eventType == watchGameBlock {
				unsubscribeGameBlockChangeEventChannel <- true

				var watchGameBlockEvent watchGameBlockEvent
				json.Unmarshal(msg, &watchGameBlockEvent)
				unsubscribeGameBlockChangeEventChannel =
					gameentity.SubscribeGameBlockChangeEvent(
						watchGameBlockEvent.Payload.FromX,
						watchGameBlockEvent.Payload.FromY,
						watchGameBlockEvent.Payload.ToX,
						watchGameBlockEvent.Payload.ToY,
						func(gameUnits [][]*gameentity.GameUnit) {
							conn.WriteJSON(gameUnits)
						},
					)
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
