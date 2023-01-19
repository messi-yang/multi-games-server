package redispub

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/rediseprovider"
)

type redisPublisher struct {
	redisProvider *rediseprovider.Provider
}

func NewRedisPublisher() intgrevent.IntgrEventPublisher {
	return &redisPublisher{
		redisProvider: rediseprovider.New(),
	}
}

func (redisPublisher *redisPublisher) Publish(channel string, message []byte) error {
	err := redisPublisher.redisProvider.Publish(channel, message)
	if err != nil {
		return err
	}
	return nil
}
