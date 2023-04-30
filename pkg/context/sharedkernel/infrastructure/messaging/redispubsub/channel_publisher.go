package redispubsub

import (
	"context"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/messaging/redisclient"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/util/jsonutil"
	"github.com/go-redis/redis/v9"
)

type ChannelPublisher interface {
	Publish(channel string, event any) error
}

type channelPublisher struct {
	redisClient *redis.Client
}

func NewChannelPublisher() ChannelPublisher {
	return &channelPublisher{redisClient: redisclient.GetRedisClient()}
}

func (channelPublisher *channelPublisher) Publish(channel string, event any) error {
	message := jsonutil.Marshal(event)
	return channelPublisher.redisClient.Publish(context.TODO(), channel, message).Err()
}
