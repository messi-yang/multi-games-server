package livegameintgreventcontroller

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent/addplayerrequestedintgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent/buliditemrequestedintgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent/destroyitemrequestedintgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent/removeplayerrequestedintgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent/zoommaprangerequestedintgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/interface/messaging/redisintgreventsubscriber"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/gameserver/application/service/livegameappservice"
)

func New(liveGameAppService livegameappservice.Service) {
	liveGameAdminChannelUnsubscriber := redisintgreventsubscriber.New().Subscribe(
		intgrevent.CreateLiveGameAdminChannel(),
		func(message []byte) {
			integrationEvent, err := intgrevent.Parse(message)
			if err != nil {
				return
			}

			if integrationEvent.Name == addplayerrequestedintgrevent.EVENT_NAME {

				event := addplayerrequestedintgrevent.Deserialize(message)
				liveGameAppService.AddPlayerToLiveGame(event.LiveGameId, event.PlayerId)

			} else if integrationEvent.Name == destroyitemrequestedintgrevent.EVENT_NAME {

				event := destroyitemrequestedintgrevent.Deserialize(message)
				liveGameAppService.DestroyItemInLiveGame(event.LiveGameId, event.Location)

			} else if integrationEvent.Name == buliditemrequestedintgrevent.EVENT_NAME {

				event := buliditemrequestedintgrevent.Deserialize(message)
				liveGameAppService.BuildItemInLiveGame(event.LiveGameId, event.Location, event.ItemId)

			} else if integrationEvent.Name == zoommaprangerequestedintgrevent.EVENT_NAME {

				event := zoommaprangerequestedintgrevent.Deserialize(message)
				liveGameAppService.AddZoomedMapRangeToLiveGame(event.LiveGameId, event.PlayerId, event.MapRange)

			} else if integrationEvent.Name == removeplayerrequestedintgrevent.EVENT_NAME {

				event := removeplayerrequestedintgrevent.Deserialize(message)
				liveGameAppService.RemovePlayerFromLiveGame(event.LiveGameId, event.PlayerId)
				liveGameAppService.RemoveZoomedMapRangeFromLiveGame(event.LiveGameId, event.PlayerId)

			}
		})
	defer liveGameAdminChannelUnsubscriber()

	closeConnFlag := make(chan bool)
	for {
		<-closeConnFlag
		return
	}
}
