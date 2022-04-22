package gamesocketcontroller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DumDumGeniuss/game-of-liberty-computer/workers/gameworker"
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

	sessionHash := generateRandomHash(12)

	conn.SetCloseHandler(func(code int, text string) error {
		closeConn <- true
		return nil
	})

	go func() {
		defer func() {
			gameworker.UnsubscribeGameBlockChangeEvent(sessionHash)
		}()

		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				break
			}
			actionType, err := getActionTypeFromMessage(msg)
			if err != nil {
				break
			}

			if *actionType == watchGameBlock {
				var watchGameBlockAction watchGameBlockAction
				json.Unmarshal(msg, &watchGameBlockAction)
				gameworker.SubscribeGameBlockChangeEvent(
					sessionHash,
					gameworker.GameBlockArea{
						FromX: watchGameBlockAction.Payload.Area.From.X,
						FromY: watchGameBlockAction.Payload.Area.From.Y,
						ToX:   watchGameBlockAction.Payload.Area.To.X,
						ToY:   watchGameBlockAction.Payload.Area.To.Y,
					},
					func(gameUnits [][]*gameworker.GameUnit) {
						event := gameBlockUpdatedEvent{
							Type: gaemBlockUpdated,
							Payload: gameBlockUpdatedEventPayload{
								Area:  watchGameBlockAction.Payload.Area,
								Units: gameUnits,
							},
						}
						conn.WriteJSON(event)
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
