package gamesocketrouter

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/sockethandlers/gamesockethandler"
	"github.com/gin-gonic/gin"
)

func SetRouter(engine *gin.Engine) {
	router := engine.Group("/game-socket")
	router.GET("/", gamesockethandler.Handler)
}
