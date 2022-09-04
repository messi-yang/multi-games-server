package gameunitmaptickedeventbus

import (
	"fmt"

	"github.com/asaskevich/EventBus"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/event/gameunitmaptickedevent"
	"github.com/google/uuid"
)

type eventBus struct {
	eventBus   EventBus.Bus
	eventTopic string
}

type eventBusCallback = func()

var eventBusInstance *eventBus

func GetEventBus() gameunitmaptickedevent.Event {
	if eventBusInstance == nil {
		eventBusInstance = &eventBus{
			eventBus:   EventBus.New(),
			eventTopic: "GAME_COMPUTED",
		}
	}
	return eventBusInstance
}

func (gue *eventBus) Publish(gameId uuid.UUID) {
	topic := fmt.Sprintf("%s-%s", gue.eventTopic, gameId)
	gue.eventBus.Publish(topic)
}

func (gue *eventBus) Subscribe(gameId uuid.UUID, callback eventBusCallback) (unsubscriber func()) {
	topic := fmt.Sprintf("%s-%s", gue.eventTopic, gameId)
	gue.eventBus.Subscribe(topic, callback)

	return func() {
		gue.eventBus.Unsubscribe(topic, callback)
	}
}
