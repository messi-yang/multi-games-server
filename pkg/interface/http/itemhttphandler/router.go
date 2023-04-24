package itemhttphandler

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/itemappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/pgrepo"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	itemRepo, err := pgrepo.NewItemRepo()
	if err != nil {
		panic(err)
	}
	itemAppService := itemappsrv.NewService(itemRepo)
	httpHandler := newHttpHandler(itemAppService)

	routerGroup := router.Group("/api/items")
	routerGroup.GET("/", httpHandler.queryItems)
}
