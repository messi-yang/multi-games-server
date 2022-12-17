package redis

import (
	commonappevent "github.com/dum-dum-genius/game-of-liberty-computer/server/common/application/event"
	commonnotification "github.com/dum-dum-genius/game-of-liberty-computer/server/common/application/notification"
	commonredis "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification/redis"
)

type RedisReviveUnitsRequestedSubscriber struct {
	redisProvider *commonredis.RedisProvider
}

func NewRedisReviveUnitsRequestedSubscriber() (commonnotification.NotificationSubscriber[*commonappevent.ReviveUnitsRequestedAppEvent], error) {
	return &RedisReviveUnitsRequestedSubscriber{
		redisProvider: commonredis.NewRedisProvider(),
	}, nil
}

func (subscriber *RedisReviveUnitsRequestedSubscriber) Subscribe(handler func(*commonappevent.ReviveUnitsRequestedAppEvent)) func() {
	unsubscriber := subscriber.redisProvider.Subscribe(
		commonappevent.NewReviveUnitsRequestedAppEventChannel(),
		func(message []byte) {
			event := commonappevent.DeserializeReviveUnitsRequestedAppEvent(message)
			handler(&event)
		})

	return func() {
		unsubscriber()
	}
}
