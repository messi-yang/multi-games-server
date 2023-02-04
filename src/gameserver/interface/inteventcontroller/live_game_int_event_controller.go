package inteventcontroller

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/interface/redissub"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/gameserver/application/appservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/jsonmarshaller"
)

func NewLiveGameIntEventController(liveGameAppService appservice.LiveGameAppService) {
	liveGameAdminChannelUnsubscriber := redissub.New().Subscribe(
		intevent.CreateLiveGameAdminChannel(),
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
				liveGameAppService.JoinGame(event.LiveGameId, event.PlayerId)
			case intevent.MoveRequestedIntEventName:
				event, err := jsonmarshaller.Unmarshal[intevent.MoveRequestedIntEvent](message)
				if err != nil {
					return
				}
				liveGameAppService.MovePlayer(event.LiveGameId, event.PlayerId, event.Direction)
			case intevent.DestroyItemRequestedIntEventName:
				event, err := jsonmarshaller.Unmarshal[intevent.DestroyItemRequestedIntEvent](message)
				if err != nil {
					return
				}
				liveGameAppService.DestroyItem(event.LiveGameId, event.PlayerId, event.Location)
			case intevent.BuildItemRequestedIntEventName:
				event, err := jsonmarshaller.Unmarshal[intevent.BuildItemRequestedIntEvent](message)
				if err != nil {
					return
				}
				liveGameAppService.BuildItem(event.LiveGameId, event.PlayerId, event.Location, event.ItemId)
			case intevent.LeaveGameRequestedIntEventName:
				event, err := jsonmarshaller.Unmarshal[intevent.LeaveGameRequestedIntEvent](message)
				if err != nil {
					return
				}
				liveGameAppService.LeaveGame(event.LiveGameId, event.PlayerId)
			}
		})
	defer liveGameAdminChannelUnsubscriber()

	closeConnFlag := make(chan bool)
	for {
		<-closeConnFlag
		return
	}
}
