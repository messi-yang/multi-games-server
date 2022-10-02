package eventbus

import (
	"github.com/asaskevich/EventBus"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/application/eventbus"
)

type eventBus struct {
	eventBus EventBus.Bus
}

type eventBusCallback = func(event []byte)

var eventBusInstance *eventBus

func GetEventBus() eventbus.EventBus {
	if eventBusInstance == nil {
		eventBusInstance = &eventBus{
			eventBus: EventBus.New(),
		}
	}
	return eventBusInstance
}

func (gue *eventBus) Publish(topic string, event []byte) {
	// I don't what the heck is going on, without this "go" we can't let the method finish
	go gue.eventBus.Publish(topic, event)
}

func (gue *eventBus) Subscribe(topic string, callback eventBusCallback) (unsubscriber func()) {
	gue.eventBus.Subscribe(topic, callback)

	return func() {
		gue.eventBus.Unsubscribe(topic, callback)
	}
}
