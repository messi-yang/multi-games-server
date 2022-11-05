package gameclient

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/gameclient/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/gameclient/interface/http/gamehandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/gameclient/presenter/gamehandlerpresenter"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/sandbox/redis"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/service/sandboxservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/eventbusredis"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/infrastructureservice"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	redisInfrastructureService := infrastructureservice.NewRedisInfrastructureService()
	sandboxDomainService := sandboxservice.NewSandboxDomainService(sandboxservice.SandboxDomainServiceConfiguration{
		SandboxRepository: redis.NewSandboxRedis(redis.SandboxRedisConfiguration{
			RedisInfrastructureService: redisInfrastructureService,
		}),
	})
	gameApplicationService := applicationservice.NewGameApplicationService(applicationservice.GameApplicationServiceConfiguration{
		SandboxDomainService: sandboxDomainService,
	})
	redisIntegrationEventBus := eventbusredis.NewRedisIntegrationEventBus(eventbusredis.RedisIntegrationEventBusCallbackConfiguration{
		RedisInfrastructureService: redisInfrastructureService,
	})

	router.Group("/ws/game").GET("/", gamehandler.NewHandler(gamehandler.HandlerConfiguration{
		IntegrationEventBus:    redisIntegrationEventBus,
		GameApplicationService: gameApplicationService,
		GameHandlerPresenter:   gamehandlerpresenter.NewGameHandlerPresenter(),
	}))

	router.Run()
}
