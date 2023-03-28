package assethttpcontroller

import "github.com/gin-gonic/gin"

func Setup(router *gin.Engine) {
	router.Static("/asset", "./asset")
}
