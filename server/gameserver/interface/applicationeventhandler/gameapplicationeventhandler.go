package applicationeventhandler

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/common/dto/jsondto"
	applicationservice "github.com/dum-dum-genius/game-of-liberty-computer/server/gameserver/application/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/gameserver/port/adapter/notification/gameredis"
)

type GameIntegrationEventHandlerConfiguration struct {
	LiveGameApplicationService applicationservice.LiveGameApplicationService
}

func NewGameIntegrationEventHandler(
	configuration GameIntegrationEventHandlerConfiguration,
) {
	redisReviveUnitsRequestedSubscriber, _ := gameredis.NewRedisReviveUnitsRequestedSubscriber()
	redisReviveUnitsRequestedSubscriberUnsubscriber := redisReviveUnitsRequestedSubscriber.Subscribe(func(event gameredis.RedisReviveUnitsRequestedIntegrationEvent) {
		liveGameId, err := event.LiveGameId.ToValueObject()
		if err != nil {
			return
		}
		coordinates, err := jsondto.ParseCoordinateJsonDtos(event.Coordinates)
		if err != nil {
			return
		}
		configuration.LiveGameApplicationService.ReviveUnitsInLiveGame(liveGameId, coordinates)
	})
	defer redisReviveUnitsRequestedSubscriberUnsubscriber()

	redisAddPlayerRequestedSubscriber, _ := gameredis.NewRedisAddPlayerRequestedSubscriber()
	redisAddPlayerRequestedSubscriberUnsubscriber := redisAddPlayerRequestedSubscriber.Subscribe(func(event gameredis.RedisAddPlayerRequestedIntegrationEvent) {
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

	redisRemovePlayerRequestedSubscriber, _ := gameredis.NewRedisRemovePlayerRequestedSubscriber()
	redisRemovePlayerRequestedSubscriberUnsubscriber := redisRemovePlayerRequestedSubscriber.Subscribe(func(event gameredis.RedisRemovePlayerRequestedIntegrationEvent) {
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

	redisZoomAreaRequestedSubscriber, _ := gameredis.NewRedisZoomAreaRequestedSubscriber()
	redisZoomAreaRequestedSubscriberUnsubscriber := redisZoomAreaRequestedSubscriber.Subscribe(func(event gameredis.RedisZoomAreaRequestedIntegrationEvent) {
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
