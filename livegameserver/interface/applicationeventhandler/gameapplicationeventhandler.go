package applicationeventhandler

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegameserver/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/notification/redis"
	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/presenter/presenterdto"
)

type GameIntegrationEventHandlerConfiguration struct {
	LiveGameApplicationService *applicationservice.LiveGameApplicationService
}

func NewGameIntegrationEventHandler(
	configuration GameIntegrationEventHandlerConfiguration,
	liveGameId livegamemodel.LiveGameId,
) {
	redisReviveUnitsRequestedSubscriber, _ := redis.NewRedisReviveUnitsRequestedSubscriber()
	redisReviveUnitsRequestedSubscriberUnsubscriber := redisReviveUnitsRequestedSubscriber.Subscribe(func(event redis.RedisReviveUnitsRequestedIntegrationEvent) {
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
	defer redisReviveUnitsRequestedSubscriberUnsubscriber()

	redisAddPlayerRequestedSubscriber, _ := redis.NewRedisAddPlayerRequestedSubscriber()
	redisAddPlayerRequestedSubscriberUnsubscriber := redisAddPlayerRequestedSubscriber.Subscribe(func(event redis.RedisAddPlayerRequestedIntegrationEvent) {
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
	defer redisAddPlayerRequestedSubscriberUnsubscriber()

	redisRemovePlayerRequestedSubscriber, _ := redis.NewRedisRemovePlayerRequestedSubscriber()
	redisRemovePlayerRequestedSubscriberUnsubscriber := redisRemovePlayerRequestedSubscriber.Subscribe(func(event redis.RedisRemovePlayerRequestedIntegrationEvent) {
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
	defer redisRemovePlayerRequestedSubscriberUnsubscriber()

	redisZoomAreaRequestedSubscriber, _ := redis.NewRedisZoomAreaRequestedSubscriber()
	redisZoomAreaRequestedSubscriberUnsubscriber := redisZoomAreaRequestedSubscriber.Subscribe(func(event redis.RedisZoomAreaRequestedIntegrationEvent) {
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
	defer redisZoomAreaRequestedSubscriberUnsubscriber()

	closeConnFlag := make(chan bool)
	for {
		<-closeConnFlag
		return
	}
}
