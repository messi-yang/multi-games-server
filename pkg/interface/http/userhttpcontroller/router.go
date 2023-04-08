package userhttpcontroller

import "github.com/gin-gonic/gin"

func Setup(router *gin.Engine) {
	routerGroup := router.Group("/api/users")
	routerGroup.GET("/", getUsersHandler)
}
