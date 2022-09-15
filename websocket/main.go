package websocket

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/websocket/gamesocketcontroller"
	"github.com/gin-gonic/gin"
)

func SetWebsocketRouters() {
	router := gin.Default()

	gameRouterGroup := router.Group("/ws/game")
	gameRouterGroup.GET("/", gamesocketcontroller.Controller)

	router.Run()
}
