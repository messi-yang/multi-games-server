package redispubsub

import (
	"context"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/messaging/redisclient"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/util/jsonutil"
	"github.com/go-redis/redis/v9"
)

type ChannelSubscriber[T any] interface {
	Subscribe(channel string, handler func(T)) (channelUnsubscriber func())
}

type channelSubscriber[T any] struct {
	redisClient *redis.Client
}

func NewChannelSubscriber[T any]() ChannelSubscriber[T] {
	return &channelSubscriber[T]{redisClient: redisclient.NewRedisClient()}
}

func (channelSubscriber *channelSubscriber[T]) Subscribe(channel string, handler func(T)) (channelUnsubscriber func()) {
	pubsub := channelSubscriber.redisClient.Subscribe(context.TODO(), channel)
	go func() {
		for msg := range pubsub.Channel() {
			message, err := jsonutil.Unmarshal[T]([]byte(msg.Payload))
			if err != nil {
				return
			}
			handler(message)
		}
	}()

	channelUnsubscriber = func() {
		pubsub.Close()
	}
	return channelUnsubscriber
}
