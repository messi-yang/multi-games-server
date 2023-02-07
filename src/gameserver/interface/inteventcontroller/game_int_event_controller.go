package inteventcontroller

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/interface/redissub"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/gameserver/application/appservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/jsonmarshaller"
)

func NewGameIntEventController(gameAppService appservice.GameAppService) {
	gameAdminChannelUnsubscriber := redissub.New().Subscribe(
		intevent.CreateGameAdminChannel(),
		func(message []byte) {
			intEvent, err := jsonmarshaller.Unmarshal[intevent.GenericIntEvent](message)
			if err != nil {
				return
			}

			switch intEvent.Name {
			case intevent.JoinGameRequestedIntEventName:
				event, err := jsonmarshaller.Unmarshal[intevent.JoinGameRequestedIntEvent](message)
				if err != nil {
					return
				}
				gameAppService.JoinGame(event.GameId, event.PlayerId)
			case intevent.MoveRequestedIntEventName:
				event, err := jsonmarshaller.Unmarshal[intevent.MoveRequestedIntEvent](message)
				if err != nil {
					return
				}
				gameAppService.MovePlayer(event.GameId, event.PlayerId, event.Direction)
			case intevent.DestroyItemRequestedIntEventName:
				event, err := jsonmarshaller.Unmarshal[intevent.DestroyItemRequestedIntEvent](message)
				if err != nil {
					return
				}
				gameAppService.DestroyItem(event.GameId, event.PlayerId, event.Location)
			case intevent.PlaceItemRequestedIntEventName:
				event, err := jsonmarshaller.Unmarshal[intevent.PlaceItemRequestedIntEvent](message)
				if err != nil {
					return
				}
				gameAppService.PlaceItem(event.GameId, event.PlayerId, event.Location, event.ItemId)
			case intevent.LeaveGameRequestedIntEventName:
				event, err := jsonmarshaller.Unmarshal[intevent.LeaveGameRequestedIntEvent](message)
				if err != nil {
					return
				}
				gameAppService.LeaveGame(event.GameId, event.PlayerId)
			}
		})
	defer gameAdminChannelUnsubscriber()

	closeConnFlag := make(chan bool)
	for {
		<-closeConnFlag
		return
	}
}
