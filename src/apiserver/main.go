package apiserver

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/application/service/itemappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/application/service/livegameappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/interface/controller/itemcontroller"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/interface/controller/livegamecontroller"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/infrastructure/messaging/redisintgreventpublisher"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/infrastructure/persistence/gamepsqlrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/infrastructure/persistence/itemmemoryrepo"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/gormdb"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	router.Use(cors.New(corsConfig))

	gormDb, err := gormdb.New()
	if err != nil {
		panic(err)
	}
	gameRepo := gamepsqlrepo.New(gormDb)
	intgrEventPublisher := redisintgreventpublisher.New()
	itemRepo := itemmemoryrepo.New()
	liveGameAppService := livegameappservice.New(intgrEventPublisher)
	itemAppService := itemappservice.New(itemRepo)

	itemController := itemcontroller.New(itemAppService)

	router.Group("/ws/game").GET("/", livegamecontroller.NewController(
		gameRepo,
		liveGameAppService,
	))

	router.GET("/items", itemController.HandleGetAllItems)

	router.Run()
}
