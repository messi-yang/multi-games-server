package redis

import (
	commonappevent "github.com/dum-dum-genius/game-of-liberty-computer/server/common/application/event"
	commonnotification "github.com/dum-dum-genius/game-of-liberty-computer/server/common/application/notification"
	commonredis "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification/redis"
)

type RedisDestroyItemRequestedSubscriber struct {
	redisProvider *commonredis.RedisProvider
}

func NewRedisDestroyItemRequestedSubscriber() (commonnotification.NotificationSubscriber[*commonappevent.DestroyItemRequestedAppEvent], error) {
	return &RedisDestroyItemRequestedSubscriber{
		redisProvider: commonredis.NewRedisProvider(),
	}, nil
}

func (subscriber *RedisDestroyItemRequestedSubscriber) Subscribe(handler func(*commonappevent.DestroyItemRequestedAppEvent)) func() {
	unsubscriber := subscriber.redisProvider.Subscribe(
		commonappevent.NewDestroyItemRequestedAppEventChannel(),
		func(message []byte) {
			event := commonappevent.DeserializeDestroyItemRequestedAppEvent(message)
			handler(&event)
		})

	return func() {
		unsubscriber()
	}
}
