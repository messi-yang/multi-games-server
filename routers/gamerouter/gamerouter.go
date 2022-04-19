package gamerouter

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/controllers/gamecontroller"
	"github.com/gin-gonic/gin"
)

func SetRouter(engine *gin.Engine) {
	router := engine.Group("/game")
	router.GET("/units", gamecontroller.GetController)
}
