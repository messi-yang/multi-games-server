package worldhttphandler

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/worldappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/pgrepo"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	worldRepo, err := pgrepo.NewWorldRepo()
	if err != nil {
		panic(err)
	}
	itemRepo, err := pgrepo.NewItemRepo()
	if err != nil {
		panic(err)
	}
	unitRepo, err := pgrepo.NewUnitRepo()
	if err != nil {
		panic(err)
	}
	worldAppService := worldappsrv.NewService(worldRepo, unitRepo, itemRepo)
	httphandler := newHttpHandler(worldAppService)

	routerGroup := router.Group("/api/worlds")
	routerGroup.GET("/:worldId", httphandler.getWorld)
	routerGroup.GET("/", httphandler.queryWorlds)
	routerGroup.POST("/", httphandler.createWorld)
}
