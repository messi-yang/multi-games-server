package gamerhttphandler

import "github.com/gin-gonic/gin"

func Setup(router *gin.Engine) {
	gamerAppService, err := provideGamerAppService()
	if err != nil {
		panic(err)
	}
	httpHandler := newHttpHandler(gamerAppService)

	routerGroup := router.Group("/api/gamers")
	routerGroup.GET("/", httpHandler.queryGamers)
}
