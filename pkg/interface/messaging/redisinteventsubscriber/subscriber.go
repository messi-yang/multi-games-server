package redisinteventsubscriber

import (
	"context"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/intevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/common/client/redisclient"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/common/util/jsonutil"
	"github.com/go-redis/redis/v9"
)

type subscriber[T intevent.Event] struct {
	redisClient *redis.Client
}

func New[T intevent.Event]() intevent.Subscriber[T] {
	redisClient := redisclient.New()
	return &subscriber[T]{redisClient: redisClient}
}

func (subscriber *subscriber[T]) Subscribe(channel string, handler func(T)) (unsubscriber func()) {
	pubsub := subscriber.redisClient.Subscribe(context.TODO(), channel)
	go func() {
		for msg := range pubsub.Channel() {
			intEvent, err := jsonutil.Unmarshal[T]([]byte(msg.Payload))
			if err != nil {
				return
			}
			handler(intEvent)
		}
	}()

	unsubscriber = func() {
		pubsub.Close()
	}
	return unsubscriber
}
