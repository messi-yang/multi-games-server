package redissub

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/rediseprovider"
)

type redisSubscriber struct {
	redisProvider *rediseprovider.Provider
}

func New() intgrevent.IntgrEventSubscriber {
	return &redisSubscriber{
		redisProvider: rediseprovider.New(),
	}
}

func (subscriber *redisSubscriber) Subscribe(channel string, handler func([]byte)) func() {
	unsubscriber := subscriber.redisProvider.Subscribe(channel, handler)

	return unsubscriber
}
