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
		liveGameId := livegamemodel.NewLiveGameId(event.GameId)
		configuration.LiveGameApplicationService.ReviveUnitsInLiveGame(liveGameId, event.Coordinates)
	})
	defer redisReviveUnitsRequestedListenerUnsubscriber()

	redisAddPlayerRequestedListener, _ := redis.NewRedisAddPlayerRequestedListener()
	redisAddPlayerRequestedListenerUnsubscriber := redisAddPlayerRequestedListener.Subscribe(func(event redis.RedisAddPlayerRequestedIntegrationEvent) {
		liveGameId := livegamemodel.NewLiveGameId(event.GameId)
		configuration.LiveGameApplicationService.AddPlayerToLiveGame(liveGameId, event.PlayerId)
	})
	defer redisAddPlayerRequestedListenerUnsubscriber()

	redisRemovePlayerRequestedListener, _ := redis.NewRedisRemovePlayerRequestedListener()
	redisRemovePlayerRequestedListenerUnsubscriber := redisRemovePlayerRequestedListener.Subscribe(func(event redis.RedisRemovePlayerRequestedIntegrationEvent) {
		liveGameId := livegamemodel.NewLiveGameId(event.GameId)
		configuration.LiveGameApplicationService.RemovePlayerFromLiveGame(liveGameId, event.PlayerId)
		configuration.LiveGameApplicationService.RemoveZoomedAreaFromLiveGame(liveGameId, event.PlayerId)
	})
	defer redisRemovePlayerRequestedListenerUnsubscriber()

	redisZoomAreaRequestedListener, _ := redis.NewRedisZoomAreaRequestedListener()
	redisZoomAreaRequestedListenerUnsubscriber := redisZoomAreaRequestedListener.Subscribe(func(event redis.RedisZoomAreaRequestedIntegrationEvent) {
		liveGameId := livegamemodel.NewLiveGameId(event.GameId)
		configuration.LiveGameApplicationService.AddZoomedAreaToLiveGame(liveGameId, event.PlayerId, event.Area)
	})
	defer redisZoomAreaRequestedListenerUnsubscriber()

	closeConnFlag := make(chan bool)
	for {
		<-closeConnFlag
		return
	}
}
