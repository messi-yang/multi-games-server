package eventbusredis

import (
	"github.com/asaskevich/EventBus"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/application/eventbus"
)

type integrationEventBusRedis struct {
	eventBus EventBus.Bus
}

type integrationEventBusRedisCallback = func(event []byte)

var integrationEventBusRedisInstance *integrationEventBusRedis

func GetIntegrationEventBusRedis() eventbus.IntegrationEventBus {
	if integrationEventBusRedisInstance == nil {
		integrationEventBusRedisInstance = &integrationEventBusRedis{
			eventBus: EventBus.New(),
		}
	}
	return integrationEventBusRedisInstance
}

func (gue *integrationEventBusRedis) Publish(topic string, event []byte) {
	// I don't what the heck is going on, without this "go" we can't let the method finish
	go gue.eventBus.Publish(topic, event)
}

func (gue *integrationEventBusRedis) Subscribe(topic string, callback integrationEventBusRedisCallback) (unsubscriber func()) {
	gue.eventBus.Subscribe(topic, callback)

	return func() {
		gue.eventBus.Unsubscribe(topic, callback)
	}
}
