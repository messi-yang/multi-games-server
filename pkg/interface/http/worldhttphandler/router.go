package worldhttphandler

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/worldappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/pgrepository"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	worldRepository, err := pgrepository.NewWorldRepository()
	if err != nil {
		panic(err)
	}
	itemRepository, err := pgrepository.NewItemRepository()
	if err != nil {
		panic(err)
	}
	unitRepository, err := pgrepository.NewUnitRepository()
	if err != nil {
		panic(err)
	}
	worldAppService := worldappservice.NewService(worldRepository, unitRepository, itemRepository)
	httphandler := newHttpHandler(worldAppService)

	routerGroup := router.Group("/api/worlds")
	routerGroup.GET("/:worldId", httphandler.getWorld)
	routerGroup.GET("/", httphandler.queryWorlds)
	routerGroup.POST("/", httphandler.createWorld)
}
