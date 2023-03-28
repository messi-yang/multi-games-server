package gamesocketcontroller

import (
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	routerGroup := router.Group("/ws/game")
	routerGroup.GET("/", gameConnectionHandler)
}
