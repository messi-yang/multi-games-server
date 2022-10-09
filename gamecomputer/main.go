package gamecomputer

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/infrastructure/config"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/infrastructure/memoryrepository"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/interface/integrationeventhandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/interface/task"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/domainservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/entity"
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
	gameRoomDomainService := domainservice.NewGameRoomDomainService(domainservice.GameRoomDomainServiceConfiguration{
		GameRoomRealtimeRepository: memoryrepository.NewGameRoomRealtimeMemoryRepository(),
		GameRoomPersistentRepository: redisrepository.NewGameRoomPersistentRedisRepository(redisrepository.GameRoomPersistentRedisRepositoryConfiguration{
			RedisInfrastructureService: redisInfrastructureService,
		}),
	})
	gameRoomApplicationService := applicationservice.NewGameRoomApplicationService(
		applicationservice.GameRoomApplicationServiceConfiguration{
			GameRoomDomainService: gameRoomDomainService,
			IntegrationEventBus:   redisIntegrationEventBus,
		},
	)

	size := config.GetConfig().GetGameMapSize()
	mapSize, _ := valueobject.NewMapSize(size, size)
	game := entity.NewGame(mapSize, time.Second.Microseconds())
	newGameRoomId, err := gameRoomApplicationService.CreateGameRoom(game)
	if err != nil {
		panic(err.Error())
	}
	gameRoomApplicationService.LoadGameRoom(newGameRoomId)
	redisInfrastructureService.Set("game_id", []byte(newGameRoomId.String()))

	task.NewTickUnitMapTask(task.TickUnitMapTaskConfiguration{
		GameRoomApplicationService: gameRoomApplicationService,
	})

	integrationeventhandler.NewGameRoomIntegrationEventHandler(
		integrationeventhandler.GameRoomIntegrationEventHandlerConfiguration{
			GameRoomApplicationService: gameRoomApplicationService,
			IntegrationEventBus:        redisIntegrationEventBus,
		},
		newGameRoomId,
	)
}
