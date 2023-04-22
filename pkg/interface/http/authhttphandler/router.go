package authhttphandler

import "github.com/gin-gonic/gin"

func Setup(router *gin.Engine) {
	googleAuthInfraService := provideGoogleOauthInfraService()
	identityAppService, err := provideIdentityAppService()
	if err != nil {
		panic(err)
	}
	httpHandler := newHttpHandler(googleAuthInfraService, identityAppService)

	routerGroup := router.Group("/auth")
	routerGroup.GET("/oauth2/google", httpHandler.goToGoogleAuthUrl)
	routerGroup.GET("/oauth2/google/redirect", httpHandler.googleAuthCallback)
}
