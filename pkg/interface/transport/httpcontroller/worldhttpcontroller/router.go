package worldhttpcontroller

import "github.com/gin-gonic/gin"

func SetRouter(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/", QueryHandler)
}
