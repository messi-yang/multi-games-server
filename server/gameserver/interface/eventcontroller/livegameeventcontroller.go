package eventcontroller

import (
	commonappevent "github.com/dum-dum-genius/game-of-liberty-computer/server/common/application/event"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/gameserver/application/appservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/gameserver/port/adapter/notification/redis"
)

func NewLiveGameEventController(liveGameAppService appservice.LiveGameAppService) {
	redisDestroyItemRequestedSubscriber, _ := redis.NewRedisDestroyItemRequestedSubscriber()
	redisDestroyItemRequestedSubscriberUnsubscriber := redisDestroyItemRequestedSubscriber.Subscribe(func(event *commonappevent.DestroyItemRequestedAppEvent) {
		liveGameId, err := event.GetLiveGameId()
		if err != nil {
			return
		}
		coordinate, err := event.GetCoordinate()
		if err != nil {
			return
		}

		liveGameAppService.DestroyItemInLiveGame(liveGameId, coordinate)
	})
	defer redisDestroyItemRequestedSubscriberUnsubscriber()

	redisBuildItemRequestedSubscriber, _ := redis.NewRedisBuildItemRequestedSubscriber()
	redisBuildItemRequestedSubscriberUnsubscriber := redisBuildItemRequestedSubscriber.Subscribe(func(event *commonappevent.BuildItemRequestedAppEvent) {
		liveGameId, err := event.GetLiveGameId()
		if err != nil {
			return
		}
		coordinate, err := event.GetCoordinate()
		if err != nil {
			return
		}
		itemId, err := event.GetItemId()
		if err != nil {
			return
		}

		liveGameAppService.BuildItemInLiveGame(liveGameId, coordinate, itemId)
	})
	defer redisBuildItemRequestedSubscriberUnsubscriber()

	redisAddPlayerRequestedSubscriber, _ := redis.NewRedisAddPlayerRequestedSubscriber()
	redisAddPlayerRequestedSubscriberUnsubscriber := redisAddPlayerRequestedSubscriber.Subscribe(
		func(event *commonappevent.AddPlayerRequestedAppEvent) {
			liveGameId, err := event.GetLiveGameId()
			if err != nil {
				return
			}
			playerId, err := event.GetPlayerId()
			if err != nil {
				return
			}
			liveGameAppService.AddPlayerToLiveGame(liveGameId, playerId)
		},
	)
	defer redisAddPlayerRequestedSubscriberUnsubscriber()

	redisRemovePlayerRequestedSubscriber, _ := redis.NewRedisRemovePlayerRequestedSubscriber()
	redisRemovePlayerRequestedSubscriberUnsubscriber := redisRemovePlayerRequestedSubscriber.Subscribe(
		func(event *commonappevent.RemovePlayerRequestedAppEvent) {
			liveGameId, err := event.GetLiveGameId()
			if err != nil {
				return
			}
			playerId, err := event.GetPlayerId()
			if err != nil {
				return
			}
			liveGameAppService.RemovePlayerFromLiveGame(liveGameId, playerId)
			liveGameAppService.RemoveZoomedAreaFromLiveGame(liveGameId, playerId)
		},
	)
	defer redisRemovePlayerRequestedSubscriberUnsubscriber()

	redisZoomAreaRequestedSubscriber, _ := redis.NewRedisZoomAreaRequestedSubscriber()
	redisZoomAreaRequestedSubscriberUnsubscriber := redisZoomAreaRequestedSubscriber.Subscribe(
		func(event *commonappevent.ZoomAreaRequestedAppEvent) {
			liveGameId, err := event.GetLiveGameId()
			if err != nil {
				return
			}
			playerId, err := event.GetPlayerId()
			if err != nil {
				return
			}
			area, err := event.GetArea()
			if err != nil {
				return
			}
			liveGameAppService.AddZoomedAreaToLiveGame(liveGameId, playerId, area)
		},
	)
	defer redisZoomAreaRequestedSubscriberUnsubscriber()

	closeConnFlag := make(chan bool)
	for {
		<-closeConnFlag
		return
	}
}
