package gamecomputer

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/infrastructure/config"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/interface/integrationeventhandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/game/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/service/gameservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/service/sandboxservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/eventbusredis"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/infrastructureservice"
)

func Start() {
	redisService := infrastructureservice.NewRedisService()
	redisIntegrationEventBus := eventbusredis.NewRedisIntegrationEventBus(eventbusredis.RedisIntegrationEventBusCallbackConfiguration{
		RedisService: redisService,
	})
	sandboxDomainService, _ := sandboxservice.NewSandboxDomainService(sandboxservice.WithSandboxRedis())
	gameDomainService, _ := gameservice.NewGameService(gameservice.WithGameMemory())
	gameApplicationService := applicationservice.NewGameApplicationService(
		applicationservice.GameApplicationServiceConfiguration{
			GameService:         gameDomainService,
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
