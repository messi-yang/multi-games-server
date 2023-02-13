package redisinteventpublisher

import (
	"context"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/intevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/json"
)

type publisher struct {
}

func New() intevent.IntEventPublisher {
	return &publisher{}
}

func (publisher *publisher) Publish(channel string, event intevent.IntEvent) error {
	message := json.Marshal(event)
	err := redisClient.Publish(context.TODO(), channel, message).Err()
	if err != nil {
		return err
	}
	return nil
}
