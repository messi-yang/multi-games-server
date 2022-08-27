package gameunitsupdatedeventbus

import (
	"fmt"
	"time"

	"github.com/asaskevich/EventBus"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/coordinatedto"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/event/gameunitsupdatedevent"
	"github.com/google/uuid"
)

type gameUnitsUpdatedEventBus struct {
	eventBus   EventBus.Bus
	eventTopic string
}

type gameUnitsUpdatedEventCallback = func(coordinateDTOs []coordinatedto.CoordinateDTO, updatedAt time.Time)

var gameUnitsUpdatedEventInstance *gameUnitsUpdatedEventBus

func GetUnitsUpdatedEventBus() gameunitsupdatedevent.UnitsUpdatedEvent {
	if gameUnitsUpdatedEventInstance == nil {
		gameUnitsUpdatedEventInstance = &gameUnitsUpdatedEventBus{
			eventBus:   EventBus.New(),
			eventTopic: "GAME_UNITS_UPDATED",
		}
	}
	return gameUnitsUpdatedEventInstance
}

func (gue *gameUnitsUpdatedEventBus) Publish(gameId uuid.UUID, coordinateDTOs []coordinatedto.CoordinateDTO, updatedAt time.Time) {
	topic := fmt.Sprintf("%s-%s", gue.eventTopic, gameId)
	gue.eventBus.Publish(topic, coordinateDTOs, updatedAt)
}

func (gue *gameUnitsUpdatedEventBus) Subscribe(gameId uuid.UUID, callback gameUnitsUpdatedEventCallback) (unsubscriber func()) {
	topic := fmt.Sprintf("%s-%s", gue.eventTopic, gameId)
	gue.eventBus.Subscribe(topic, callback)

	return func() {
		gue.eventBus.Unsubscribe(topic, callback)
	}
}
