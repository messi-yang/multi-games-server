package gameunitsrevivedeventbus

import (
	"fmt"

	"github.com/asaskevich/EventBus"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/event/gameunitsrevivedevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/valueobject"
	"github.com/google/uuid"
)

type eventBus struct {
	eventBus   EventBus.Bus
	eventTopic string
}

type eventBusCallback = func(coordinates []valueobject.Coordinate)

var eventBusInstance *eventBus

func GetEventBus() gameunitsrevivedevent.Event {
	if eventBusInstance == nil {
		eventBusInstance = &eventBus{
			eventBus:   EventBus.New(),
			eventTopic: "GAME_UNITS_UPDATED",
		}
	}
	return eventBusInstance
}

func (gue *eventBus) Publish(gameId uuid.UUID, coordinates []valueobject.Coordinate) {
	topic := fmt.Sprintf("%s-%s", gue.eventTopic, gameId)
	gue.eventBus.Publish(topic, coordinates)
}

func (gue *eventBus) Subscribe(gameId uuid.UUID, callback eventBusCallback) (unsubscriber func()) {
	topic := fmt.Sprintf("%s-%s", gue.eventTopic, gameId)
	gue.eventBus.Subscribe(topic, callback)

	return func() {
		gue.eventBus.Unsubscribe(topic, callback)
	}
}
