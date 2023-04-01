package worldhttpcontroller

import "github.com/gin-gonic/gin"

func Setup(router *gin.Engine) {
	routerGroup := router.Group("/api/worlds")
	routerGroup.GET("/:worldId", getWorldHandler)
	routerGroup.GET("/", getWorldsHandler)
	routerGroup.POST("/", createWorldHandler)
}
