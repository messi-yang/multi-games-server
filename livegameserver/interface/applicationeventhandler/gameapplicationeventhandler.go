package applicationeventhandler

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/common/infrastructure/rediseventbus"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegame/domain/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegame/port/adapter/applicationevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegameserver/application/applicationservice"
)

type GameApplicationEventHandlerConfiguration struct {
	LiveGameApplicationService *applicationservice.LiveGameApplicationService
}

func NewGameApplicationEventHandler(
	configuration GameApplicationEventHandlerConfiguration,
	gameId valueobject.GameId,
) {
	reviveUnitsRequestedApplicationEventUnsubscriber := rediseventbus.NewRedisApplicationEventBus(
		rediseventbus.WithRedisInfrastructureService[applicationevent.ReviveUnitsRequestedApplicationEvent](),
	).Subscribe(
		applicationevent.NewReviveUnitsRequestedApplicationEventTopic(gameId.GetId()),
		func(event applicationevent.ReviveUnitsRequestedApplicationEvent) {
			configuration.LiveGameApplicationService.ReviveUnitsInLiveGame(gameId, event.Coordinates)
		},
	)
	defer reviveUnitsRequestedApplicationEventUnsubscriber()

	addPlayerRequestedApplicationEventUnsubscriber := rediseventbus.NewRedisApplicationEventBus(
		rediseventbus.WithRedisInfrastructureService[applicationevent.AddPlayerRequestedApplicationEvent](),
	).Subscribe(
		applicationevent.NewAddPlayerRequestedApplicationEventTopic(gameId.GetId()),
		func(event applicationevent.AddPlayerRequestedApplicationEvent) {
			playerId := event.PlayerId
			configuration.LiveGameApplicationService.AddPlayerToLiveGame(gameId, playerId)
		},
	)
	defer addPlayerRequestedApplicationEventUnsubscriber()

	removePlayerRequestedApplicationEventUnsubscriber := rediseventbus.NewRedisApplicationEventBus(
		rediseventbus.WithRedisInfrastructureService[applicationevent.RemovePlayerRequestedApplicationEvent](),
	).Subscribe(
		applicationevent.NewRemovePlayerRequestedApplicationEventTopic(gameId.GetId()),
		func(event applicationevent.RemovePlayerRequestedApplicationEvent) {
			configuration.LiveGameApplicationService.RemovePlayerFromLiveGame(gameId, event.PlayerId)
			configuration.LiveGameApplicationService.RemoveZoomedAreaFromLiveGame(gameId, event.PlayerId)
		},
	)
	defer removePlayerRequestedApplicationEventUnsubscriber()

	zoomAreaRequestedApplicationEventSubscriber := rediseventbus.NewRedisApplicationEventBus(
		rediseventbus.WithRedisInfrastructureService[applicationevent.ZoomAreaRequestedApplicationEvent](),
	).Subscribe(
		applicationevent.NewZoomAreaRequestedApplicationEventTopic(gameId.GetId()),
		func(event applicationevent.ZoomAreaRequestedApplicationEvent) {
			configuration.LiveGameApplicationService.AddZoomedAreaToLiveGame(gameId, event.PlayerId, event.Area)
		},
	)
	defer zoomAreaRequestedApplicationEventSubscriber()

	closeConnFlag := make(chan bool)
	for {
		<-closeConnFlag
		return
	}
}
