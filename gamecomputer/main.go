package gamecomputer

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/infrastructure/config"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/infrastructure/memoryrepository"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/interface/integrationeventhandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/interface/task"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/domainservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/eventbusredis"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/infrastructureservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/redisrepository"
	"github.com/google/uuid"
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

	games := gameDomainService.GetAllGames()
	var gameId uuid.UUID

	if len(games) > 0 {
		gameRoomApplicationService.LoadGameRoom(games[0])
		gameId = games[0].GetId()
	} else {
		size := config.GetConfig().GetGameMapSize()
		mapSize, _ := valueobject.NewMapSize(size, size)
		game, _ := gameDomainService.CreateGame(mapSize)
		gameRoomApplicationService.LoadGameRoom(game)
		gameId = game.GetId()
	}

	task.NewTickUnitMapTask(task.TickUnitMapTaskConfiguration{
		GameRoomApplicationService: gameRoomApplicationService,
	})

	integrationeventhandler.NewGameRoomIntegrationEventHandler(
		integrationeventhandler.GameRoomIntegrationEventHandlerConfiguration{
			GameRoomApplicationService: gameRoomApplicationService,
			IntegrationEventBus:        redisIntegrationEventBus,
		},
		gameId,
	)
}
