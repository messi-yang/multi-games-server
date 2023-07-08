package redisservermessagemediator

import (
	"context"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/redisclient"
	"github.com/go-redis/redis/v9"
)

// Server Message Mediator, used for the messages going between servers and okay to be lost
type Mediator interface {
	Send(channel string, message []byte)
	Receive(channel string, messageHandler func(message []byte)) (unsubscriber func())
}

type mediate struct {
	redisClient *redis.Client
}

func NewMediator() Mediator {
	return &mediate{
		redisClient: redisclient.GetRedisClient(),
	}
}

func (mediate *mediate) Send(channel string, message []byte) {
	mediate.redisClient.Publish(context.TODO(), channel, message)
}

func (mediate *mediate) Receive(channel string, messageHandler func(message []byte)) (messageUnsubscriber func()) {
	pubsub := mediate.redisClient.Subscribe(context.TODO(), channel)
	go func() {
		for msg := range pubsub.Channel() {
			messageHandler([]byte(msg.Payload))
		}
	}()

	return func() {
		pubsub.Close()
	}
}
