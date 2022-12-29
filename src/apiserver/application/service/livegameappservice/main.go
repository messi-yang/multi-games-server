package livegameappservice

import (
	commonappevent "github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/event"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/integrationeventpublisher"
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
	notificationPublisher integrationeventpublisher.Publisher
}

func New(notificationPublisher integrationeventpublisher.Publisher) Service {
	return &serve{notificationPublisher: notificationPublisher}
}

func (serve *serve) RequestToAddPlayer(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId) {
	serve.notificationPublisher.Publish(
		commonappevent.NewAddPlayerRequestedAppEventChannel(),
		commonappevent.NewAddPlayerRequestedAppEvent(liveGameId, playerId),
	)
}

func (serve *serve) RequestToZoomArea(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId, area commonmodel.Area) {
	serve.notificationPublisher.Publish(
		commonappevent.NewZoomAreaRequestedAppEventChannel(),
		commonappevent.NewZoomAreaRequestedAppEvent(liveGameId, playerId, area),
	)
}

func (serve *serve) RequestToBuildItem(liveGameId livegamemodel.LiveGameId, coordinate commonmodel.Coordinate, itemId itemmodel.ItemId) {
	serve.notificationPublisher.Publish(
		commonappevent.NewBuildItemRequestedAppEventChannel(),
		commonappevent.NewBuildItemRequestedAppEvent(liveGameId, coordinate, itemId),
	)
}

func (serve *serve) RequestToDestroyItem(liveGameId livegamemodel.LiveGameId, coordinate commonmodel.Coordinate) {
	serve.notificationPublisher.Publish(
		commonappevent.NewDestroyItemRequestedAppEventChannel(),
		commonappevent.NewDestroyItemRequestedAppEvent(liveGameId, coordinate),
	)
}

func (serve *serve) RequestToRemovePlayer(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId) {
	serve.notificationPublisher.Publish(
		commonappevent.NewRemovePlayerRequestedAppEventChannel(),
		commonappevent.NewRemovePlayerRequestedAppEvent(liveGameId, playerId),
	)
}
