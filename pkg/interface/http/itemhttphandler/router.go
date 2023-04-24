package itemhttphandler

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/itemappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/pgrepo"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	itemRepository, err := pgrepo.NewItemRepository()
	if err != nil {
		panic(err)
	}
	itemAppService := itemappsrv.NewService(itemRepository)
	httpHandler := newHttpHandler(itemAppService)

	routerGroup := router.Group("/api/items")
	routerGroup.GET("/", httpHandler.queryItems)
}
