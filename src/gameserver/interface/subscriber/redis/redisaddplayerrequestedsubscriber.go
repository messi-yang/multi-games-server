package redis

import (
	commonredis "github.com/dum-dum-genius/game-of-liberty-computer/src/common/adapter/notification/redis"
	commonappevent "github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/event"
	commonnotification "github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/notification"
)

type RedisAddPlayerRequestedSubscriber struct {
	redisProvider *commonredis.RedisProvider
}

func NewRedisAddPlayerRequestedSubscriber() (commonnotification.NotificationSubscriber[*commonappevent.AddPlayerRequestedAppEvent], error) {
	return &RedisAddPlayerRequestedSubscriber{
		redisProvider: commonredis.NewRedisProvider(),
	}, nil
}

func (subscriber *RedisAddPlayerRequestedSubscriber) Subscribe(handler func(*commonappevent.AddPlayerRequestedAppEvent)) func() {
	unsubscriber := subscriber.redisProvider.Subscribe(
		commonappevent.NewAddPlayerRequestedAppEventChannel(),
		func(message []byte) {
			event := commonappevent.DeserializeAddPlayerRequestedAppEvent(message)
			handler(&event)
		})

	return func() {
		unsubscriber()
	}
}
