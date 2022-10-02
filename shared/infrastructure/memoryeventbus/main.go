package memoryeventbus

import (
	"github.com/asaskevich/EventBus"
)

type MemoryEventBus struct {
	eventBus EventBus.Bus
}

type eventBusCallback = func(event []byte)

var eventBusInstance *MemoryEventBus

func GetEventBus() *MemoryEventBus {
	if eventBusInstance == nil {
		eventBusInstance = &MemoryEventBus{
			eventBus: EventBus.New(),
		}
	}
	return eventBusInstance
}

func (gue *MemoryEventBus) Publish(topic string, event []byte) {
	// I don't what the heck is going on, without this "go" we can't let the method finish
	go gue.eventBus.Publish(topic, event)
}

func (gue *MemoryEventBus) Subscribe(topic string, callback eventBusCallback) (unsubscriber func()) {
	gue.eventBus.Subscribe(topic, callback)

	return func() {
		gue.eventBus.Unsubscribe(topic, callback)
	}
}
