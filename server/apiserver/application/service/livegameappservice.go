package service

import (
	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/livegamemodel"
	commonappevent "github.com/dum-dum-genius/game-of-liberty-computer/server/common/application/event"
	commonnotification "github.com/dum-dum-genius/game-of-liberty-computer/server/common/application/notification"
)

type LiveGameAppService interface {
	RequestToAddPlayer(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId)
	RequestToZoomArea(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId, area gamecommonmodel.Area)
	RequestToBuildItem(liveGameId livegamemodel.LiveGameId, coordinate gamecommonmodel.Coordinate, itemId itemmodel.ItemId)
	RequestToRemovePlayer(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId)
}

type liveGameAppServe struct {
	notificationPublisher commonnotification.NotificationPublisher
}

func NewLiveGameAppService(notificationPublisher commonnotification.NotificationPublisher) LiveGameAppService {
	return &liveGameAppServe{notificationPublisher: notificationPublisher}
}

func (serve *liveGameAppServe) RequestToAddPlayer(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) {
	serve.notificationPublisher.Publish(
		commonappevent.NewAddPlayerRequestedAppEventChannel(),
		commonappevent.NewAddPlayerRequestedAppEvent(liveGameId, playerId),
	)
}

func (serve *liveGameAppServe) RequestToZoomArea(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId, area gamecommonmodel.Area) {
	serve.notificationPublisher.Publish(
		commonappevent.NewZoomAreaRequestedAppEventChannel(),
		commonappevent.NewZoomAreaRequestedAppEvent(liveGameId, playerId, area),
	)
}

func (serve *liveGameAppServe) RequestToBuildItem(liveGameId livegamemodel.LiveGameId, coordinate gamecommonmodel.Coordinate, itemId itemmodel.ItemId) {
	serve.notificationPublisher.Publish(
		commonappevent.NewBuildItemRequestedAppEventChannel(),
		commonappevent.NewBuildItemRequestedAppEvent(liveGameId, coordinate, itemId),
	)
}

func (serve *liveGameAppServe) RequestToRemovePlayer(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) {
	serve.notificationPublisher.Publish(
		commonappevent.NewRemovePlayerRequestedAppEventChannel(),
		commonappevent.NewRemovePlayerRequestedAppEvent(liveGameId, playerId),
	)
}
