package livegameeventcontroller

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/integrationevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/integrationevent/addplayerrequestedintegrationevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/integrationevent/buliditemrequestedintegrationevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/integrationevent/destroyitemrequestedintegrationevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/integrationevent/removeplayerrequestedintegrationevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/integrationevent/zoomarearequestedintegrationevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/interface/integrationevent/redisintegrationeventsubscriber"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/gameserver/application/service/livegameappservice"
)

func New(liveGameAppService livegameappservice.Service) {
	integrationEventSubscriberUnsubscriber := redisintegrationeventsubscriber.New().Subscribe(
		integrationevent.CreateLiveGameAdminChannel(),
		func(message []byte) {
			integrationEvent := integrationevent.New(message)

			if integrationEvent.Name == addplayerrequestedintegrationevent.EVENT_NAME {
				event := addplayerrequestedintegrationevent.Deserialize(message)
				liveGameId, err := event.GetLiveGameId()
				if err != nil {
					return
				}
				playerId, err := event.GetPlayerId()
				if err != nil {
					return
				}
				liveGameAppService.AddPlayerToLiveGame(liveGameId, playerId)
			} else if integrationEvent.Name == destroyitemrequestedintegrationevent.EVENT_NAME {
				event := destroyitemrequestedintegrationevent.Deserialize(message)
				liveGameId, err := event.GetLiveGameId()
				if err != nil {
					return
				}
				coordinate, err := event.GetCoordinate()
				if err != nil {
					return
				}

				liveGameAppService.DestroyItemInLiveGame(liveGameId, coordinate)
			} else if integrationEvent.Name == buliditemrequestedintegrationevent.EVENT_NAME {
				event := buliditemrequestedintegrationevent.Deserialize(message)
				liveGameId, err := event.GetLiveGameId()
				if err != nil {
					return
				}
				coordinate, err := event.GetCoordinate()
				if err != nil {
					return
				}
				itemId, err := event.GetItemId()
				if err != nil {
					return
				}

				liveGameAppService.BuildItemInLiveGame(liveGameId, coordinate, itemId)
			} else if integrationEvent.Name == zoomarearequestedintegrationevent.EVENT_NAME {
				event := zoomarearequestedintegrationevent.Deserialize(message)
				liveGameId, err := event.GetLiveGameId()
				if err != nil {
					return
				}
				playerId, err := event.GetPlayerId()
				if err != nil {
					return
				}
				area, err := event.GetArea()
				if err != nil {
					return
				}
				liveGameAppService.AddZoomedAreaToLiveGame(liveGameId, playerId, area)
			} else if integrationEvent.Name == removeplayerrequestedintegrationevent.EVENT_NAME {
				event := removeplayerrequestedintegrationevent.Deserialize(message)
				liveGameId, err := event.GetLiveGameId()
				if err != nil {
					return
				}
				playerId, err := event.GetPlayerId()
				if err != nil {
					return
				}
				liveGameAppService.RemovePlayerFromLiveGame(liveGameId, playerId)
				liveGameAppService.RemoveZoomedAreaFromLiveGame(liveGameId, playerId)
			}
		})
	defer integrationEventSubscriberUnsubscriber()

	closeConnFlag := make(chan bool)
	for {
		<-closeConnFlag
		return
	}
}
