package apiserver

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/service/gameservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/apiserver/application/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/apiserver/interface/httpcontroller/livegamecontroller"
	commonredis "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification/redis"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/persistence/postgres"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	postgresClient, err := postgres.NewPostgresClient()
	if err != nil {
		panic(err)
	}
	gameRepository := postgres.NewPostgresGameRepository(postgresClient)
	gameService := gameservice.NewGameService(gameRepository)
	notificationPublisher := commonredis.NewRedisNotificationPublisher()
	liveGameAppService := service.NewLiveGameAppService(notificationPublisher)
	gameAppService := service.NewGameAppService(gameService)

	router.Group("/ws/game").GET("/", livegamecontroller.NewController(
		gameRepository,
		gameAppService,
		liveGameAppService,
	))

	router.Run()
}
