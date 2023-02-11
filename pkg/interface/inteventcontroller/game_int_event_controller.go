package inteventcontroller

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/intevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/newappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/redissub"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/library/jsonmarshaller"
)

func NewGameIntEventController(newGameAppService newappservice.GameAppService) {
	gameAdminChannelUnsubscriber := redissub.New().Subscribe(
		intevent.CreateGameAdminChannel(),
		func(message []byte) {
			intEvent, err := jsonmarshaller.Unmarshal[intevent.GenericIntEvent](message)
			if err != nil {
				return
			}

			switch intEvent.Name {
			case intevent.MoveRequestedIntEventName:
				event, err := jsonmarshaller.Unmarshal[intevent.MoveRequestedIntEvent](message)
				if err != nil {
					return
				}
				newGameAppService.MovePlayer(event.GameId, event.PlayerId, event.Direction)
			case intevent.DestroyItemRequestedIntEventName:
				event, err := jsonmarshaller.Unmarshal[intevent.DestroyItemRequestedIntEvent](message)
				if err != nil {
					return
				}
				newGameAppService.DestroyItem(event.GameId, event.PlayerId, event.Location)
			case intevent.PlaceItemRequestedIntEventName:
				event, err := jsonmarshaller.Unmarshal[intevent.PlaceItemRequestedIntEvent](message)
				if err != nil {
					return
				}
				newGameAppService.PlaceItem(event.GameId, event.PlayerId, event.Location, event.ItemId)
			case intevent.LeaveGameRequestedIntEventName:
				event, err := jsonmarshaller.Unmarshal[intevent.LeaveGameRequestedIntEvent](message)
				if err != nil {
					return
				}
				newGameAppService.LeaveGame(event.GameId, event.PlayerId)
			}
		})
	defer gameAdminChannelUnsubscriber()

	closeConnFlag := make(chan bool)
	for {
		<-closeConnFlag
		return
	}
}
