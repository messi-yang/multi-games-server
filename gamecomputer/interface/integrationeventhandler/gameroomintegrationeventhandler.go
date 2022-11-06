package integrationeventhandler

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/application/eventbus"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/integrationevent"
	"github.com/google/uuid"
)

type GameIntegrationEventHandlerConfiguration struct {
	IntegrationEventBus    eventbus.IntegrationEventBus
	GameApplicationService *applicationservice.GameApplicationService
}

func NewGameIntegrationEventHandler(
	configuration GameIntegrationEventHandlerConfiguration,
	gameId uuid.UUID,
) {
	reviveUnitsRequestedIntegrationEventUnsubscriber := configuration.IntegrationEventBus.Subscribe(
		integrationevent.NewReviveUnitsRequestedIntegrationEventTopic(gameId),
		func(event []byte) {
			var reviveUnitsRequestedIntegrationEvent integrationevent.ReviveUnitsRequestedIntegrationEvent
			json.Unmarshal(event, &reviveUnitsRequestedIntegrationEvent)

			configuration.GameApplicationService.ReviveUnitsInGame(gameId, reviveUnitsRequestedIntegrationEvent.Payload.Coordinates)
		},
	)
	defer reviveUnitsRequestedIntegrationEventUnsubscriber()

	addPlayerRequestedIntegrationEventUnsubscriber := configuration.IntegrationEventBus.Subscribe(
		integrationevent.NewAddPlayerRequestedIntegrationEventTopic(gameId),
		func(event []byte) {
			var addPlayerRequestedIntegrationEvent integrationevent.AddPlayerRequestedIntegrationEvent
			json.Unmarshal(event, &addPlayerRequestedIntegrationEvent)

			playerId := addPlayerRequestedIntegrationEvent.Payload.PlayerId
			configuration.GameApplicationService.AddPlayerToGame(gameId, playerId)
		},
	)
	defer addPlayerRequestedIntegrationEventUnsubscriber()

	removePlayerRequestedIntegrationEventUnsubscriber := configuration.IntegrationEventBus.Subscribe(
		integrationevent.NewRemovePlayerRequestedIntegrationEventTopic(gameId),
		func(event []byte) {
			var removePlayerRequestedIntegrationEvent integrationevent.RemovePlayerRequestedIntegrationEvent
			json.Unmarshal(event, &removePlayerRequestedIntegrationEvent)

			configuration.GameApplicationService.RemovePlayerFromGame(gameId, removePlayerRequestedIntegrationEvent.Payload.PlayerId)
			configuration.GameApplicationService.RemoveZoomedAreaFromGame(gameId, removePlayerRequestedIntegrationEvent.Payload.PlayerId)
		},
	)
	defer removePlayerRequestedIntegrationEventUnsubscriber()

	zoomAreaRequestedIntegrationEventSubscriber := configuration.IntegrationEventBus.Subscribe(
		integrationevent.NewZoomAreaRequestedIntegrationEventTopic(gameId),
		func(event []byte) {
			var zoomAreaRequestedIntegrationEvent integrationevent.ZoomAreaRequestedIntegrationEvent
			json.Unmarshal(event, &zoomAreaRequestedIntegrationEvent)

			configuration.GameApplicationService.AddZoomedAreaToGame(gameId, zoomAreaRequestedIntegrationEvent.Payload.PlayerId, zoomAreaRequestedIntegrationEvent.Payload.Area)
		},
	)
	defer zoomAreaRequestedIntegrationEventSubscriber()

	closeConnFlag := make(chan bool)
	for {
		<-closeConnFlag
		return
	}
}
