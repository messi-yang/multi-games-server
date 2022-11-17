package applicationeventhandler

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegameserver/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/integrationeventlistener"
)

type GameIntegrationEventHandlerConfiguration struct {
	LiveGameApplicationService *applicationservice.LiveGameApplicationService
}

func NewGameIntegrationEventHandler(
	configuration GameIntegrationEventHandlerConfiguration,
	liveGameId livegamemodel.LiveGameId,
) {
	reviveUnitsRequestedListener, _ := integrationeventlistener.NewReviveUnitsRequestedListener()
	reviveUnitsRequestedListenerUnsubscriber := reviveUnitsRequestedListener.Subscribe(func(event integrationeventlistener.ReviveUnitsRequestedIntegrationEvent) {
		liveGameId := livegamemodel.NewLiveGameId(event.GameId)
		configuration.LiveGameApplicationService.ReviveUnitsInLiveGame(liveGameId, event.Coordinates)
	})
	defer reviveUnitsRequestedListenerUnsubscriber()

	addPlayerRequestedListener, _ := integrationeventlistener.NewAddPlayerRequestedListener()
	addPlayerRequestedListenerUnsubscriber := addPlayerRequestedListener.Subscribe(func(event integrationeventlistener.AddPlayerRequestedIntegrationEvent) {
		liveGameId := livegamemodel.NewLiveGameId(event.GameId)
		configuration.LiveGameApplicationService.AddPlayerToLiveGame(liveGameId, event.PlayerId)
	})
	defer addPlayerRequestedListenerUnsubscriber()

	removePlayerRequestedListener, _ := integrationeventlistener.NewRemovePlayerRequestedListener()
	removePlayerRequestedListenerUnsubscriber := removePlayerRequestedListener.Subscribe(func(event integrationeventlistener.RemovePlayerRequestedIntegrationEvent) {
		liveGameId := livegamemodel.NewLiveGameId(event.GameId)
		configuration.LiveGameApplicationService.RemovePlayerFromLiveGame(liveGameId, event.PlayerId)
		configuration.LiveGameApplicationService.RemoveZoomedAreaFromLiveGame(liveGameId, event.PlayerId)
	})
	defer removePlayerRequestedListenerUnsubscriber()

	zoomAreaRequestedListener, _ := integrationeventlistener.NewZoomAreaRequestedListener()
	zoomAreaRequestedListenerUnsubscriber := zoomAreaRequestedListener.Subscribe(func(event integrationeventlistener.ZoomAreaRequestedIntegrationEvent) {
		liveGameId := livegamemodel.NewLiveGameId(event.GameId)
		configuration.LiveGameApplicationService.AddZoomedAreaToLiveGame(liveGameId, event.PlayerId, event.Area)
	})
	defer zoomAreaRequestedListenerUnsubscriber()

	closeConnFlag := make(chan bool)
	for {
		<-closeConnFlag
		return
	}
}
