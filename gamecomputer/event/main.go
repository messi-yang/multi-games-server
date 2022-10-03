package event

import (
	"encoding/json"
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/application/service/gameroomservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/infrastructure/memory/gameroommemory"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/config"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/memoryeventbus"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/dto/coordinatedto"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/dto/playerdto"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/event/addplayerrequestedevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/event/removeplayerrequestedevent"
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
		reviveunitsrequestedevent.NewEventTopic(gameId),
		func(event []byte) {
			var reviveUnitsRequestedEvent reviveunitsrequestedevent.Event
			json.Unmarshal(event, &reviveUnitsRequestedEvent)

			coordinates, err := coordinatedto.FromDtoList(reviveUnitsRequestedEvent.Payload.Coordinates)
			if err != nil {
				return
			}

			gameRoomService.ReviveUnits(gameId, coordinates)
		},
	)
	defer reviveUnitsRequestedEventUnsubscriber()

	addPlayerRequestedEventUnsubscriber := eventBus.Subscribe(
		addplayerrequestedevent.NewEventTopic(gameId),
		func(event []byte) {
			var addPlayerRequestedEvent addplayerrequestedevent.Event
			json.Unmarshal(event, &addPlayerRequestedEvent)

			player := playerdto.FromDto(addPlayerRequestedEvent.Payload.Player)
			gameRoomService.AddPlayer(gameId, player)

			fmt.Println(gameRoomService.GetPlayers(gameId))
		},
	)
	defer addPlayerRequestedEventUnsubscriber()

	removePlayerRequestedEventUnsubscriber := eventBus.Subscribe(
		removeplayerrequestedevent.NewEventTopic(gameId),
		func(event []byte) {
			var removePlayerRequestedEvent removeplayerrequestedevent.Event
			json.Unmarshal(event, &removePlayerRequestedEvent)

			gameRoomService.RemovePlayer(gameId, removePlayerRequestedEvent.Payload.PlayerId)

			fmt.Println(gameRoomService.GetPlayers(gameId))
		},
	)
	defer removePlayerRequestedEventUnsubscriber()

	closeConnFlag := make(chan bool)
	for {
		<-closeConnFlag
		return
	}
}
