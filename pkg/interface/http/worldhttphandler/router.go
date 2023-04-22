package worldhttphandler

import "github.com/gin-gonic/gin"

func Setup(router *gin.Engine) {
	worldAppService, err := provideWorldAppService()
	if err != nil {
		panic(err)
	}
	httphandler := newHttpHandler(worldAppService)

	routerGroup := router.Group("/api/worlds")
	routerGroup.GET("/:worldId", httphandler.getWorld)
	routerGroup.GET("/", httphandler.queryWorlds)
	routerGroup.POST("/", httphandler.createWorld)
}
