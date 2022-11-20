package applicationeventhandler

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegameserver/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/messaging/redislistener"
	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/presenter/presenterdto"
)

type GameIntegrationEventHandlerConfiguration struct {
	LiveGameApplicationService *applicationservice.LiveGameApplicationService
}

func NewGameIntegrationEventHandler(
	configuration GameIntegrationEventHandlerConfiguration,
	liveGameId livegamemodel.LiveGameId,
) {
	redisReviveUnitsRequestedListener, _ := redislistener.NewRedisReviveUnitsRequestedListener()
	redisReviveUnitsRequestedListenerUnsubscriber := redisReviveUnitsRequestedListener.Subscribe(func(event redislistener.RedisReviveUnitsRequestedIntegrationEvent) {
		liveGameId, err := event.LiveGameId.ToValueObject()
		if err != nil {
			return
		}
		coordinates, err := presenterdto.ParseCoordinatePresenterDtos(event.Coordinates)
		if err != nil {
			return
		}
		configuration.LiveGameApplicationService.ReviveUnitsInLiveGame(liveGameId, coordinates)
	})
	defer redisReviveUnitsRequestedListenerUnsubscriber()

	redisAddPlayerRequestedListener, _ := redislistener.NewRedisAddPlayerRequestedListener()
	redisAddPlayerRequestedListenerUnsubscriber := redisAddPlayerRequestedListener.Subscribe(func(event redislistener.RedisAddPlayerRequestedIntegrationEvent) {
		liveGameId, err := event.LiveGameId.ToValueObject()
		if err != nil {
			return
		}
		playerId, err := event.PlayerId.ToValueObject()
		if err != nil {
			return
		}
		configuration.LiveGameApplicationService.AddPlayerToLiveGame(liveGameId, playerId)
	})
	defer redisAddPlayerRequestedListenerUnsubscriber()

	redisRemovePlayerRequestedListener, _ := redislistener.NewRedisRemovePlayerRequestedListener()
	redisRemovePlayerRequestedListenerUnsubscriber := redisRemovePlayerRequestedListener.Subscribe(func(event redislistener.RedisRemovePlayerRequestedIntegrationEvent) {
		liveGameId, err := event.LiveGameId.ToValueObject()
		if err != nil {
			return
		}
		playerId, err := event.PlayerId.ToValueObject()
		if err != nil {
			return
		}
		configuration.LiveGameApplicationService.RemovePlayerFromLiveGame(liveGameId, playerId)
		configuration.LiveGameApplicationService.RemoveZoomedAreaFromLiveGame(liveGameId, playerId)
	})
	defer redisRemovePlayerRequestedListenerUnsubscriber()

	redisZoomAreaRequestedListener, _ := redislistener.NewRedisZoomAreaRequestedListener()
	redisZoomAreaRequestedListenerUnsubscriber := redisZoomAreaRequestedListener.Subscribe(func(event redislistener.RedisZoomAreaRequestedIntegrationEvent) {
		liveGameId, err := event.LiveGameId.ToValueObject()
		if err != nil {
			return
		}
		playerId, err := event.PlayerId.ToValueObject()
		if err != nil {
			return
		}
		area, err := event.Area.ToValueObject()
		if err != nil {
			return
		}
		configuration.LiveGameApplicationService.AddZoomedAreaToLiveGame(liveGameId, playerId, area)
	})
	defer redisZoomAreaRequestedListenerUnsubscriber()

	closeConnFlag := make(chan bool)
	for {
		<-closeConnFlag
		return
	}
}
