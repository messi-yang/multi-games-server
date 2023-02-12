package redisinteventsubscriber

import (
	"context"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/intevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/json"
)

type subscriber[T intevent.IntEvent] struct{}

func New[T intevent.IntEvent]() intevent.IntEventSubscriber[T] {
	return &subscriber[T]{}
}

func (subscriber *subscriber[T]) Subscribe(channel string, handler func(T)) func() {
	pubsub := redisClient.Subscribe(context.TODO(), channel)
	go func() {
		for msg := range pubsub.Channel() {
			intEvent, err := json.Unmarshal[T]([]byte(msg.Payload))
			if err != nil {
				continue
			}
			handler(intEvent)
		}
	}()

	return func() {
		pubsub.Close()
	}
}
