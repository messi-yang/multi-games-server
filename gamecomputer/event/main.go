package event

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/application/service/gameroomservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/infrastructure/memory/gameroommemory"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/config"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/memoryeventbus"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/event/reviveunitsrequestedevent"
)

func Controller() {
	gameId := config.GetConfig().GetGameId()

	gameRoomRepository := gameroommemory.GetRepository()
	eventBus := memoryeventbus.GetEventBus()
	gameRoomService := gameroomservice.NewService(
		gameroomservice.Configuration{
			GameRoomRepository: gameRoomRepository,
			EventBus:           eventBus,
		},
	)

	reviveUnitsRequestedEventUnsubscriber := eventBus.Subscribe(
		reviveunitsrequestedevent.NewEventTopic(),
		func(event []byte) {
			var reviveUnitsRequestedEvent reviveunitsrequestedevent.Event
			json.Unmarshal(event, &reviveUnitsRequestedEvent)

			coordinates := make([]valueobject.Coordinate, 0)
			for _, coordInPayload := range reviveUnitsRequestedEvent.Payload.Coordinates {
				coordinate, _ := valueobject.NewCoordinate(coordInPayload.X, coordInPayload.Y)
				coordinates = append(coordinates, coordinate)
			}

			gameRoomService.ReviveUnits(gameId, coordinates)
		},
	)
	defer reviveUnitsRequestedEventUnsubscriber()

	closeConnFlag := make(chan bool)
	for {
		<-closeConnFlag
		return
	}
}
