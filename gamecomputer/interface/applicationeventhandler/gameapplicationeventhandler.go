package applicationeventhandler

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/common/infrastructure/rediseventbus"
	"github.com/dum-dum-genius/game-of-liberty-computer/game/port/adapter/applicationevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/application/applicationservice"
	"github.com/google/uuid"
)

type GameApplicationEventHandlerConfiguration struct {
	GameApplicationService *applicationservice.GameApplicationService
}

func NewGameApplicationEventHandler(
	configuration GameApplicationEventHandlerConfiguration,
	gameId uuid.UUID,
) {
	reviveUnitsRequestedApplicationEventUnsubscriber := rediseventbus.NewRedisApplicationEventBus(
		rediseventbus.WithRedisInfrastructureService[applicationevent.ReviveUnitsRequestedApplicationEvent](),
	).Subscribe(
		applicationevent.NewReviveUnitsRequestedApplicationEventTopic(gameId),
		func(event applicationevent.ReviveUnitsRequestedApplicationEvent) {
			configuration.GameApplicationService.ReviveUnitsInGame(gameId, event.Coordinates)
		},
	)
	defer reviveUnitsRequestedApplicationEventUnsubscriber()

	addPlayerRequestedApplicationEventUnsubscriber := rediseventbus.NewRedisApplicationEventBus(
		rediseventbus.WithRedisInfrastructureService[applicationevent.AddPlayerRequestedApplicationEvent](),
	).Subscribe(
		applicationevent.NewAddPlayerRequestedApplicationEventTopic(gameId),
		func(event applicationevent.AddPlayerRequestedApplicationEvent) {
			playerId := event.PlayerId
			configuration.GameApplicationService.AddPlayerToGame(gameId, playerId)
		},
	)
	defer addPlayerRequestedApplicationEventUnsubscriber()

	removePlayerRequestedApplicationEventUnsubscriber := rediseventbus.NewRedisApplicationEventBus(
		rediseventbus.WithRedisInfrastructureService[applicationevent.RemovePlayerRequestedApplicationEvent](),
	).Subscribe(
		applicationevent.NewRemovePlayerRequestedApplicationEventTopic(gameId),
		func(event applicationevent.RemovePlayerRequestedApplicationEvent) {
			configuration.GameApplicationService.RemovePlayerFromGame(gameId, event.PlayerId)
			configuration.GameApplicationService.RemoveZoomedAreaFromGame(gameId, event.PlayerId)
		},
	)
	defer removePlayerRequestedApplicationEventUnsubscriber()

	zoomAreaRequestedApplicationEventSubscriber := rediseventbus.NewRedisApplicationEventBus(
		rediseventbus.WithRedisInfrastructureService[applicationevent.ZoomAreaRequestedApplicationEvent](),
	).Subscribe(
		applicationevent.NewZoomAreaRequestedApplicationEventTopic(gameId),
		func(event applicationevent.ZoomAreaRequestedApplicationEvent) {
			configuration.GameApplicationService.AddZoomedAreaToGame(gameId, event.PlayerId, event.Area)
		},
	)
	defer zoomAreaRequestedApplicationEventSubscriber()

	closeConnFlag := make(chan bool)
	for {
		<-closeConnFlag
		return
	}
}
