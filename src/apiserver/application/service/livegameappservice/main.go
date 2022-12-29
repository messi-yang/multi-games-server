package livegameappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/integrationevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/integrationevent/addplayerrequestedintegrationevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/integrationevent/buliditemrequestedintegrationevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/integrationevent/destroyitemrequestedintegrationevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/integrationevent/removeplayerrequestedintegrationevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/integrationevent/zoomarearequestedintegrationevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/livegamemodel"
)

type Service interface {
	RequestToAddPlayer(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId)
	RequestToZoomArea(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId, area commonmodel.Area)
	RequestToBuildItem(liveGameId livegamemodel.LiveGameId, coordinate commonmodel.Coordinate, itemId itemmodel.ItemId)
	RequestToDestroyItem(liveGameId livegamemodel.LiveGameId, coordinate commonmodel.Coordinate)
	RequestToRemovePlayer(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId)
}

type serve struct {
	notificationPublisher integrationevent.Publisher
}

func New(notificationPublisher integrationevent.Publisher) Service {
	return &serve{notificationPublisher: notificationPublisher}
}

func (serve *serve) RequestToAddPlayer(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId) {
	serve.notificationPublisher.Publish(
		integrationevent.CreateLiveGameAdminChannel(),
		addplayerrequestedintegrationevent.New(liveGameId, playerId).Serialize(),
	)
}

func (serve *serve) RequestToZoomArea(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId, area commonmodel.Area) {
	serve.notificationPublisher.Publish(
		integrationevent.CreateLiveGameAdminChannel(),
		zoomarearequestedintegrationevent.New(liveGameId, playerId, area).Serialize(),
	)
}

func (serve *serve) RequestToBuildItem(liveGameId livegamemodel.LiveGameId, coordinate commonmodel.Coordinate, itemId itemmodel.ItemId) {
	serve.notificationPublisher.Publish(
		integrationevent.CreateLiveGameAdminChannel(),
		buliditemrequestedintegrationevent.New(liveGameId, coordinate, itemId).Serialize(),
	)
}

func (serve *serve) RequestToDestroyItem(liveGameId livegamemodel.LiveGameId, coordinate commonmodel.Coordinate) {
	serve.notificationPublisher.Publish(
		integrationevent.CreateLiveGameAdminChannel(),
		destroyitemrequestedintegrationevent.New(liveGameId, coordinate).Serialize(),
	)
}

func (serve *serve) RequestToRemovePlayer(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId) {
	serve.notificationPublisher.Publish(
		integrationevent.CreateLiveGameAdminChannel(),
		removeplayerrequestedintegrationevent.New(liveGameId, playerId).Serialize(),
	)
}
