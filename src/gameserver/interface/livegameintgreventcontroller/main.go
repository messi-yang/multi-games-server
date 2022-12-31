package livegameintgreventcontroller

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent/addplayerrequestedintgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent/buliditemrequestedintgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent/destroyitemrequestedintgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent/removeplayerrequestedintgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent/zoomarearequestedintgrevent"
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
				liveGameAppService.DestroyItemInLiveGame(event.LiveGameId, event.Coordinate)

			} else if integrationEvent.Name == buliditemrequestedintgrevent.EVENT_NAME {

				event := buliditemrequestedintgrevent.Deserialize(message)
				liveGameAppService.BuildItemInLiveGame(event.LiveGameId, event.Coordinate, event.ItemId)

			} else if integrationEvent.Name == zoomarearequestedintgrevent.EVENT_NAME {

				event := zoomarearequestedintgrevent.Deserialize(message)
				liveGameAppService.AddZoomedAreaToLiveGame(event.LiveGameId, event.PlayerId, event.Area)

			} else if integrationEvent.Name == removeplayerrequestedintgrevent.EVENT_NAME {

				event := removeplayerrequestedintgrevent.Deserialize(message)
				liveGameAppService.RemovePlayerFromLiveGame(event.LiveGameId, event.PlayerId)
				liveGameAppService.RemoveZoomedAreaFromLiveGame(event.LiveGameId, event.PlayerId)

			}
		})
	defer liveGameAdminChannelUnsubscriber()

	closeConnFlag := make(chan bool)
	for {
		<-closeConnFlag
		return
	}
}
