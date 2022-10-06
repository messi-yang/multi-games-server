package integrationeventhandler

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/infrastructure/repositorymemory"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/eventbusredis"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/infrastructureservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/dto"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/integrationevent"
	"github.com/google/uuid"
)

func HandleGameRoomIntegrationEvent() {
	redisInfrastructureService := infrastructureservice.NewRedisInfrastructureService()
	gameIdBytes, err := redisInfrastructureService.Get("game_id")
	if err != nil {
		return
	}
	gameId, err := uuid.Parse(string(gameIdBytes))
	if err != nil {
		return
	}

	gameRoomRepositoryMemory := repositorymemory.NewGameRoomRepositoryMemory()
	integrationEventBusRedis := eventbusredis.NewIntegrationEventBusRedis(
		eventbusredis.IntegrationEventBusRedisCallbackConfiguration{
			RedisInfrastructureService: infrastructureservice.NewRedisInfrastructureService(),
		},
	)
	gameRoomApplicationService := applicationservice.NewGameRoomApplicationService(
		applicationservice.GameRoomApplicationServiceConfiguration{
			GameRoomRepository:  gameRoomRepositoryMemory,
			IntegrationEventBus: integrationEventBusRedis,
		},
	)

	reviveUnitsRequestedIntegrationEventUnsubscriber := integrationEventBusRedis.Subscribe(
		integrationevent.NewReviveUnitsRequestedIntegrationEventTopic(gameId),
		func(event []byte) {
			var reviveUnitsRequestedIntegrationEvent integrationevent.ReviveUnitsRequestedIntegrationEvent
			json.Unmarshal(event, &reviveUnitsRequestedIntegrationEvent)

			coordinates, err := dto.ParseCoordinateDtos(reviveUnitsRequestedIntegrationEvent.Payload.Coordinates)
			if err != nil {
				return
			}

			gameRoomApplicationService.ReviveUnits(gameId, coordinates)
		},
	)
	defer reviveUnitsRequestedIntegrationEventUnsubscriber()

	addPlayerRequestedIntegrationEventUnsubscriber := integrationEventBusRedis.Subscribe(
		integrationevent.NewAddPlayerRequestedIntegrationEventTopic(gameId),
		func(event []byte) {
			var addPlayerRequestedIntegrationEvent integrationevent.AddPlayerRequestedIntegrationEvent
			json.Unmarshal(event, &addPlayerRequestedIntegrationEvent)

			player := addPlayerRequestedIntegrationEvent.Payload.Player.ToEntity()
			gameRoomApplicationService.AddPlayer(gameId, player)
		},
	)
	defer addPlayerRequestedIntegrationEventUnsubscriber()

	removePlayerRequestedIntegrationEventUnsubscriber := integrationEventBusRedis.Subscribe(
		integrationevent.NewRemovePlayerRequestedIntegrationEventTopic(gameId),
		func(event []byte) {
			var removePlayerRequestedIntegrationEvent integrationevent.RemovePlayerRequestedIntegrationEvent
			json.Unmarshal(event, &removePlayerRequestedIntegrationEvent)

			gameRoomApplicationService.RemovePlayer(gameId, removePlayerRequestedIntegrationEvent.Payload.PlayerId)
			gameRoomApplicationService.RemoveZoomedArea(gameId, removePlayerRequestedIntegrationEvent.Payload.PlayerId)
		},
	)
	defer removePlayerRequestedIntegrationEventUnsubscriber()

	zoomAreaRequestedIntegrationEventSubscriber := integrationEventBusRedis.Subscribe(
		integrationevent.NewZoomAreaRequestedIntegrationEventTopic(gameId),
		func(event []byte) {
			var zoomAreaRequestedIntegrationEvent integrationevent.ZoomAreaRequestedIntegrationEvent
			json.Unmarshal(event, &zoomAreaRequestedIntegrationEvent)

			area, err := zoomAreaRequestedIntegrationEvent.Payload.Area.ToValueObject()
			if err != nil {
				return
			}
			gameRoomApplicationService.AddZoomedArea(gameId, zoomAreaRequestedIntegrationEvent.Payload.PlayerId, area)
		},
	)
	defer zoomAreaRequestedIntegrationEventSubscriber()

	closeConnFlag := make(chan bool)
	for {
		<-closeConnFlag
		return
	}
}
