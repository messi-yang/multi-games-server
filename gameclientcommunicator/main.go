package gameclientcommunicator

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/gameclientcommunicator/http/gameroomsockethandler"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	gameSocketRouter := router.Group("/ws/game")
	gameSocketRouter.GET("/", gameroomsockethandler.Handler)

	router.Run()
}
