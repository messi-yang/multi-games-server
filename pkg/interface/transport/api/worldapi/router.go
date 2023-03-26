package worldapi

import "github.com/gin-gonic/gin"

func SetRouter(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/", queryWorldHandler)
	routerGroup.POST("/", createWorldHandler)
}
