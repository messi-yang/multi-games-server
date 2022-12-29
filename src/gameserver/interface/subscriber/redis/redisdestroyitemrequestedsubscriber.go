package redis

import (
	commonappevent "github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/event"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/integrationeventsubscriber"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/rediseprovider"
)

type RedisDestroyItemRequestedSubscriber struct {
	redisProvider *rediseprovider.Provider
}

func NewRedisDestroyItemRequestedSubscriber() (integrationeventsubscriber.DeprecatedSubscriber[*commonappevent.DestroyItemRequestedAppEvent], error) {
	return &RedisDestroyItemRequestedSubscriber{
		redisProvider: rediseprovider.New(),
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
