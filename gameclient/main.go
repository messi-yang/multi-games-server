package gameclient

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/gameclient/socketcontrollers/gamesocketcontroller"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	gameSocketRouter := router.Group("/ws/game")
	gameSocketRouter.GET("/", gamesocketcontroller.Controller)

	router.Run()
}
