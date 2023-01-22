package appservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/jsonmarshaller"
	"github.com/samber/lo"
)

type LiveGameAppService interface {
	SendErroredServerEvent(presenter Presenter, clientMessage string)
	SendItemsUpdatedServerEvent(presenter Presenter)
	SendGameJoinedServerEvent(presenter Presenter, myPlayerVm viewmodel.PlayerVm, otherPlayerVms []viewmodel.PlayerVm, playerVm viewmodel.PlayerVm, mapSizeVm viewmodel.SizeVm, viewVm viewmodel.ViewVm)
	SendPlayersUpdatedServerEvent(presenter Presenter, myPlayerVm viewmodel.PlayerVm, otherPlayerVms []viewmodel.PlayerVm)
	SendViewUpdatedServerEvent(presenter Presenter, viewVm viewmodel.ViewVm)
	RequestToJoinGame(liveGameIdVm string, playerIdVm string)
	RequestToChangeCamera(liveGameIdVm string, playerIdVm string, cameraVm viewmodel.CameraVm)
	RequestToBuildItem(liveGameIdVm string, locationVm viewmodel.LocationVm, itemIdVm string)
	RequestToDestroyItem(liveGameIdVm string, locationVm viewmodel.LocationVm)
	RequestToLeaveGame(liveGameIdVm string, playerIdVm string)
}

type liveGameAppServe struct {
	IntEventPublisher intevent.IntEventPublisher
	itemRepo          itemmodel.Repo
}

func NewLiveGameAppService(IntEventPublisher intevent.IntEventPublisher, itemRepo itemmodel.Repo) LiveGameAppService {
	return &liveGameAppServe{IntEventPublisher: IntEventPublisher, itemRepo: itemRepo}
}

func (liveGameAppServe *liveGameAppServe) SendErroredServerEvent(presenter Presenter, clientMessage string) {
	event := ErroredServerEvent{}
	event.Type = ErroredServerEventType
	event.Payload.ClientMessage = clientMessage
	presenter.OnSuccess(event)
}

func (liveGameAppServe *liveGameAppServe) SendGameJoinedServerEvent(presenter Presenter, myPlayerVm viewmodel.PlayerVm, otherPlayerVms []viewmodel.PlayerVm, playerVm viewmodel.PlayerVm, mapSizeVm viewmodel.SizeVm, viewVm viewmodel.ViewVm) {
	event := GameJoinedServerEvent{}
	event.Type = GameJoinedServerEventType
	event.Payload.MyPlayer = myPlayerVm
	event.Payload.OtherPlayers = otherPlayerVms
	event.Payload.Player = playerVm
	event.Payload.MapSize = mapSizeVm
	event.Payload.View = viewVm
	presenter.OnSuccess(event)
}

func (liveGameAppServe *liveGameAppServe) SendPlayersUpdatedServerEvent(presenter Presenter, myPlayerVm viewmodel.PlayerVm, otherPlayerVms []viewmodel.PlayerVm) {
	event := PlayersUpdatedServerEvent{}
	event.Type = PlayersUpdatedServerEventType
	event.Payload.MyPlayer = myPlayerVm
	event.Payload.OtherPlayers = otherPlayerVms
	presenter.OnSuccess(event)
}

func (liveGameAppServe *liveGameAppServe) SendViewUpdatedServerEvent(presenter Presenter, viewVm viewmodel.ViewVm) {
	event := ViewUpdatedServerEvent{}
	event.Type = ViewUpdatedServerEventType
	event.Payload.View = viewVm
	presenter.OnSuccess(event)
}

func (liveGameAppServe *liveGameAppServe) SendItemsUpdatedServerEvent(presenter Presenter) {
	items := liveGameAppServe.itemRepo.GetAllItems()
	itemVms := lo.Map(items, func(item itemmodel.ItemAgg, _ int) viewmodel.ItemVm {
		return viewmodel.NewItemVm(item)
	})

	event := ItemsUpdatedServerEvent{}
	event.Type = ItemsUpdatedServerEventType
	event.Payload.Items = itemVms
	presenter.OnSuccess(event)
}

func (liveGameAppServe *liveGameAppServe) RequestToJoinGame(liveGameIdVm string, playerIdVm string) {
	liveGameAppServe.IntEventPublisher.Publish(
		intevent.CreateLiveGameAdminChannel(),
		jsonmarshaller.Marshal(intevent.NewJoinGameRequestedIntEvent(liveGameIdVm, playerIdVm)),
	)
}

func (liveGameAppServe *liveGameAppServe) RequestToChangeCamera(liveGameIdVm string, playerIdVm string, cameraVm viewmodel.CameraVm) {
	liveGameAppServe.IntEventPublisher.Publish(
		intevent.CreateLiveGameAdminChannel(),
		jsonmarshaller.Marshal(intevent.NewChangeCameraRequestedIntEvent(liveGameIdVm, playerIdVm, cameraVm)),
	)
}

func (liveGameAppServe *liveGameAppServe) RequestToBuildItem(liveGameIdVm string, locationVm viewmodel.LocationVm, itemIdVm string) {
	liveGameAppServe.IntEventPublisher.Publish(
		intevent.CreateLiveGameAdminChannel(),
		jsonmarshaller.Marshal(intevent.NewBuildItemRequestedIntEvent(liveGameIdVm, locationVm, itemIdVm)),
	)
}

func (liveGameAppServe *liveGameAppServe) RequestToDestroyItem(liveGameIdVm string, locationVm viewmodel.LocationVm) {
	liveGameAppServe.IntEventPublisher.Publish(
		intevent.CreateLiveGameAdminChannel(),
		jsonmarshaller.Marshal(intevent.NewDestroyItemRequestedIntEvent(liveGameIdVm, locationVm)),
	)
}

func (liveGameAppServe *liveGameAppServe) RequestToLeaveGame(liveGameIdVm string, playerIdVm string) {
	liveGameAppServe.IntEventPublisher.Publish(
		intevent.CreateLiveGameAdminChannel(),
		jsonmarshaller.Marshal(intevent.NewLeaveGameRequestedIntEvent(liveGameIdVm, playerIdVm)),
	)
}
