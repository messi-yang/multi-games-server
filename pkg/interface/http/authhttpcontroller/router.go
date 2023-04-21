package authhttpcontroller

import "github.com/gin-gonic/gin"

func Setup(router *gin.Engine) {
	routerGroup := router.Group("/auth")
	routerGroup.GET("/oauth2/google", goToGoogleAuthUrlHandler)
	routerGroup.GET("/oauth2/google/redirect", googleAuthCallbackHandler)
}
