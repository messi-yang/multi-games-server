package redisintgreventpublisher

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/messaging/intgreventpublisher"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/rediseprovider"
)

type publisher struct {
	redisProvider *rediseprovider.Provider
}

func New() intgreventpublisher.Publisher {
	return &publisher{
		redisProvider: rediseprovider.New(),
	}
}

func (publisher *publisher) Publish(channel string, message []byte) error {
	err := publisher.redisProvider.Publish(channel, message)
	if err != nil {
		return err
	}
	return nil
}
