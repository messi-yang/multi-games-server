package livegameintgreventcontroller

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/interface/messaging/redisintgreventsubscriber"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/gameserver/application/service/livegameappservice"
)

func New(liveGameAppService livegameappservice.Service) {
	liveGameAdminChannelUnsubscriber := redisintgreventsubscriber.New().Subscribe(
		intgrevent.CreateLiveGameAdminChannel(),
		func(message []byte) {
			intgrEvent, err := intgrevent.Unmarshal[intgrevent.GenericIntgrEvent](message)
			if err != nil {
				return
			}

			switch intgrEvent.Name {
			case intgrevent.ChangeCameraRequestedIntgrEventName:
				event, err := intgrevent.Unmarshal[intgrevent.ChangeCameraRequestedIntgrEvent](message)
				if err != nil {
					return
				}
				liveGameAppService.ChangeCamera(event.LiveGameId, event.PlayerId, event.Camera)
			case intgrevent.JoinGameRequestedIntgrEventName:
				event, err := intgrevent.Unmarshal[intgrevent.JoinGameRequestedIntgrEvent](message)
				if err != nil {
					return
				}
				liveGameAppService.AddPlayer(event.LiveGameId, event.PlayerId)
			case intgrevent.DestroyItemRequestedIntgrEventName:
				event, err := intgrevent.Unmarshal[intgrevent.DestroyItemRequestedIntgrEvent](message)
				if err != nil {
					return
				}
				liveGameAppService.DestroyItem(event.LiveGameId, event.Location)
			case intgrevent.BuildItemRequestedIntgrEventName:
				event, err := intgrevent.Unmarshal[intgrevent.BuildItemRequestedIntgrEvent](message)
				if err != nil {
					return
				}
				liveGameAppService.BuildItem(event.LiveGameId, event.Location, event.ItemId)
			case intgrevent.LeaveGameRequestedIntgrEventName:
				event, err := intgrevent.Unmarshal[intgrevent.LeaveGameRequestedIntgrEvent](message)
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
