package gameclient

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/gameclient/gamesocketcontroller"
	"github.com/gin-gonic/gin"
)

func SetWebsocketRouters() {
	router := gin.Default()

	gameSocketRouter := router.Group("/ws/game")
	gameSocketRouter.GET("/", gamesocketcontroller.Controller)

	router.Run()
}
