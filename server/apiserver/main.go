package apiserver

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/domainservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/apiserver/application/appservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/apiserver/interface/httpcontroller/itemcontroller"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/apiserver/interface/httpcontroller/livegamecontroller"
	commonredis "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification/redis"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/persistence/postgres"
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
	itemDomainService := domainservice.NewItemDomainServe()
	liveGameAppService := appservice.NewLiveGameAppService(notificationPublisher)
	itemAppService := appservice.NewItemAppService(itemDomainService)

	itemController := itemcontroller.NewItemController(itemAppService)

	router.Group("/ws/game").GET("/", livegamecontroller.NewController(
		itemDomainService,
		gameRepository,
		liveGameAppService,
	))

	router.GET("/items", itemController.HandleGetAllItems)

	router.Run()
}
