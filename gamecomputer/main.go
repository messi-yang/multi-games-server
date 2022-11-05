package gamecomputer

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/infrastructure/config"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/infrastructure/memoryrepository"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/interface/integrationeventhandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/game/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/service/gameroomdomainservice"
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
	gameRoomDomainService := gameroomdomainservice.NewGameRoomDomainService(gameroomdomainservice.GameRoomDomainServiceConfiguration{
		GameRoomRepository: memoryrepository.NewGameRoomMemoryRepository(),
	})
	gameRoomApplicationService := applicationservice.NewGameRoomApplicationService(
		applicationservice.GameRoomApplicationServiceConfiguration{
			GameRoomDomainService: gameRoomDomainService,
			IntegrationEventBus:   redisIntegrationEventBus,
		},
	)

	gameId, err := sandboxDomainService.GetFirstSandboxId()

	if err != nil {
		size := config.GetConfig().GetGameMapSize()
		mapSize, _ := valueobject.NewMapSize(size, size)
		newSandbox, _ := sandboxDomainService.CreateSandbox(mapSize)
		gameRoomApplicationService.CreateGameRoom(newSandbox)
		gameId = newSandbox.GetId()
	} else {
		game, _ := sandboxDomainService.GetSandbox(gameId)
		gameRoomApplicationService.CreateGameRoom(game)
	}

	integrationeventhandler.NewGameRoomIntegrationEventHandler(
		integrationeventhandler.GameRoomIntegrationEventHandlerConfiguration{
			GameRoomApplicationService: gameRoomApplicationService,
			IntegrationEventBus:        redisIntegrationEventBus,
		},
		gameId,
	)
}
