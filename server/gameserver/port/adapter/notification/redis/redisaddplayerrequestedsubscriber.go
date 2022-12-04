package redis

import (
	"encoding/json"

	commonnotification "github.com/dum-dum-genius/game-of-liberty-computer/server/common/notification"
	commonredis "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification/redis"
	commonredisdto "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification/redis/dto"
)

type RedisAddPlayerRequestedSubscriber struct {
	redisProvider *commonredis.RedisProvider
}

func NewRedisAddPlayerRequestedSubscriber() (commonnotification.NotificationSubscriber[commonredisdto.RedisAddPlayerRequestedEvent], error) {
	return &RedisAddPlayerRequestedSubscriber{
		redisProvider: commonredis.NewRedisProvider(),
	}, nil
}

func (subscriber *RedisAddPlayerRequestedSubscriber) Subscribe(handler func(commonredisdto.RedisAddPlayerRequestedEvent)) func() {
	unsubscriber := subscriber.redisProvider.Subscribe(commonredisdto.NewRedisAddPlayerRequestedEventChannel(), func(message []byte) {
		var event commonredisdto.RedisAddPlayerRequestedEvent
		json.Unmarshal(message, &event)
		handler(event)
	})

	return func() {
		unsubscriber()
	}
}
