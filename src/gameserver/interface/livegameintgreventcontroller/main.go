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
			case intgrevent.AddPlayerRequestedIntgrEventName:
				event, err := intgrevent.Unmarshal[intgrevent.AddPlayerRequestedIntgrEvent](message)
				if err != nil {
					return
				}
				liveGameAppService.AddPlayerToLiveGame(event.LiveGameId, event.PlayerId)
			case intgrevent.DestroyItemRequestedEventName:
				event, err := intgrevent.Unmarshal[intgrevent.DestroyItemRequestedEvent](message)
				if err != nil {
					return
				}
				liveGameAppService.DestroyItemInLiveGame(event.LiveGameId, event.Location)
			case intgrevent.BuildItemRequestedIntgrEventName:
				event, err := intgrevent.Unmarshal[intgrevent.BuildItemRequestedIntgrEvent](message)
				if err != nil {
					return
				}
				liveGameAppService.BuildItemInLiveGame(event.LiveGameId, event.Location, event.ItemId)
			case intgrevent.ObserveMapRangeRequestedEventName:
				event, err := intgrevent.Unmarshal[intgrevent.ObserveMapRangeRequestedEvent](message)
				if err != nil {
					return
				}
				liveGameAppService.AddObservedMapRangeToLiveGame(event.LiveGameId, event.PlayerId, event.MapRange)
			case intgrevent.RemovePlayerRequestedEventName:
				event, err := intgrevent.Unmarshal[intgrevent.RemovePlayerRequestedEvent](message)
				if err != nil {
					return
				}
				liveGameAppService.RemovePlayerFromLiveGame(event.LiveGameId, event.PlayerId)
				liveGameAppService.RemoveObservedMapRangeFromLiveGame(event.LiveGameId, event.PlayerId)
			}
		})
	defer liveGameAdminChannelUnsubscriber()

	closeConnFlag := make(chan bool)
	for {
		<-closeConnFlag
		return
	}
}
