package gamerhttpcontroller

import "github.com/gin-gonic/gin"

func Setup(router *gin.Engine) {
	routerGroup := router.Group("/api/gamers")
	routerGroup.GET("/", getGamersHandler)
}
