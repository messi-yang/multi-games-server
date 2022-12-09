package apiserver

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/service/gameservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/apiserver/application/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/apiserver/interface/httpcontroller/livegamecontroller"
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
	liveGameApplicationService, _ := service.NewLiveGameApplicationService(
		service.WithRedisNotificationPublisher(),
	)
	gameApplicationService := service.NewGameApplicationService(
		gameService,
	)

	router.Group("/ws/game").GET("/", livegamecontroller.NewController(livegamecontroller.Configuration{
		LiveGameApplicationService: liveGameApplicationService,
		GameApplicationService:     gameApplicationService,
	}))

	router.Run()
}
