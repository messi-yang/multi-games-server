package redis

import (
	commonappevent "github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/event"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/integrationeventsubscriber"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/rediseprovider"
)

type RedisAddPlayerRequestedSubscriber struct {
	redisProvider *rediseprovider.Provider
}

func NewRedisAddPlayerRequestedSubscriber() (integrationeventsubscriber.DeprecatedSubscriber[*commonappevent.AddPlayerRequestedAppEvent], error) {
	return &RedisAddPlayerRequestedSubscriber{
		redisProvider: rediseprovider.New(),
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
