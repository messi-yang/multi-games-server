package worldhttphandler

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/worldappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/pgrepo"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	worldRepository, err := pgrepo.NewWorldRepository()
	if err != nil {
		panic(err)
	}
	itemRepository, err := pgrepo.NewItemRepository()
	if err != nil {
		panic(err)
	}
	unitRepository, err := pgrepo.NewUnitRepository()
	if err != nil {
		panic(err)
	}
	worldAppService := worldappsrv.NewService(worldRepository, unitRepository, itemRepository)
	httphandler := newHttpHandler(worldAppService)

	routerGroup := router.Group("/api/worlds")
	routerGroup.GET("/:worldId", httphandler.getWorld)
	routerGroup.GET("/", httphandler.queryWorlds)
	routerGroup.POST("/", httphandler.createWorld)
}
