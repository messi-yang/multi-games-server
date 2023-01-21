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
			case intevent.ChangeCameraRequestedIntEventName:
				event, err := jsonmarshaller.Unmarshal[intevent.ChangeCameraRequestedIntEvent](message)
				if err != nil {
					return
				}
				liveGameAppService.ChangeCamera(event.LiveGameId, event.PlayerId, event.Camera)
			case intevent.JoinGameRequestedIntEventName:
				event, err := jsonmarshaller.Unmarshal[intevent.JoinGameRequestedIntEvent](message)
				if err != nil {
					return
				}
				liveGameAppService.AddPlayer(event.LiveGameId, event.PlayerId)
			case intevent.DestroyItemRequestedIntEventName:
				event, err := jsonmarshaller.Unmarshal[intevent.DestroyItemRequestedIntEvent](message)
				if err != nil {
					return
				}
				liveGameAppService.DestroyItem(event.LiveGameId, event.Location)
			case intevent.BuildItemRequestedIntEventName:
				event, err := jsonmarshaller.Unmarshal[intevent.BuildItemRequestedIntEvent](message)
				if err != nil {
					return
				}
				liveGameAppService.BuildItem(event.LiveGameId, event.Location, event.ItemId)
			case intevent.LeaveGameRequestedIntEventName:
				event, err := jsonmarshaller.Unmarshal[intevent.LeaveGameRequestedIntEvent](message)
				if err != nil {
					return
				}
				liveGameAppService.RemovePlayer(event.LiveGameId, event.PlayerId)
			}
		})
	defer liveGameAdminChannelUnsubscriber()

	closeConnFlag := make(chan bool)
	for {
		<-closeConnFlag
		return
	}
}
