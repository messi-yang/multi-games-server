package gamecomputer

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/infrastructure/config"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/interface/integrationeventhandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/game/memory"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/game/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/service/gameservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/service/sandboxservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/eventbusredis"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/infrastructureservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/redisrepository"
)

func Start() {
	redisInfrastructureService := infrastructureservice.NewRedisInfrastructureService()
	redisIntegrationEventBus := eventbusredis.NewRedisIntegrationEventBus(eventbusredis.RedisIntegrationEventBusCallbackConfiguration{
		RedisInfrastructureService: redisInfrastructureService,
	})
	sandboxDomainService := sandboxservice.NewSandboxDomainService(sandboxservice.SandboxDomainServiceConfiguration{
		SandboxRepository: redisrepository.NewSandboxRedisRepository(redisrepository.SandboxRedisRepositoryConfiguration{
			RedisInfrastructureService: redisInfrastructureService,
		}),
	})
	gameDomainService := gameservice.NewGameDomainService(gameservice.GameDomainServiceConfiguration{
		GameRepository: memory.NewGameMemory(),
	})
	gameApplicationService := applicationservice.NewGameApplicationService(
		applicationservice.GameApplicationServiceConfiguration{
			GameDomainService:   gameDomainService,
			IntegrationEventBus: redisIntegrationEventBus,
		},
	)

	gameId, err := sandboxDomainService.GetFirstSandboxId()

	if err != nil {
		size := config.GetConfig().GetGameMapSize()
		mapSize, _ := valueobject.NewMapSize(size, size)
		newSandbox, _ := sandboxDomainService.CreateSandbox(mapSize)
		gameApplicationService.CreateGame(newSandbox)
		gameId = newSandbox.GetId()
	} else {
		game, _ := sandboxDomainService.GetSandbox(gameId)
		gameApplicationService.CreateGame(game)
	}

	integrationeventhandler.NewGameIntegrationEventHandler(
		integrationeventhandler.GameIntegrationEventHandlerConfiguration{
			GameApplicationService: gameApplicationService,
			IntegrationEventBus:    redisIntegrationEventBus,
		},
		gameId,
	)
}
