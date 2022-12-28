package redis

import (
	commonredis "github.com/dum-dum-genius/game-of-liberty-computer/server/common/adapter/notification/redis"
	commonappevent "github.com/dum-dum-genius/game-of-liberty-computer/server/common/application/event"
	commonnotification "github.com/dum-dum-genius/game-of-liberty-computer/server/common/application/notification"
)

type RedisZoomAreaRequestedSubscriber struct {
	redisProvider *commonredis.RedisProvider
}

func NewRedisZoomAreaRequestedSubscriber() (commonnotification.NotificationSubscriber[*commonappevent.ZoomAreaRequestedAppEvent], error) {
	return &RedisZoomAreaRequestedSubscriber{
		redisProvider: commonredis.NewRedisProvider(),
	}, nil
}

func (subscriber *RedisZoomAreaRequestedSubscriber) Subscribe(handler func(*commonappevent.ZoomAreaRequestedAppEvent)) func() {
	unsubscriber := subscriber.redisProvider.Subscribe(
		commonappevent.NewZoomAreaRequestedAppEventChannel(),
		func(message []byte) {
			event := commonappevent.DeserializeZoomAreaRequestedAppEvent(message)
			handler(&event)
		})

	return func() {
		unsubscriber()
	}
}
