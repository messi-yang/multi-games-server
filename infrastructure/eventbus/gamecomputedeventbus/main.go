package gamecomputedeventbus

import (
	"fmt"

	"github.com/asaskevich/EventBus"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/event/gamecomputedevent"
	"github.com/google/uuid"
)

type gameComputeEventBus struct {
	eventBus   EventBus.Bus
	eventTopic string
}

type gameComputeEventCallback = func()

var gameComputeEventInstance *gameComputeEventBus

func GetGameComputedEventBus() gamecomputedevent.GameComputedEvent {
	if gameComputeEventInstance == nil {
		gameComputeEventInstance = &gameComputeEventBus{
			eventBus:   EventBus.New(),
			eventTopic: "GAME_COMPUTED",
		}
	}
	return gameComputeEventInstance
}

func (gue *gameComputeEventBus) Publish(gameId uuid.UUID) {
	topic := fmt.Sprintf("%s-%s", gue.eventTopic, gameId)
	gue.eventBus.Publish(topic)
}

func (gue *gameComputeEventBus) Subscribe(gameId uuid.UUID, callback gameComputeEventCallback) (unsubscriber func()) {
	topic := fmt.Sprintf("%s-%s", gue.eventTopic, gameId)
	gue.eventBus.Subscribe(topic, callback)

	return func() {
		gue.eventBus.Unsubscribe(topic, callback)
	}
}
