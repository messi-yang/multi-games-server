package gamesockethandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func Setup(router *gin.Engine) {
	gameAppService, err := provideGameAppService()
	if err != nil {
		panic(err)
	}
	httpHandler := newHttpHandler(gameAppService, websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	})

	routerGroup := router.Group("/ws/game")
	routerGroup.GET("/", httpHandler.gameConnection)
}
