package redisinteventpublisher

import (
	"context"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/intevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/common/util/jsonutil"
	"github.com/go-redis/redis/v9"
)

type publisher struct {
	redisClient *redis.Client
}

func New(redisClient *redis.Client) intevent.Publisher {
	return &publisher{redisClient: redisClient}
}

func (publisher *publisher) Publish(channel string, event intevent.Event) error {
	message := jsonutil.Marshal(event)
	return publisher.redisClient.Publish(context.TODO(), channel, message).Err()
}
