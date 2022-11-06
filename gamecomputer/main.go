package gamecomputer

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/infrastructure/config"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/interface/integrationeventhandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/game/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/service/gameservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/eventbusredis"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/infrastructureservice"
)

func Start() {
	redisService := infrastructureservice.NewRedisService()
	redisIntegrationEventBus := eventbusredis.NewRedisIntegrationEventBus(eventbusredis.RedisIntegrationEventBusCallbackConfiguration{
		RedisService: redisService,
	})
	gameService, _ := gameservice.NewGameService(gameservice.WithGameMemory())
	gameApplicationService := applicationservice.NewGameApplicationService(
		applicationservice.GameApplicationServiceConfiguration{
			GameService:         gameService,
			IntegrationEventBus: redisIntegrationEventBus,
		},
	)

	size := config.GetConfig().GetGameDimension()
	dimension, _ := valueobject.NewDimension(size, size)
	gameId, _ := gameApplicationService.CreateGame(dimension)

	integrationeventhandler.NewGameIntegrationEventHandler(
		integrationeventhandler.GameIntegrationEventHandlerConfiguration{
			GameApplicationService: gameApplicationService,
			IntegrationEventBus:    redisIntegrationEventBus,
		},
		gameId,
	)
}
