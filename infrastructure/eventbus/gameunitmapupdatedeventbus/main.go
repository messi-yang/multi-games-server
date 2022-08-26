package gameunitmapupdatedeventbus

import (
	"fmt"

	"github.com/asaskevich/EventBus"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/event/gameunitmapupdatedevent"
	"github.com/google/uuid"
)

type gameUnitMapUpdatedEventBus struct {
	eventBus   EventBus.Bus
	eventTopic string
}

type gameUnitMapUpdatedEventCallback = func()

var gameUnitMapUpdatedEventInstance *gameUnitMapUpdatedEventBus

func GetGameUnitMapUpdatedEventBus() gameunitmapupdatedevent.GameUnitMapUpdatedEvent {
	if gameUnitMapUpdatedEventInstance == nil {
		gameUnitMapUpdatedEventInstance = &gameUnitMapUpdatedEventBus{
			eventBus:   EventBus.New(),
			eventTopic: "GAME_COMPUTED",
		}
	}
	return gameUnitMapUpdatedEventInstance
}

func (gue *gameUnitMapUpdatedEventBus) Publish(gameId uuid.UUID) {
	topic := fmt.Sprintf("%s-%s", gue.eventTopic, gameId)
	gue.eventBus.Publish(topic)
}

func (gue *gameUnitMapUpdatedEventBus) Subscribe(gameId uuid.UUID, callback gameUnitMapUpdatedEventCallback) (unsubscriber func()) {
	topic := fmt.Sprintf("%s-%s", gue.eventTopic, gameId)
	gue.eventBus.Subscribe(topic, callback)

	return func() {
		gue.eventBus.Unsubscribe(topic, callback)
	}
}
