package eventbusredis

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/application/eventbus"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/infrastructureservice"
)

type redisIntegrationEventBus struct {
	redisService infrastructureservice.RedisService
}

type redisIntegrationEventBusCallback = func(event []byte)

type RedisIntegrationEventBusCallbackConfiguration struct {
	RedisService infrastructureservice.RedisService
}

var redisIntegrationEventBusInstance *redisIntegrationEventBus

func NewRedisIntegrationEventBus(config RedisIntegrationEventBusCallbackConfiguration) eventbus.IntegrationEventBus {
	if redisIntegrationEventBusInstance == nil {
		redisIntegrationEventBusInstance = &redisIntegrationEventBus{
			redisService: infrastructureservice.NewRedisService(),
		}
	}
	return redisIntegrationEventBusInstance
}

func (gue *redisIntegrationEventBus) Publish(topic string, event []byte) {
	gue.redisService.Publish(topic, event)
}

func (gue *redisIntegrationEventBus) Subscribe(topic string, callback redisIntegrationEventBusCallback) (unsubscriber func()) {
	redisUnsubscriber := gue.redisService.Subscribe(topic, callback)

	return func() {
		redisUnsubscriber()
	}
}
