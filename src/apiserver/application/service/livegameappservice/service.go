package livegameappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
	"github.com/samber/lo"
)

type Service interface {
	QueryItems(presenter Presenter)
	SendErroredServerEvent(presenter Presenter, clientMessage string)
	SendGameJoinedServerEvent(presenter Presenter, playerIdVm string, cameraVm viewmodel.CameraVm, dimensionVm viewmodel.DimensionVm, viewVm viewmodel.ViewVm)
	SendCameraChangedServerEvent(presenter Presenter, cameraVm viewmodel.CameraVm, viewVm viewmodel.ViewVm)
	SendViewUpdatedServerEvent(presenter Presenter, cameraVm viewmodel.CameraVm, viewVm viewmodel.ViewVm)
	RequestToJoinGame(liveGameIdVm string, playerIdVm string)
	RequestToChangeCamera(liveGameIdVm string, playerIdVm string, cameraVm viewmodel.CameraVm)
	RequestToBuildItem(liveGameIdVm string, locationVm viewmodel.LocationVm, itemIdVm string)
	RequestToDestroyItem(liveGameIdVm string, locationVm viewmodel.LocationVm)
	RequestToLeaveGame(liveGameIdVm string, playerIdVm string)
}

type serve struct {
	intgrEventPublisher intgrevent.IntgrEventPublisher
	itemRepo            itemmodel.Repo
}

func New(intgrEventPublisher intgrevent.IntgrEventPublisher, itemRepo itemmodel.Repo) Service {
	return &serve{intgrEventPublisher: intgrEventPublisher, itemRepo: itemRepo}
}

func (serve *serve) SendErroredServerEvent(presenter Presenter, clientMessage string) {
	event := ErroredServerEvent{}
	event.Type = ErroredServerEventType
	event.Payload.ClientMessage = clientMessage
	presenter.OnSuccess(event)
}

func (serve *serve) SendGameJoinedServerEvent(presenter Presenter, playerIdVm string, cameraVm viewmodel.CameraVm, dimensionVm viewmodel.DimensionVm, viewVm viewmodel.ViewVm) {
	event := GameJoinedServerEvent{}
	event.Type = GameJoinedServerEventType
	event.Payload.PlayerId = playerIdVm
	event.Payload.Camera = cameraVm
	event.Payload.Dimension = dimensionVm
	event.Payload.View = viewVm
	presenter.OnSuccess(event)
}

func (serve *serve) SendCameraChangedServerEvent(presenter Presenter, cameraVm viewmodel.CameraVm, viewVm viewmodel.ViewVm) {
	event := CameraChangedServerEvent{}
	event.Type = CameraChangedServerEventType
	event.Payload.Camera = cameraVm
	event.Payload.View = viewVm
	presenter.OnSuccess(event)
}

func (serve *serve) SendViewUpdatedServerEvent(presenter Presenter, cameraVm viewmodel.CameraVm, viewVm viewmodel.ViewVm) {
	event := ViewUpdatedServerEvent{}
	event.Type = ViewUpdatedServerEventType
	event.Payload.Camera = cameraVm
	event.Payload.View = viewVm
	presenter.OnSuccess(event)
}

func (serve *serve) QueryItems(presenter Presenter) {
	items := serve.itemRepo.GetAllItems()
	itemVms := lo.Map(items, func(item itemmodel.Item, _ int) viewmodel.ItemVm {
		return viewmodel.NewItemVm(item)
	})

	event := ItemsUpdatedServerEvent{}
	event.Type = ItemsUpdatedServerEventType
	event.Payload.Items = itemVms
	presenter.OnSuccess(event)
}

func (serve *serve) RequestToJoinGame(liveGameIdVm string, playerIdVm string) {
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameAdminChannel(),
		intgrevent.Marshal(intgrevent.NewJoinGameRequestedIntgrEvent(liveGameIdVm, playerIdVm)),
	)
}

func (serve *serve) RequestToChangeCamera(liveGameIdVm string, playerIdVm string, cameraVm viewmodel.CameraVm) {
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameAdminChannel(),
		intgrevent.Marshal(intgrevent.NewChangeCameraRequestedIntgrEvent(liveGameIdVm, playerIdVm, cameraVm)),
	)
}

func (serve *serve) RequestToBuildItem(liveGameIdVm string, locationVm viewmodel.LocationVm, itemIdVm string) {
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameAdminChannel(),
		intgrevent.Marshal(intgrevent.NewBuildItemRequestedIntgrEvent(liveGameIdVm, locationVm, itemIdVm)),
	)
}

func (serve *serve) RequestToDestroyItem(liveGameIdVm string, locationVm viewmodel.LocationVm) {
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameAdminChannel(),
		intgrevent.Marshal(intgrevent.NewDestroyItemRequestedIntgrEvent(liveGameIdVm, locationVm)),
	)
}

func (serve *serve) RequestToLeaveGame(liveGameIdVm string, playerIdVm string) {
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameAdminChannel(),
		intgrevent.Marshal(intgrevent.NewLeaveGameRequestedIntgrEvent(liveGameIdVm, playerIdVm)),
	)
}
