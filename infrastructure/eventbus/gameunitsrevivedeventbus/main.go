package gameunitsrevivedeventbus

import (
	"fmt"
	"time"

	"github.com/asaskevich/EventBus"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/coordinatedto"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/event/gameunitsrevivedevent"
	"github.com/google/uuid"
)

type gameUnitsRevivedEventBus struct {
	eventBus   EventBus.Bus
	eventTopic string
}

type gameUnitsRevivedEventCallback = func(coordinateDTOs []coordinatedto.CoordinateDTO, updatedAt time.Time)

var gameUnitsRevivedEventInstance *gameUnitsRevivedEventBus

func GetUnitsRevivedEventBus() gameunitsrevivedevent.Event {
	if gameUnitsRevivedEventInstance == nil {
		gameUnitsRevivedEventInstance = &gameUnitsRevivedEventBus{
			eventBus:   EventBus.New(),
			eventTopic: "GAME_UNITS_UPDATED",
		}
	}
	return gameUnitsRevivedEventInstance
}

func (gue *gameUnitsRevivedEventBus) Publish(gameId uuid.UUID, coordinateDTOs []coordinatedto.CoordinateDTO, updatedAt time.Time) {
	topic := fmt.Sprintf("%s-%s", gue.eventTopic, gameId)
	gue.eventBus.Publish(topic, coordinateDTOs, updatedAt)
}

func (gue *gameUnitsRevivedEventBus) Subscribe(gameId uuid.UUID, callback gameUnitsRevivedEventCallback) (unsubscriber func()) {
	topic := fmt.Sprintf("%s-%s", gue.eventTopic, gameId)
	gue.eventBus.Subscribe(topic, callback)

	return func() {
		gue.eventBus.Unsubscribe(topic, callback)
	}
}
