package eventbusredis

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/application/eventbus"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/infrastructureservice"
)

type redisIntegrationEventBus struct {
	redisService infrastructureservice.RedisService
}

type redisIntegrationEventBusCallbackConfiguration func(eventBus *redisIntegrationEventBus) error

func NewRedisIntegrationEventBus(cfgs ...redisIntegrationEventBusCallbackConfiguration) (eventbus.IntegrationEventBus, error) {
	service := &redisIntegrationEventBus{}
	for _, cfg := range cfgs {
		err := cfg(service)
		if err != nil {
			return nil, err
		}
	}
	return service, nil
}

func WithRedisService() redisIntegrationEventBusCallbackConfiguration {
	redisService := infrastructureservice.NewRedisService()
	return func(eventBus *redisIntegrationEventBus) error {
		eventBus.redisService = redisService
		return nil
	}
}

func (gue *redisIntegrationEventBus) Publish(topic string, event []byte) {
	gue.redisService.Publish(topic, event)
}

func (gue *redisIntegrationEventBus) Subscribe(topic string, callback func(event []byte)) (unsubscriber func()) {
	redisUnsubscriber := gue.redisService.Subscribe(topic, callback)

	return func() {
		redisUnsubscriber()
	}
}
