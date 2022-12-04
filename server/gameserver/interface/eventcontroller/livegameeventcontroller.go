package eventcontroller

import (
	commonjsondto "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/common/dto/jsondto"
	commonredisdto "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification/redis/dto"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/gameserver/application/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/gameserver/port/adapter/notification/redis"
)

type LiveGameEventControllerConfiguration struct {
	LiveGameApplicationService service.LiveGameApplicationService
}

func NewLiveGameEventController(
	configuration LiveGameEventControllerConfiguration,
) {
	redisReviveUnitsRequestedSubscriber, _ := redis.NewRedisReviveUnitsRequestedSubscriber()
	redisReviveUnitsRequestedSubscriberUnsubscriber := redisReviveUnitsRequestedSubscriber.Subscribe(
		func(event commonredisdto.RedisReviveUnitsRequestedEvent) {
			liveGameId, err := event.LiveGameId.ToValueObject()
			if err != nil {
				return
			}
			coordinates, err := commonjsondto.ParseCoordinateJsonDtos(event.Coordinates)
			if err != nil {
				return
			}
			configuration.LiveGameApplicationService.ReviveUnitsInLiveGame(liveGameId, coordinates)
		},
	)
	defer redisReviveUnitsRequestedSubscriberUnsubscriber()

	redisAddPlayerRequestedSubscriber, _ := redis.NewRedisAddPlayerRequestedSubscriber()
	redisAddPlayerRequestedSubscriberUnsubscriber := redisAddPlayerRequestedSubscriber.Subscribe(
		func(event commonredisdto.RedisAddPlayerRequestedEvent) {
			liveGameId, err := event.LiveGameId.ToValueObject()
			if err != nil {
				return
			}
			playerId, err := event.PlayerId.ToValueObject()
			if err != nil {
				return
			}
			configuration.LiveGameApplicationService.AddPlayerToLiveGame(liveGameId, playerId)
		},
	)
	defer redisAddPlayerRequestedSubscriberUnsubscriber()

	redisRemovePlayerRequestedSubscriber, _ := redis.NewRedisRemovePlayerRequestedSubscriber()
	redisRemovePlayerRequestedSubscriberUnsubscriber := redisRemovePlayerRequestedSubscriber.Subscribe(
		func(event commonredisdto.RedisRemovePlayerRequestedEvent) {
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
		},
	)
	defer redisRemovePlayerRequestedSubscriberUnsubscriber()

	redisZoomAreaRequestedSubscriber, _ := redis.NewRedisZoomAreaRequestedSubscriber()
	redisZoomAreaRequestedSubscriberUnsubscriber := redisZoomAreaRequestedSubscriber.Subscribe(
		func(event commonredisdto.RedisZoomAreaRequestedEvent) {
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
		},
	)
	defer redisZoomAreaRequestedSubscriberUnsubscriber()

	closeConnFlag := make(chan bool)
	for {
		<-closeConnFlag
		return
	}
}
