package gameapi

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/service/gamesocketappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/common/client/redisclient"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/infrastructure/messaging/redisinteventpublisher"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/infrastructure/persistence/memrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/infrastructure/persistence/postgres"
	"github.com/gin-gonic/gin"
)

func SetRouter(routerGroup *gin.RouterGroup) {
	redisClient := redisclient.New()
	intEventPublisher := redisinteventpublisher.New(redisClient)

	itemRepository, err := postgres.NewItemRepository()
	if err != nil {
		panic(err)
	}
	playerRepository := memrepo.NewPlayerMemRepository()
	worldRepository, err := postgres.NewWorldRepository()
	if err != nil {
		panic(err)
	}
	unitRepository, err := postgres.NewUnitRepository()
	if err != nil {
		panic(err)
	}
	gameSocketAppService := gamesocketappservice.NewService(intEventPublisher, worldRepository, playerRepository, unitRepository, itemRepository)

	gameSocketApiController := NewController(gameSocketAppService, redisClient)

	routerGroup.GET("/", gameSocketApiController.HandleGameConnection)
}
