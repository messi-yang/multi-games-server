package apiserver

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/application/service/itemappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/application/service/livegameappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/interface/controller/itemcontroller"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/apiserver/interface/controller/livegamecontroller"
	commonredis "github.com/dum-dum-genius/game-of-liberty-computer/src/common/adapter/notification/redis"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/infrastructure/persistence/postgres"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/service/itemdomainservice"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	router.Use(cors.New(corsConfig))

	postgresClient, err := postgres.NewPostgresClient()
	if err != nil {
		panic(err)
	}
	gameRepository := postgres.NewPostgresGameRepository(postgresClient)
	notificationPublisher := commonredis.NewRedisNotificationPublisher()
	itemDomainService := itemdomainservice.New()
	liveGameAppService := livegameappservice.New(notificationPublisher)
	itemAppService := itemappservice.New(itemDomainService)

	itemController := itemcontroller.New(itemAppService)

	router.Group("/ws/game").GET("/", livegamecontroller.NewController(
		gameRepository,
		liveGameAppService,
	))

	router.GET("/items", itemController.HandleGetAllItems)

	router.Run()
}
