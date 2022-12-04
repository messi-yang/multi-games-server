package redis

import (
	"encoding/json"

	commonapplicationevent "github.com/dum-dum-genius/game-of-liberty-computer/server/common/application/event"
	commonnotification "github.com/dum-dum-genius/game-of-liberty-computer/server/common/notification"
	commonredis "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification/redis"
)

type RedisZoomAreaRequestedSubscriber struct {
	redisProvider *commonredis.RedisProvider
}

func NewRedisZoomAreaRequestedSubscriber() (commonnotification.NotificationSubscriber[commonapplicationevent.ZoomAreaRequestedApplicationEvent], error) {
	return &RedisZoomAreaRequestedSubscriber{
		redisProvider: commonredis.NewRedisProvider(),
	}, nil
}

func (subscriber *RedisZoomAreaRequestedSubscriber) Subscribe(handler func(commonapplicationevent.ZoomAreaRequestedApplicationEvent)) func() {
	unsubscriber := subscriber.redisProvider.Subscribe(commonapplicationevent.NewZoomAreaRequestedApplicationEventChannel(), func(message []byte) {
		var event commonapplicationevent.ZoomAreaRequestedApplicationEvent
		json.Unmarshal(message, &event)
		handler(event)
	})

	return func() {
		unsubscriber()
	}
}
