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
	sandboxService, _ := sandboxservice.NewSandboxService(sandboxservice.WithSandboxRedis())
	gameService, _ := gameservice.NewGameService(gameservice.WithGameMemory())
	gameApplicationService := applicationservice.NewGameApplicationService(
		applicationservice.GameApplicationServiceConfiguration{
			GameService:         gameService,
			IntegrationEventBus: redisIntegrationEventBus,
		},
	)

	gameId, err := sandboxService.GetFirstSandboxId()

	if err != nil {
		size := config.GetConfig().GetGameMapSize()
		mapSize, _ := valueobject.NewMapSize(size, size)
		newSandbox, _ := sandboxService.CreateSandbox(mapSize)
		gameApplicationService.CreateGame(newSandbox)
		gameId = newSandbox.GetId()
	} else {
		game, _ := sandboxService.GetSandbox(gameId)
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
