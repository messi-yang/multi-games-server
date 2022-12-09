package service

import (
	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/livegamemodel"
	commonapplicationevent "github.com/dum-dum-genius/game-of-liberty-computer/server/common/application/event"
	commonnotification "github.com/dum-dum-genius/game-of-liberty-computer/server/common/application/notification"
)

type LiveGameApplicationService interface {
	RequestToAddPlayer(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId)
	RequestToZoomArea(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId, area gamecommonmodel.Area)
	RequestToReviveUnits(liveGameId livegamemodel.LiveGameId, coordinates []gamecommonmodel.Coordinate)
	RequestToRemovePlayer(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId)
}

type liveGameApplicationServe struct {
	notificationPublisher commonnotification.NotificationPublisher
}

func NewLiveGameApplicationService(notificationPublisher commonnotification.NotificationPublisher) LiveGameApplicationService {
	return &liveGameApplicationServe{notificationPublisher: notificationPublisher}
}

func (serve *liveGameApplicationServe) RequestToAddPlayer(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) {
	serve.notificationPublisher.Publish(
		commonapplicationevent.NewAddPlayerRequestedApplicationEventChannel(),
		commonapplicationevent.NewAddPlayerRequestedApplicationEvent(liveGameId, playerId),
	)
}

func (serve *liveGameApplicationServe) RequestToZoomArea(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId, area gamecommonmodel.Area) {
	serve.notificationPublisher.Publish(
		commonapplicationevent.NewZoomAreaRequestedApplicationEventChannel(),
		commonapplicationevent.NewZoomAreaRequestedApplicationEvent(liveGameId, playerId, area),
	)
}

func (serve *liveGameApplicationServe) RequestToReviveUnits(liveGameId livegamemodel.LiveGameId, coordinates []gamecommonmodel.Coordinate) {
	serve.notificationPublisher.Publish(
		commonapplicationevent.NewReviveUnitsRequestedApplicationEventChannel(),
		commonapplicationevent.NewReviveUnitsRequestedApplicationEvent(liveGameId, coordinates),
	)
}

func (serve *liveGameApplicationServe) RequestToRemovePlayer(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) {
	serve.notificationPublisher.Publish(
		commonapplicationevent.NewRemovePlayerRequestedApplicationEventChannel(),
		commonapplicationevent.NewRemovePlayerRequestedApplicationEvent(liveGameId, playerId),
	)
}
