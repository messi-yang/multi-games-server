package gamerouter

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/controllers/gamecontroller"
	"github.com/gin-gonic/gin"
)

func SetRouter(router *gin.RouterGroup) {
	router.GET("/", gamecontroller.GetController)
}
