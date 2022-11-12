package applicationeventhandler

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/common/infrastructure/rediseventbus"
	liveGameValueObject "github.com/dum-dum-genius/game-of-liberty-computer/livegame/domain/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegame/port/adapter/applicationevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegameserver/application/applicationservice"
)

type GameApplicationEventHandlerConfiguration struct {
	LiveGameApplicationService *applicationservice.LiveGameApplicationService
}

func NewGameApplicationEventHandler(
	configuration GameApplicationEventHandlerConfiguration,
	liveGameId liveGameValueObject.LiveGameId,
) {
	reviveUnitsRequestedApplicationEventUnsubscriber := rediseventbus.NewRedisApplicationEventBus(
		rediseventbus.WithRedisInfrastructureService[applicationevent.ReviveUnitsRequestedApplicationEvent](),
	).Subscribe(
		applicationevent.NewReviveUnitsRequestedApplicationEventTopic(liveGameId.GetId()),
		func(event applicationevent.ReviveUnitsRequestedApplicationEvent) {
			configuration.LiveGameApplicationService.ReviveUnitsInLiveGame(liveGameId, event.Coordinates)
		},
	)
	defer reviveUnitsRequestedApplicationEventUnsubscriber()

	addPlayerRequestedApplicationEventUnsubscriber := rediseventbus.NewRedisApplicationEventBus(
		rediseventbus.WithRedisInfrastructureService[applicationevent.AddPlayerRequestedApplicationEvent](),
	).Subscribe(
		applicationevent.NewAddPlayerRequestedApplicationEventTopic(liveGameId.GetId()),
		func(event applicationevent.AddPlayerRequestedApplicationEvent) {
			playerId := event.PlayerId
			configuration.LiveGameApplicationService.AddPlayerToLiveGame(liveGameId, playerId)
		},
	)
	defer addPlayerRequestedApplicationEventUnsubscriber()

	removePlayerRequestedApplicationEventUnsubscriber := rediseventbus.NewRedisApplicationEventBus(
		rediseventbus.WithRedisInfrastructureService[applicationevent.RemovePlayerRequestedApplicationEvent](),
	).Subscribe(
		applicationevent.NewRemovePlayerRequestedApplicationEventTopic(liveGameId.GetId()),
		func(event applicationevent.RemovePlayerRequestedApplicationEvent) {
			configuration.LiveGameApplicationService.RemovePlayerFromLiveGame(liveGameId, event.PlayerId)
			configuration.LiveGameApplicationService.RemoveZoomedAreaFromLiveGame(liveGameId, event.PlayerId)
		},
	)
	defer removePlayerRequestedApplicationEventUnsubscriber()

	zoomAreaRequestedApplicationEventSubscriber := rediseventbus.NewRedisApplicationEventBus(
		rediseventbus.WithRedisInfrastructureService[applicationevent.ZoomAreaRequestedApplicationEvent](),
	).Subscribe(
		applicationevent.NewZoomAreaRequestedApplicationEventTopic(liveGameId.GetId()),
		func(event applicationevent.ZoomAreaRequestedApplicationEvent) {
			configuration.LiveGameApplicationService.AddZoomedAreaToLiveGame(liveGameId, event.PlayerId, event.Area)
		},
	)
	defer zoomAreaRequestedApplicationEventSubscriber()

	closeConnFlag := make(chan bool)
	for {
		<-closeConnFlag
		return
	}
}
