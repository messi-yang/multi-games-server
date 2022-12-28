package redis

import (
	commonredis "github.com/dum-dum-genius/game-of-liberty-computer/src/common/adapter/notification/redis"
	commonappevent "github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/event"
	commonnotification "github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/notification"
)

type RedisRemovePlayerRequestedSubscriber struct {
	redisProvider *commonredis.RedisProvider
}

func NewRedisRemovePlayerRequestedSubscriber() (commonnotification.NotificationSubscriber[*commonappevent.RemovePlayerRequestedAppEvent], error) {
	return &RedisRemovePlayerRequestedSubscriber{
		redisProvider: commonredis.NewRedisProvider(),
	}, nil
}

func (subscriber *RedisRemovePlayerRequestedSubscriber) Subscribe(handler func(*commonappevent.RemovePlayerRequestedAppEvent)) func() {
	unsubscriber := subscriber.redisProvider.Subscribe(
		commonappevent.NewRemovePlayerRequestedAppEventChannel(),
		func(message []byte) {
			event := commonappevent.DeserializeRemovePlayerRequestedAppEvent(message)
			handler(&event)
		})

	return func() {
		unsubscriber()
	}
}
