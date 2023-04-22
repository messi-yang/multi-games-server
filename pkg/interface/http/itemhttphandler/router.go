package itemhttphandler

import "github.com/gin-gonic/gin"

func Setup(router *gin.Engine) {
	itemAppService, err := provideItemAppService()
	if err != nil {
		panic(err)
	}
	httpHandler := newHttpHandler(itemAppService)

	routerGroup := router.Group("/api/items")
	routerGroup.GET("/", httpHandler.queryItems)
}
