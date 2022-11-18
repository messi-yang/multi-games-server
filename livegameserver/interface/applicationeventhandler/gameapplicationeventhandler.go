package applicationeventhandler

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegameserver/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/messaging/redis"
)

type GameIntegrationEventHandlerConfiguration struct {
	LiveGameApplicationService *applicationservice.LiveGameApplicationService
}

func NewGameIntegrationEventHandler(
	configuration GameIntegrationEventHandlerConfiguration,
	liveGameId livegamemodel.LiveGameId,
) {
	redisReviveUnitsRequestedListener, _ := redis.NewRedisReviveUnitsRequestedListener()
	redisReviveUnitsRequestedListenerUnsubscriber := redisReviveUnitsRequestedListener.Subscribe(func(event redis.RedisReviveUnitsRequestedIntegrationEvent) {
		liveGameId, err := event.LiveGameId.ToValueObject()
		if err != nil {
			return
		}
		configuration.LiveGameApplicationService.ReviveUnitsInLiveGame(liveGameId, event.Coordinates)
	})
	defer redisReviveUnitsRequestedListenerUnsubscriber()

	redisAddPlayerRequestedListener, _ := redis.NewRedisAddPlayerRequestedListener()
	redisAddPlayerRequestedListenerUnsubscriber := redisAddPlayerRequestedListener.Subscribe(func(event redis.RedisAddPlayerRequestedIntegrationEvent) {
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

	redisRemovePlayerRequestedListener, _ := redis.NewRedisRemovePlayerRequestedListener()
	redisRemovePlayerRequestedListenerUnsubscriber := redisRemovePlayerRequestedListener.Subscribe(func(event redis.RedisRemovePlayerRequestedIntegrationEvent) {
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

	redisZoomAreaRequestedListener, _ := redis.NewRedisZoomAreaRequestedListener()
	redisZoomAreaRequestedListenerUnsubscriber := redisZoomAreaRequestedListener.Subscribe(func(event redis.RedisZoomAreaRequestedIntegrationEvent) {
		liveGameId, err := event.LiveGameId.ToValueObject()
		if err != nil {
			return
		}
		playerId, err := event.PlayerId.ToValueObject()
		if err != nil {
			return
		}
		configuration.LiveGameApplicationService.AddZoomedAreaToLiveGame(liveGameId, playerId, event.Area)
	})
	defer redisZoomAreaRequestedListenerUnsubscriber()

	closeConnFlag := make(chan bool)
	for {
		<-closeConnFlag
		return
	}
}
