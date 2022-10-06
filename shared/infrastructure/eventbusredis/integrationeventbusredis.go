package eventbusredis

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/application/eventbus"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/infrastructureservice"
)

type integrationEventBusRedis struct {
	redisInfrastructureService infrastructureservice.RedisInfrastructureService
}

type integrationEventBusRedisCallback = func(event []byte)

var integrationEventBusRedisInstance *integrationEventBusRedis

func GetIntegrationEventBusRedis() eventbus.IntegrationEventBus {
	if integrationEventBusRedisInstance == nil {
		integrationEventBusRedisInstance = &integrationEventBusRedis{
			redisInfrastructureService: infrastructureservice.NewRedisInfrastructureService(),
		}
	}
	return integrationEventBusRedisInstance
}

func (gue *integrationEventBusRedis) Publish(topic string, event []byte) {
	gue.redisInfrastructureService.Publish(topic, event)
}

func (gue *integrationEventBusRedis) Subscribe(topic string, callback integrationEventBusRedisCallback) (unsubscriber func()) {
	redisUnsubscriber := gue.redisInfrastructureService.Subscribe(topic, callback)

	return func() {
		redisUnsubscriber()
	}
}
