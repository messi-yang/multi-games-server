package redis

import (
	commonapplicationevent "github.com/dum-dum-genius/game-of-liberty-computer/server/common/application/event"
	commonnotification "github.com/dum-dum-genius/game-of-liberty-computer/server/common/application/notification"
	commonredis "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification/redis"
)

type RedisRemovePlayerRequestedSubscriber struct {
	redisProvider *commonredis.RedisProvider
}

func NewRedisRemovePlayerRequestedSubscriber() (commonnotification.NotificationSubscriber[*commonapplicationevent.RemovePlayerRequestedApplicationEvent], error) {
	return &RedisRemovePlayerRequestedSubscriber{
		redisProvider: commonredis.NewRedisProvider(),
	}, nil
}

func (subscriber *RedisRemovePlayerRequestedSubscriber) Subscribe(handler func(*commonapplicationevent.RemovePlayerRequestedApplicationEvent)) func() {
	unsubscriber := subscriber.redisProvider.Subscribe(
		commonapplicationevent.NewRemovePlayerRequestedApplicationEventChannel(),
		func(message []byte) {
			event := commonapplicationevent.DeserializeRemovePlayerRequestedApplicationEvent(message)
			handler(&event)
		})

	return func() {
		unsubscriber()
	}
}
