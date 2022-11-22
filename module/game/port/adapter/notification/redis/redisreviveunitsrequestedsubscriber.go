package redis

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/module/common/notification"
	commonredis "github.com/dum-dum-genius/game-of-liberty-computer/module/common/port/adapter/notification/redis"
	"github.com/dum-dum-genius/game-of-liberty-computer/module/game/domain/model/gamecommonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/module/game/domain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/module/game/port/adapter/presenter/presenterdto"
)

type RedisReviveUnitsRequestedIntegrationEvent struct {
	LiveGameId  presenterdto.LiveGameIdPresenterDto   `json:"liveGameId"`
	Coordinates []presenterdto.CoordinatePresenterDto `json:"coordinates"`
}

func NewRedisReviveUnitsRequestedIntegrationEvent(liveGameId livegamemodel.LiveGameId, coordinates []gamecommonmodel.Coordinate) RedisReviveUnitsRequestedIntegrationEvent {
	return RedisReviveUnitsRequestedIntegrationEvent{
		LiveGameId:  presenterdto.NewLiveGameIdPresenterDto(liveGameId),
		Coordinates: presenterdto.NewCoordinatePresenterDtos(coordinates),
	}
}

var RedisReviveUnitsRequestedSubscriberChannel string = "revive-units-requested"

type RedisReviveUnitsRequestedSubscriber struct {
	redisProvider *commonredis.RedisProvider
}

func NewRedisReviveUnitsRequestedSubscriber() (notification.NotificationSubscriber[RedisReviveUnitsRequestedIntegrationEvent], error) {
	return &RedisReviveUnitsRequestedSubscriber{
		redisProvider: commonredis.NewRedisProvider(),
	}, nil
}

func (subscriber *RedisReviveUnitsRequestedSubscriber) Subscribe(handler func(RedisReviveUnitsRequestedIntegrationEvent)) func() {
	unsubscriber := subscriber.redisProvider.Subscribe(RedisReviveUnitsRequestedSubscriberChannel, func(message []byte) {
		var event RedisReviveUnitsRequestedIntegrationEvent
		json.Unmarshal(message, &event)
		handler(event)
	})

	return func() {
		unsubscriber()
	}
}
