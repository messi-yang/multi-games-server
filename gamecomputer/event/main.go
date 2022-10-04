package event

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/infrastructure/repositorymemory"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/config"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/eventbusredis"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/dto/areadto"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/dto/coordinatedto"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/dto/playerdto"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/event/addplayerrequestedevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/event/removeplayerrequestedevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/event/reviveunitsrequestedevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/event/zoomarearequestedevent"
)

func Controller() {
	gameId := config.GetConfig().GetGameId()

	gameRoomRepositoryMemory := repositorymemory.GetGameRoomRepositoryMemory()
	integrationEventBusRedis := eventbusredis.GetIntegrationEventBusRedis()
	gameRoomApplicationService := applicationservice.NewGameRoomApplicationService(
		applicationservice.GameRoomApplicationServiceConfiguration{
			GameRoomRepository:  gameRoomRepositoryMemory,
			IntegrationEventBus: integrationEventBusRedis,
		},
	)

	reviveUnitsRequestedEventUnsubscriber := integrationEventBusRedis.Subscribe(
		reviveunitsrequestedevent.NewEventTopic(gameId),
		func(event []byte) {
			var reviveUnitsRequestedEvent reviveunitsrequestedevent.Event
			json.Unmarshal(event, &reviveUnitsRequestedEvent)

			coordinates, err := coordinatedto.FromDtoList(reviveUnitsRequestedEvent.Payload.Coordinates)
			if err != nil {
				return
			}

			gameRoomApplicationService.ReviveUnits(gameId, coordinates)
		},
	)
	defer reviveUnitsRequestedEventUnsubscriber()

	addPlayerRequestedEventUnsubscriber := integrationEventBusRedis.Subscribe(
		addplayerrequestedevent.NewEventTopic(gameId),
		func(event []byte) {
			var addPlayerRequestedEvent addplayerrequestedevent.Event
			json.Unmarshal(event, &addPlayerRequestedEvent)

			player := playerdto.FromDto(addPlayerRequestedEvent.Payload.Player)
			gameRoomApplicationService.AddPlayer(gameId, player)
		},
	)
	defer addPlayerRequestedEventUnsubscriber()

	removePlayerRequestedEventUnsubscriber := integrationEventBusRedis.Subscribe(
		removeplayerrequestedevent.NewEventTopic(gameId),
		func(event []byte) {
			var removePlayerRequestedEvent removeplayerrequestedevent.Event
			json.Unmarshal(event, &removePlayerRequestedEvent)

			gameRoomApplicationService.RemovePlayer(gameId, removePlayerRequestedEvent.Payload.PlayerId)
			gameRoomApplicationService.RemoveZoomedArea(gameId, removePlayerRequestedEvent.Payload.PlayerId)
		},
	)
	defer removePlayerRequestedEventUnsubscriber()

	zoomAreaRequestedEventSubscriber := integrationEventBusRedis.Subscribe(
		zoomarearequestedevent.NewEventTopic(gameId),
		func(event []byte) {
			var zoomAreaRequestedEvent zoomarearequestedevent.Event
			json.Unmarshal(event, &zoomAreaRequestedEvent)

			area, err := areadto.FromDto(zoomAreaRequestedEvent.Payload.Area)
			if err != nil {
				return
			}
			gameRoomApplicationService.AddZoomedArea(gameId, zoomAreaRequestedEvent.Payload.PlayerId, area)
		},
	)
	defer zoomAreaRequestedEventSubscriber()

	closeConnFlag := make(chan bool)
	for {
		<-closeConnFlag
		return
	}
}
