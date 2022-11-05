package integrationeventhandler

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/application/eventbus"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/dto"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/integrationevent"
	"github.com/google/uuid"
)

type GameIntegrationEventHandlerConfiguration struct {
	IntegrationEventBus    eventbus.IntegrationEventBus
	GameApplicationService applicationservice.GameApplicationService
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

			coordinates, err := dto.ParseCoordinateDtos(reviveUnitsRequestedIntegrationEvent.Payload.Coordinates)
			if err != nil {
				return
			}

			configuration.GameApplicationService.ReviveUnitsInGame(gameId, coordinates)
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

			area, err := zoomAreaRequestedIntegrationEvent.Payload.Area.ToValueObject()
			if err != nil {
				return
			}
			configuration.GameApplicationService.AddZoomedAreaToGame(gameId, zoomAreaRequestedIntegrationEvent.Payload.PlayerId, area)
		},
	)
	defer zoomAreaRequestedIntegrationEventSubscriber()

	closeConnFlag := make(chan bool)
	for {
		<-closeConnFlag
		return
	}
}
