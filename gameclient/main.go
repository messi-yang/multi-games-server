package gameclient

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/gameclient/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/gameclient/interface/http/gamehandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/gameclient/presenter/gamehandlerpresenter"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/eventbusredis"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	gameApplicationService, _ := applicationservice.NewGameApplicationService(applicationservice.WithGameService())
	redisIntegrationEventBus, _ := eventbusredis.NewRedisIntegrationEventBus(eventbusredis.WithRedisService())

	router.Group("/ws/game").GET("/", gamehandler.NewHandler(gamehandler.HandlerConfiguration{
		IntegrationEventBus:    redisIntegrationEventBus,
		GameApplicationService: gameApplicationService,
		GameHandlerPresenter:   gamehandlerpresenter.NewGameHandlerPresenter(),
	}))

	router.Run()
}
