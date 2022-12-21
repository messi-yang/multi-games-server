package redis

import (
	commonappevent "github.com/dum-dum-genius/game-of-liberty-computer/server/common/application/event"
	commonnotification "github.com/dum-dum-genius/game-of-liberty-computer/server/common/application/notification"
	commonredis "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification/redis"
)

type RedisBuildItemRequestedSubscriber struct {
	redisProvider *commonredis.RedisProvider
}

func NewRedisBuildItemRequestedSubscriber() (commonnotification.NotificationSubscriber[*commonappevent.BuildItemRequestedAppEvent], error) {
	return &RedisBuildItemRequestedSubscriber{
		redisProvider: commonredis.NewRedisProvider(),
	}, nil
}

func (subscriber *RedisBuildItemRequestedSubscriber) Subscribe(handler func(*commonappevent.BuildItemRequestedAppEvent)) func() {
	unsubscriber := subscriber.redisProvider.Subscribe(
		commonappevent.NewBuildItemRequestedAppEventChannel(),
		func(message []byte) {
			event := commonappevent.DeserializeBuildItemRequestedAppEvent(message)
			handler(&event)
		})

	return func() {
		unsubscriber()
	}
}
