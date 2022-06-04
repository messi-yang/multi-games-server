package gameupdateevent

import (
	"fmt"

	"github.com/asaskevich/EventBus"
	"github.com/google/uuid"
)

type GameUpdateEvent struct {
	eventBus   EventBus.Bus
	eventTopic string
}

type GameUpdateEvenCallback = func()

var gameUpdateEventImpl *GameUpdateEvent

func GetGameUpdateEventBus() *GameUpdateEvent {
	if gameUpdateEventImpl == nil {
		gameUpdateEventImpl = &GameUpdateEvent{
			eventBus:   EventBus.New(),
			eventTopic: "GAME_UPDATE_EVENT",
		}
	}
	return gameUpdateEventImpl
}

func (gue *GameUpdateEvent) Publish(gameId uuid.UUID) {
	topic := fmt.Sprintf("%s-%s", gue.eventTopic, gameId)
	gue.eventBus.Publish(topic)
}

func (gue *GameUpdateEvent) Subscribe(gameId uuid.UUID, callback GameUpdateEvenCallback) {
	topic := fmt.Sprintf("%s-%s", gue.eventTopic, gameId)
	gue.eventBus.Subscribe(topic, callback)
}

func (gue *GameUpdateEvent) Unsubscribe(gameId uuid.UUID, callback GameUpdateEvenCallback) {
	topic := fmt.Sprintf("%s-%s", gue.eventTopic, gameId)
	gue.eventBus.Subscribe(topic, callback)
}
