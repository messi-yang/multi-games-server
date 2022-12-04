package redis

import (
	commonapplicationevent "github.com/dum-dum-genius/game-of-liberty-computer/server/common/application/event"
	commonnotification "github.com/dum-dum-genius/game-of-liberty-computer/server/common/application/notification"
	commonredis "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification/redis"
)

type RedisReviveUnitsRequestedSubscriber struct {
	redisProvider *commonredis.RedisProvider
}

func NewRedisReviveUnitsRequestedSubscriber() (commonnotification.NotificationSubscriber[*commonapplicationevent.ReviveUnitsRequestedApplicationEvent], error) {
	return &RedisReviveUnitsRequestedSubscriber{
		redisProvider: commonredis.NewRedisProvider(),
	}, nil
}

func (subscriber *RedisReviveUnitsRequestedSubscriber) Subscribe(handler func(*commonapplicationevent.ReviveUnitsRequestedApplicationEvent)) func() {
	unsubscriber := subscriber.redisProvider.Subscribe(
		commonapplicationevent.NewReviveUnitsRequestedApplicationEventChannel(),
		func(message []byte) {
			event := commonapplicationevent.DeserializeReviveUnitsRequestedApplicationEvent(message)
			handler(&event)
		})

	return func() {
		unsubscriber()
	}
}
