package itemhttphandler

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/itemappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/pgrepository"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	itemRepository, err := pgrepository.NewItemRepository()
	if err != nil {
		panic(err)
	}
	itemAppService := itemappservice.NewService(itemRepository)
	httpHandler := newHttpHandler(itemAppService)

	routerGroup := router.Group("/api/items")
	routerGroup.GET("/", httpHandler.queryItems)
}
