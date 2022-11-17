package rediseventbus

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/application/eventbus"
	"github.com/dum-dum-genius/game-of-liberty-computer/common/infrastructure/service"
)

type redisIntegrationEventBus[T any] struct {
	redisInfrastructureService *service.RedisInfrastructureService
}

type redisIntegrationEventBusConfiguration[T any] func(eventBus *redisIntegrationEventBus[T])

func NewRedisIntegrationEventBus[T any](cfgs ...redisIntegrationEventBusConfiguration[T]) eventbus.IntegrationEventBus[T] {
	eventBus := &redisIntegrationEventBus[T]{}
	for _, cfg := range cfgs {
		cfg(eventBus)
	}
	return eventBus
}

func WithRedisInfrastructureService[T any]() redisIntegrationEventBusConfiguration[T] {
	return func(eventBus *redisIntegrationEventBus[T]) {
		eventBus.redisInfrastructureService = service.NewRedisInfrastructureService()
	}
}

func (gue *redisIntegrationEventBus[T]) Publish(topic string, event T) {
	jsonBytes, _ := json.Marshal(event)
	gue.redisInfrastructureService.Publish(topic, jsonBytes)
}

func (gue *redisIntegrationEventBus[T]) Subscribe(topic string, callback func(event T)) (unsubscriber func()) {
	redisUnsubscriber := gue.redisInfrastructureService.Subscribe(topic, func(event []byte) {
		var targetEvent T
		json.Unmarshal(event, &targetEvent)
		callback(targetEvent)
	})

	return func() {
		redisUnsubscriber()
	}
}
