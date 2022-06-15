package gameunitsupdatedeventbus

import (
	"fmt"

	"github.com/DumDumGeniuss/game-of-liberty-computer/application/event/gameunitsupdatedevent"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/valueobject"
	"github.com/asaskevich/EventBus"
	"github.com/google/uuid"
)

type gameUnitsUpdatedEventBus struct {
	eventBus   EventBus.Bus
	eventTopic string
}

type gameUnitsUpdatedEventCallback = func(coordinates []valueobject.Coordinate)

var gameUnitsUpdatedEventInstance *gameUnitsUpdatedEventBus

func GetGameUnitsUpdatedEventBus() gameunitsupdatedevent.GameUnitsUpdatedEvent {
	if gameUnitsUpdatedEventInstance == nil {
		gameUnitsUpdatedEventInstance = &gameUnitsUpdatedEventBus{
			eventBus:   EventBus.New(),
			eventTopic: "GAME_UNITS_UPDATED",
		}
	}
	return gameUnitsUpdatedEventInstance
}

func (gue *gameUnitsUpdatedEventBus) Publish(gameId uuid.UUID, coordinates []valueobject.Coordinate) {
	topic := fmt.Sprintf("%s-%s", gue.eventTopic, gameId)
	gue.eventBus.Publish(topic, coordinates)
}

func (gue *gameUnitsUpdatedEventBus) Subscribe(gameId uuid.UUID, callback gameUnitsUpdatedEventCallback) {
	topic := fmt.Sprintf("%s-%s", gue.eventTopic, gameId)
	gue.eventBus.Subscribe(topic, callback)
}

func (gue *gameUnitsUpdatedEventBus) Unsubscribe(gameId uuid.UUID, callback gameUnitsUpdatedEventCallback) {
	topic := fmt.Sprintf("%s-%s", gue.eventTopic, gameId)
	gue.eventBus.Unsubscribe(topic, callback)
}
