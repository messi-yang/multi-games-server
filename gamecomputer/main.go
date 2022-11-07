package gamecomputer

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/game/domain/dto"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/infrastructure/config"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/interface/integrationeventhandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/eventbusredis"
)

func Start() {
	redisIntegrationEventBus, _ := eventbusredis.NewRedisIntegrationEventBus(eventbusredis.WithRedisService())
	gameApplicationService, _ := applicationservice.NewGameApplicationService(
		applicationservice.WithGameService(),
		applicationservice.WithRedisIntegrationEventBus(),
	)

	size := config.GetConfig().GetGameDimension()
	gameId, _ := gameApplicationService.CreateGame(dto.DimensionDto{
		Width:  size,
		Height: size,
	})

	integrationeventhandler.NewGameIntegrationEventHandler(
		integrationeventhandler.GameIntegrationEventHandlerConfiguration{
			GameApplicationService: gameApplicationService,
			IntegrationEventBus:    redisIntegrationEventBus,
		},
		gameId,
	)
}
