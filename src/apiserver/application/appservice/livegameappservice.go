package appservice

import (
	commonappevent "github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/event"
	commonnotification "github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/notification"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/domainmodel/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/domainmodel/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/domainmodel/livegamemodel"
)

type LiveGameAppService interface {
	RequestToAddPlayer(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId)
	RequestToZoomArea(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId, area commonmodel.Area)
	RequestToBuildItem(liveGameId livegamemodel.LiveGameId, coordinate commonmodel.Coordinate, itemId itemmodel.ItemId)
	RequestToDestroyItem(liveGameId livegamemodel.LiveGameId, coordinate commonmodel.Coordinate)
	RequestToRemovePlayer(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId)
}

type liveGameAppServe struct {
	notificationPublisher commonnotification.NotificationPublisher
}

func NewLiveGameAppService(notificationPublisher commonnotification.NotificationPublisher) LiveGameAppService {
	return &liveGameAppServe{notificationPublisher: notificationPublisher}
}

func (serve *liveGameAppServe) RequestToAddPlayer(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId) {
	serve.notificationPublisher.Publish(
		commonappevent.NewAddPlayerRequestedAppEventChannel(),
		commonappevent.NewAddPlayerRequestedAppEvent(liveGameId, playerId),
	)
}

func (serve *liveGameAppServe) RequestToZoomArea(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId, area commonmodel.Area) {
	serve.notificationPublisher.Publish(
		commonappevent.NewZoomAreaRequestedAppEventChannel(),
		commonappevent.NewZoomAreaRequestedAppEvent(liveGameId, playerId, area),
	)
}

func (serve *liveGameAppServe) RequestToBuildItem(liveGameId livegamemodel.LiveGameId, coordinate commonmodel.Coordinate, itemId itemmodel.ItemId) {
	serve.notificationPublisher.Publish(
		commonappevent.NewBuildItemRequestedAppEventChannel(),
		commonappevent.NewBuildItemRequestedAppEvent(liveGameId, coordinate, itemId),
	)
}

func (serve *liveGameAppServe) RequestToDestroyItem(liveGameId livegamemodel.LiveGameId, coordinate commonmodel.Coordinate) {
	serve.notificationPublisher.Publish(
		commonappevent.NewDestroyItemRequestedAppEventChannel(),
		commonappevent.NewDestroyItemRequestedAppEvent(liveGameId, coordinate),
	)
}

func (serve *liveGameAppServe) RequestToRemovePlayer(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId) {
	serve.notificationPublisher.Publish(
		commonappevent.NewRemovePlayerRequestedAppEventChannel(),
		commonappevent.NewRemovePlayerRequestedAppEvent(liveGameId, playerId),
	)
}
