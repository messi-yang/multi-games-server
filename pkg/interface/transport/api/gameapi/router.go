package gameapi

import (
	"github.com/gin-gonic/gin"
)

func SetRouter(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/", gameConnectionHandler)
}
