package redisintegrationeventsubscriber

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/integrationevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/rediseprovider"
)

type subscriber struct {
	redisProvider *rediseprovider.Provider
}

func New() (integrationevent.Subscriber, error) {
	return &subscriber{
		redisProvider: rediseprovider.New(),
	}, nil
}

func (subscriber *subscriber) Subscribe(channel string, handler func([]byte)) func() {
	unSubscriber := subscriber.redisProvider.Subscribe(channel, handler)

	return unSubscriber
}
