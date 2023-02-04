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
	SendGameJoinedServerEvent(presenter Presenter, myPlayerVm viewmodel.PlayerVm, otherPlayerVms []viewmodel.PlayerVm, mapSizeVm viewmodel.SizeVm, viewVm viewmodel.ViewVm)
	SendPlayersUpdatedServerEvent(presenter Presenter, myPlayerVm viewmodel.PlayerVm, otherPlayerVms []viewmodel.PlayerVm)
	SendViewUpdatedServerEvent(presenter Presenter, viewVm viewmodel.ViewVm)
	RequestToJoinGame(liveGameIdVm string, playerIdVm string)
	RequestToMove(liveGameIdVm string, playerIdVm string, directionVm int8)
	RequestToBuildItem(liveGameIdVm string, playerIdVm string, locationVm viewmodel.LocationVm, itemIdVm string)
	RequestToDestroyItem(liveGameIdVm string, playerIdVm string, locationVm viewmodel.LocationVm)
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

func (liveGameAppServe *liveGameAppServe) SendGameJoinedServerEvent(presenter Presenter, myPlayerVm viewmodel.PlayerVm, otherPlayerVms []viewmodel.PlayerVm, mapSizeVm viewmodel.SizeVm, viewVm viewmodel.ViewVm) {
	event := GameJoinedServerEvent{}
	event.Type = GameJoinedServerEventType
	event.Payload.MyPlayer = myPlayerVm
	event.Payload.OtherPlayers = otherPlayerVms
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
	items := liveGameAppServe.itemRepo.GetAll()
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

func (liveGameAppServe *liveGameAppServe) RequestToMove(liveGameIdVm string, playerIdVm string, directionVm int8) {
	liveGameAppServe.IntEventPublisher.Publish(
		intevent.CreateLiveGameAdminChannel(),
		jsonmarshaller.Marshal(intevent.NewMoveRequestedIntEvent(liveGameIdVm, playerIdVm, directionVm)),
	)
}

func (liveGameAppServe *liveGameAppServe) RequestToBuildItem(liveGameIdVm string, playerIdVm string, locationVm viewmodel.LocationVm, itemIdVm string) {
	liveGameAppServe.IntEventPublisher.Publish(
		intevent.CreateLiveGameAdminChannel(),
		jsonmarshaller.Marshal(intevent.NewBuildItemRequestedIntEvent(liveGameIdVm, playerIdVm, locationVm, itemIdVm)),
	)
}

func (liveGameAppServe *liveGameAppServe) RequestToDestroyItem(liveGameIdVm string, playerIdVm string, locationVm viewmodel.LocationVm) {
	liveGameAppServe.IntEventPublisher.Publish(
		intevent.CreateLiveGameAdminChannel(),
		jsonmarshaller.Marshal(intevent.NewDestroyItemRequestedIntEvent(liveGameIdVm, playerIdVm, locationVm)),
	)
}

func (liveGameAppServe *liveGameAppServe) RequestToLeaveGame(liveGameIdVm string, playerIdVm string) {
	liveGameAppServe.IntEventPublisher.Publish(
		intevent.CreateLiveGameAdminChannel(),
		jsonmarshaller.Marshal(intevent.NewLeaveGameRequestedIntEvent(liveGameIdVm, playerIdVm)),
	)
}
