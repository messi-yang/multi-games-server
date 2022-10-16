package gamecomputer

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/infrastructure/config"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/infrastructure/memoryrepository"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/interface/integrationeventhandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/domainservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/eventbusredis"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/infrastructureservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/redisrepository"
)

func Start() {
	redisInfrastructureService := infrastructureservice.NewRedisInfrastructureService()
	redisIntegrationEventBus := eventbusredis.NewRedisIntegrationEventBus(eventbusredis.RedisIntegrationEventBusCallbackConfiguration{
		RedisInfrastructureService: redisInfrastructureService,
	})
	gameDomainService := domainservice.NewGameDomainService(domainservice.GameDomainServiceConfiguration{
		GameRepository: redisrepository.NewGameRedisRepository(redisrepository.GameRedisRepositoryConfiguration{
			RedisInfrastructureService: redisInfrastructureService,
		}),
	})
	gameRoomDomainService := domainservice.NewGameRoomDomainService(domainservice.GameRoomDomainServiceConfiguration{
		GameRoomRepository: memoryrepository.NewGameRoomMemoryRepository(),
	})
	gameRoomApplicationService := applicationservice.NewGameRoomApplicationService(
		applicationservice.GameRoomApplicationServiceConfiguration{
			GameRoomDomainService: gameRoomDomainService,
			IntegrationEventBus:   redisIntegrationEventBus,
		},
	)

	gameId, err := gameDomainService.GetFirstGameId()

	if err != nil {
		size := config.GetConfig().GetGameMapSize()
		mapSize, _ := valueobject.NewMapSize(size, size)
		game, _ := gameDomainService.CreateGame(mapSize)
		gameRoomApplicationService.CreateGameRoom(game)
		gameId = game.GetId()
	} else {
		game, _ := gameDomainService.GetGame(gameId)
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
