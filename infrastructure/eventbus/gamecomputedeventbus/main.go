package gamecomputedeventbus

import (
	"fmt"

	"github.com/DumDumGeniuss/game-of-liberty-computer/application/event/gamecomputedevent"
	"github.com/asaskevich/EventBus"
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

func (gue *gameComputeEventBus) Subscribe(gameId uuid.UUID, callback gameComputeEventCallback) {
	topic := fmt.Sprintf("%s-%s", gue.eventTopic, gameId)
	gue.eventBus.Subscribe(topic, callback)
}

func (gue *gameComputeEventBus) Unsubscribe(gameId uuid.UUID, callback gameComputeEventCallback) {
	topic := fmt.Sprintf("%s-%s", gue.eventTopic, gameId)
	gue.eventBus.Unsubscribe(topic, callback)
}
