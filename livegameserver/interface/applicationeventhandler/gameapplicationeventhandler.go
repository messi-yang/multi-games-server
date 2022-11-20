package applicationeventhandler

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/livegameserver/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/messaging/redissubscriber"
	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/presenter/presenterdto"
)

type GameIntegrationEventHandlerConfiguration struct {
	LiveGameApplicationService *applicationservice.LiveGameApplicationService
}

func NewGameIntegrationEventHandler(
	configuration GameIntegrationEventHandlerConfiguration,
	liveGameId livegamemodel.LiveGameId,
) {
	redisReviveUnitsRequestedSubscriber, _ := redissubscriber.NewRedisReviveUnitsRequestedSubscriber()
	redisReviveUnitsRequestedSubscriberUnsubscriber := redisReviveUnitsRequestedSubscriber.Subscribe(func(event redissubscriber.RedisReviveUnitsRequestedIntegrationEvent) {
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

	redisAddPlayerRequestedSubscriber, _ := redissubscriber.NewRedisAddPlayerRequestedSubscriber()
	redisAddPlayerRequestedSubscriberUnsubscriber := redisAddPlayerRequestedSubscriber.Subscribe(func(event redissubscriber.RedisAddPlayerRequestedIntegrationEvent) {
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

	redisRemovePlayerRequestedSubscriber, _ := redissubscriber.NewRedisRemovePlayerRequestedSubscriber()
	redisRemovePlayerRequestedSubscriberUnsubscriber := redisRemovePlayerRequestedSubscriber.Subscribe(func(event redissubscriber.RedisRemovePlayerRequestedIntegrationEvent) {
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

	redisZoomAreaRequestedSubscriber, _ := redissubscriber.NewRedisZoomAreaRequestedSubscriber()
	redisZoomAreaRequestedSubscriberUnsubscriber := redisZoomAreaRequestedSubscriber.Subscribe(func(event redissubscriber.RedisZoomAreaRequestedIntegrationEvent) {
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
