package worldapi

import "github.com/gin-gonic/gin"

func SetRouter(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/", QueryWorldHandler)
	routerGroup.POST("/", CreateWorldHandler)
}
