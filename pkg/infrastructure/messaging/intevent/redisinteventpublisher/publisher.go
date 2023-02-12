package redisinteventpublisher

import (
	"context"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/intevent"
)

type publisher struct {
}

func New() intevent.IntEventPublisher {
	return &publisher{}
}

func (publisher *publisher) Publish(channel string, event intevent.IntEvent) error {
	message := event.Marshal()
	err := redisClient.Publish(context.TODO(), channel, message).Err()
	if err != nil {
		return err
	}
	return nil
}
