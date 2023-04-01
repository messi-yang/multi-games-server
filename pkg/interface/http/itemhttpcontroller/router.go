package itemhttpcontroller

import "github.com/gin-gonic/gin"

func Setup(router *gin.Engine) {
	routerGroup := router.Group("/api/items")
	routerGroup.GET("/", queryWorldHandler)
}
