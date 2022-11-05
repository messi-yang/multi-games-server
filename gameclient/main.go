package gameclient

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/gameclient/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/gameclient/interface/http/gamehandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/gameclient/presenter/gamehandlerpresenter"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/service/sandboxservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/eventbusredis"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/infrastructureservice"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	redisService := infrastructureservice.NewRedisService()
	sandboxDomainService, _ := sandboxservice.NewSandboxDomainService(sandboxservice.WithSandboxRedis())
	gameApplicationService := applicationservice.NewGameApplicationService(applicationservice.GameApplicationServiceConfiguration{
		SandboxDomainService: sandboxDomainService,
	})
	redisIntegrationEventBus := eventbusredis.NewRedisIntegrationEventBus(eventbusredis.RedisIntegrationEventBusCallbackConfiguration{
		RedisService: redisService,
	})

	router.Group("/ws/game").GET("/", gamehandler.NewHandler(gamehandler.HandlerConfiguration{
		IntegrationEventBus:    redisIntegrationEventBus,
		GameApplicationService: gameApplicationService,
		GameHandlerPresenter:   gamehandlerpresenter.NewGameHandlerPresenter(),
	}))

	router.Run()
}
