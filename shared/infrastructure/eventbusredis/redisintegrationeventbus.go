package eventbusredis

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/application/eventbus"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/infrastructureservice"
)

type redisIntegrationEventBus struct {
	redisInfrastructureService infrastructureservice.RedisInfrastructureService
}

type redisIntegrationEventBusCallback = func(event []byte)

type RedisIntegrationEventBusCallbackConfiguration struct {
	RedisInfrastructureService infrastructureservice.RedisInfrastructureService
}

var redisIntegrationEventBusInstance *redisIntegrationEventBus

func NewRedisIntegrationEventBus(config RedisIntegrationEventBusCallbackConfiguration) eventbus.IntegrationEventBus {
	if redisIntegrationEventBusInstance == nil {
		redisIntegrationEventBusInstance = &redisIntegrationEventBus{
			redisInfrastructureService: infrastructureservice.NewRedisInfrastructureService(),
		}
	}
	return redisIntegrationEventBusInstance
}

func (gue *redisIntegrationEventBus) Publish(topic string, event []byte) {
	gue.redisInfrastructureService.Publish(topic, event)
}

func (gue *redisIntegrationEventBus) Subscribe(topic string, callback redisIntegrationEventBusCallback) (unsubscriber func()) {
	redisUnsubscriber := gue.redisInfrastructureService.Subscribe(topic, callback)

	return func() {
		redisUnsubscriber()
	}
}
