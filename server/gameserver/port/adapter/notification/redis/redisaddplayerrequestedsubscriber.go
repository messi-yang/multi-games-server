package redis

import (
	"encoding/json"

	commonapplicationevent "github.com/dum-dum-genius/game-of-liberty-computer/server/common/application/event"
	commonnotification "github.com/dum-dum-genius/game-of-liberty-computer/server/common/application/notification"
	commonredis "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification/redis"
)

type RedisAddPlayerRequestedSubscriber struct {
	redisProvider *commonredis.RedisProvider
}

func NewRedisAddPlayerRequestedSubscriber() (commonnotification.NotificationSubscriber[*commonapplicationevent.AddPlayerRequestedApplicationEvent], error) {
	return &RedisAddPlayerRequestedSubscriber{
		redisProvider: commonredis.NewRedisProvider(),
	}, nil
}

func (subscriber *RedisAddPlayerRequestedSubscriber) Subscribe(handler func(*commonapplicationevent.AddPlayerRequestedApplicationEvent)) func() {
	unsubscriber := subscriber.redisProvider.Subscribe(commonapplicationevent.NewAddPlayerRequestedApplicationEventChannel(), func(message []byte) {
		var event commonapplicationevent.AddPlayerRequestedApplicationEvent
		json.Unmarshal(message, &event)
		handler(&event)
	})

	return func() {
		unsubscriber()
	}
}
