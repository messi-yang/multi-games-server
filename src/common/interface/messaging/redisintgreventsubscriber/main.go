package redisintgreventsubscriber

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/rediseprovider"
)

type Subscriber struct {
	redisProvider *rediseprovider.Provider
}

func New() intgrevent.IntgrEventSubscriber {
	return &Subscriber{
		redisProvider: rediseprovider.New(),
	}
}

func (subscriber *Subscriber) Subscribe(channel string, handler func([]byte)) func() {
	unsubscriber := subscriber.redisProvider.Subscribe(channel, handler)

	return unsubscriber
}
