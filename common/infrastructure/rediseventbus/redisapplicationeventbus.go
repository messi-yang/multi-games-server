package rediseventbus

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/application/eventbus"
	"github.com/dum-dum-genius/game-of-liberty-computer/common/infrastructure/service"
)

type redisApplicationEventBus[T any] struct {
	redisInfrastructureService *service.RedisInfrastructureService
}

type redisApplicationEventBusConfiguration[T any] func(eventBus *redisApplicationEventBus[T])

func NewRedisApplicationEventBus[T any](cfgs ...redisApplicationEventBusConfiguration[T]) eventbus.ApplicationEventBus[T] {
	eventBus := &redisApplicationEventBus[T]{}
	for _, cfg := range cfgs {
		cfg(eventBus)
	}
	return eventBus
}

func WithRedisInfrastructureService[T any]() redisApplicationEventBusConfiguration[T] {
	return func(eventBus *redisApplicationEventBus[T]) {
		eventBus.redisInfrastructureService = service.NewRedisInfrastructureService()
	}
}

func (gue *redisApplicationEventBus[T]) Publish(topic string, event T) {
	jsonBytes, _ := json.Marshal(event)
	gue.redisInfrastructureService.Publish(topic, jsonBytes)
}

func (gue *redisApplicationEventBus[T]) Subscribe(topic string, callback func(event T)) (unsubscriber func()) {
	redisUnsubscriber := gue.redisInfrastructureService.Subscribe(topic, func(event []byte) {
		var targetEvent T
		json.Unmarshal(event, &targetEvent)
		callback(targetEvent)
	})

	return func() {
		redisUnsubscriber()
	}
}
