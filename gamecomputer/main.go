package gamecomputer

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/infrastructure/config"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/infrastructure/memoryrepository"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/interface/integrationeventhandler"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/interface/task"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/domainservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/eventbusredis"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/infrastructureservice"
)

func Start() {
	redisInfrastructureService := infrastructureservice.NewRedisInfrastructureService()
	redisIntegrationEventBus := eventbusredis.NewRedisIntegrationEventBus(eventbusredis.RedisIntegrationEventBusCallbackConfiguration{
		RedisInfrastructureService: redisInfrastructureService,
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

	size := config.GetConfig().GetGameMapSize()
	newGameRoomId, err := gameRoomApplicationService.CreateGameRoom(size, size)
	if err != nil {
		panic(err.Error())
	}
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
