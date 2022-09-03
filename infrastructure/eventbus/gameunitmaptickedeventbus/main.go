package gameunitmaptickedeventbus

import (
	"fmt"
	"time"

	"github.com/asaskevich/EventBus"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/event/gameunitmaptickedevent"
	"github.com/google/uuid"
)

type gameUnitMapTickedEventBus struct {
	eventBus   EventBus.Bus
	eventTopic string
}

type gameUnitMapTickedEventCallback = func(updatedAt time.Time)

var gameUnitMapTickedEventInstance *gameUnitMapTickedEventBus

func GetGameUnitMapTickedEventBus() gameunitmaptickedevent.Event {
	if gameUnitMapTickedEventInstance == nil {
		gameUnitMapTickedEventInstance = &gameUnitMapTickedEventBus{
			eventBus:   EventBus.New(),
			eventTopic: "GAME_COMPUTED",
		}
	}
	return gameUnitMapTickedEventInstance
}

func (gue *gameUnitMapTickedEventBus) Publish(gameId uuid.UUID, updatedAt time.Time) {
	topic := fmt.Sprintf("%s-%s", gue.eventTopic, gameId)
	gue.eventBus.Publish(topic, updatedAt)
}

func (gue *gameUnitMapTickedEventBus) Subscribe(gameId uuid.UUID, callback gameUnitMapTickedEventCallback) (unsubscriber func()) {
	topic := fmt.Sprintf("%s-%s", gue.eventTopic, gameId)
	gue.eventBus.Subscribe(topic, callback)

	return func() {
		gue.eventBus.Unsubscribe(topic, callback)
	}
}
