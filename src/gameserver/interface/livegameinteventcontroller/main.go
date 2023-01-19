package livegameinteventcontroller

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/interface/redissub"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/gameserver/application/service/livegameappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/jsonmarshaller"
)

func New(liveGameAppService livegameappservice.Service) {
	liveGameAdminChannelUnsubscriber := redissub.New().Subscribe(
		intevent.CreateLiveGameAdminChannel(),
		func(message []byte) {
			intEvent, err := jsonmarshaller.Unmarshal[intevent.GenericintEvent](message)
			if err != nil {
				return
			}

			switch intEvent.Name {
			case intevent.ChangeCameraRequestedintEventName:
				event, err := jsonmarshaller.Unmarshal[intevent.ChangeCameraRequestedintEvent](message)
				if err != nil {
					return
				}
				liveGameAppService.ChangeCamera(event.LiveGameId, event.PlayerId, event.Camera)
			case intevent.JoinGameRequestedintEventName:
				event, err := jsonmarshaller.Unmarshal[intevent.JoinGameRequestedintEvent](message)
				if err != nil {
					return
				}
				liveGameAppService.AddPlayer(event.LiveGameId, event.PlayerId)
			case intevent.DestroyItemRequestedintEventName:
				event, err := jsonmarshaller.Unmarshal[intevent.DestroyItemRequestedintEvent](message)
				if err != nil {
					return
				}
				liveGameAppService.DestroyItem(event.LiveGameId, event.Location)
			case intevent.BuildItemRequestedintEventName:
				event, err := jsonmarshaller.Unmarshal[intevent.BuildItemRequestedintEvent](message)
				if err != nil {
					return
				}
				liveGameAppService.BuildItem(event.LiveGameId, event.Location, event.ItemId)
			case intevent.LeaveGameRequestedintEventName:
				event, err := jsonmarshaller.Unmarshal[intevent.LeaveGameRequestedintEvent](message)
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
