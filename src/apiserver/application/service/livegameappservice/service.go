package livegameappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/jsonmarshaller"
	"github.com/samber/lo"
)

type Service interface {
	SendErroredServerEvent(presenter Presenter, clientMessage string)
	SendItemsUpdatedServerEvent(presenter Presenter)
	SendGameJoinedServerEvent(presenter Presenter, playerVm viewmodel.PlayerVm, cameraVm viewmodel.CameraVm, mapSizeVm viewmodel.SizeVm, viewVm viewmodel.ViewVm)
	SendCameraChangedServerEvent(presenter Presenter, cameraVm viewmodel.CameraVm)
	SendViewChangedServerEvent(presenter Presenter, viewVm viewmodel.ViewVm)
	SendViewUpdatedServerEvent(presenter Presenter, viewVm viewmodel.ViewVm)
	RequestToJoinGame(liveGameIdVm string, playerIdVm string)
	RequestToChangeCamera(liveGameIdVm string, playerIdVm string, cameraVm viewmodel.CameraVm)
	RequestToBuildItem(liveGameIdVm string, locationVm viewmodel.LocationVm, itemIdVm string)
	RequestToDestroyItem(liveGameIdVm string, locationVm viewmodel.LocationVm)
	RequestToLeaveGame(liveGameIdVm string, playerIdVm string)
}

type serve struct {
	IntEventPublisher intevent.IntEventPublisher
	itemRepo          itemmodel.Repo
}

func New(IntEventPublisher intevent.IntEventPublisher, itemRepo itemmodel.Repo) Service {
	return &serve{IntEventPublisher: IntEventPublisher, itemRepo: itemRepo}
}

func (serve *serve) SendErroredServerEvent(presenter Presenter, clientMessage string) {
	event := ErroredServerEvent{}
	event.Type = ErroredServerEventType
	event.Payload.ClientMessage = clientMessage
	presenter.OnSuccess(event)
}

func (serve *serve) SendGameJoinedServerEvent(presenter Presenter, playerVm viewmodel.PlayerVm, cameraVm viewmodel.CameraVm, mapSizeVm viewmodel.SizeVm, viewVm viewmodel.ViewVm) {
	event := GameJoinedServerEvent{}
	event.Type = GameJoinedServerEventType
	event.Payload.Player = playerVm
	event.Payload.Camera = cameraVm
	event.Payload.MapSize = mapSizeVm
	event.Payload.View = viewVm
	presenter.OnSuccess(event)
}

func (serve *serve) SendCameraChangedServerEvent(presenter Presenter, cameraVm viewmodel.CameraVm) {
	event := CameraChangedServerEvent{}
	event.Type = CameraChangedServerEventType
	event.Payload.Camera = cameraVm
	presenter.OnSuccess(event)
}

func (serve *serve) SendViewChangedServerEvent(presenter Presenter, viewVm viewmodel.ViewVm) {
	event := ViewChangedServerEvent{}
	event.Type = ViewChangedServerEventType
	event.Payload.View = viewVm
	presenter.OnSuccess(event)
}

func (serve *serve) SendViewUpdatedServerEvent(presenter Presenter, viewVm viewmodel.ViewVm) {
	event := ViewUpdatedServerEvent{}
	event.Type = ViewUpdatedServerEventType
	event.Payload.View = viewVm
	presenter.OnSuccess(event)
}

func (serve *serve) SendItemsUpdatedServerEvent(presenter Presenter) {
	items := serve.itemRepo.GetAllItems()
	itemVms := lo.Map(items, func(item itemmodel.ItemAgg, _ int) viewmodel.ItemVm {
		return viewmodel.NewItemVm(item)
	})

	event := ItemsUpdatedServerEvent{}
	event.Type = ItemsUpdatedServerEventType
	event.Payload.Items = itemVms
	presenter.OnSuccess(event)
}

func (serve *serve) RequestToJoinGame(liveGameIdVm string, playerIdVm string) {
	serve.IntEventPublisher.Publish(
		intevent.CreateLiveGameAdminChannel(),
		jsonmarshaller.Marshal(intevent.NewJoinGameRequestedintEvent(liveGameIdVm, playerIdVm)),
	)
}

func (serve *serve) RequestToChangeCamera(liveGameIdVm string, playerIdVm string, cameraVm viewmodel.CameraVm) {
	serve.IntEventPublisher.Publish(
		intevent.CreateLiveGameAdminChannel(),
		jsonmarshaller.Marshal(intevent.NewChangeCameraRequestedintEvent(liveGameIdVm, playerIdVm, cameraVm)),
	)
}

func (serve *serve) RequestToBuildItem(liveGameIdVm string, locationVm viewmodel.LocationVm, itemIdVm string) {
	serve.IntEventPublisher.Publish(
		intevent.CreateLiveGameAdminChannel(),
		jsonmarshaller.Marshal(intevent.NewBuildItemRequestedintEvent(liveGameIdVm, locationVm, itemIdVm)),
	)
}

func (serve *serve) RequestToDestroyItem(liveGameIdVm string, locationVm viewmodel.LocationVm) {
	serve.IntEventPublisher.Publish(
		intevent.CreateLiveGameAdminChannel(),
		jsonmarshaller.Marshal(intevent.NewDestroyItemRequestedintEvent(liveGameIdVm, locationVm)),
	)
}

func (serve *serve) RequestToLeaveGame(liveGameIdVm string, playerIdVm string) {
	serve.IntEventPublisher.Publish(
		intevent.CreateLiveGameAdminChannel(),
		jsonmarshaller.Marshal(intevent.NewLeaveGameRequestedintEvent(liveGameIdVm, playerIdVm)),
	)
}
