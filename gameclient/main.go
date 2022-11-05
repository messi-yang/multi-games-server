package gameclient

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/gameclient/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/gameclient/interface/http/gameroomhandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/gameclient/presenter/gameroomhandlerpresenter"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/service/sandboxservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/eventbusredis"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/infrastructureservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/redisrepository"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	redisInfrastructureService := infrastructureservice.NewRedisInfrastructureService()
	sandboxDomainService := sandboxservice.NewSandboxDomainService(sandboxservice.SandboxDomainServiceConfiguration{
		SandboxRepository: redisrepository.NewSandboxRedisRepository(redisrepository.SandboxRedisRepositoryConfiguration{
			RedisInfrastructureService: redisInfrastructureService,
		}),
	})
	gameApplicationService := applicationservice.NewGameApplicationService(applicationservice.GameApplicationServiceConfiguration{
		SandboxDomainService: sandboxDomainService,
	})
	redisIntegrationEventBus := eventbusredis.NewRedisIntegrationEventBus(eventbusredis.RedisIntegrationEventBusCallbackConfiguration{
		RedisInfrastructureService: redisInfrastructureService,
	})

	router.Group("/ws/game").GET("/", gameroomhandler.NewHandler(gameroomhandler.HandlerConfiguration{
		IntegrationEventBus:      redisIntegrationEventBus,
		GameApplicationService:   gameApplicationService,
		GameRoomHandlerPresenter: gameroomhandlerpresenter.NewGameRoomHandlerPresenter(),
	}))

	router.Run()
}
