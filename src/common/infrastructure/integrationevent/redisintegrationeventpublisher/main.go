package redisintegrationeventpublisher

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/event"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/integrationeventpublisher"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/rediseprovider"
)

type publisher struct {
	redisProvider *rediseprovider.Provider
}

func New() integrationeventpublisher.Publisher {
	return &publisher{
		redisProvider: rediseprovider.New(),
	}
}

func (publisher *publisher) Publish(channel string, event event.AppEvent) error {
	message := event.Serialize()

	err := publisher.redisProvider.Publish(channel, message)
	if err != nil {
		return err
	}
	return nil
}
