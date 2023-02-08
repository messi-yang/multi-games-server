package appservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/jsonmarshaller"
	"github.com/samber/lo"
)

type GameAppService interface {
	SendErroredServerEvent(presenter Presenter, clientMessage string)
	SendItemsUpdatedServerEvent(presenter Presenter)
	SendGameJoinedServerEvent(presenter Presenter, myPlayerVm viewmodel.PlayerVm, otherPlayerVms []viewmodel.PlayerVm, mapSizeVm viewmodel.SizeVm, viewVm viewmodel.ViewVm)
	SendPlayersUpdatedServerEvent(presenter Presenter, myPlayerVm viewmodel.PlayerVm, otherPlayerVms []viewmodel.PlayerVm)
	SendViewUpdatedServerEvent(presenter Presenter, viewVm viewmodel.ViewVm)
	RequestToJoinGame(gameIdVm string, playerIdVm string)
	RequestToMove(gameIdVm string, playerIdVm string, directionVm int8)
	RequestToPlaceItem(gameIdVm string, playerIdVm string, locationVm viewmodel.LocationVm, itemIdVm int16)
	RequestToDestroyItem(gameIdVm string, playerIdVm string, locationVm viewmodel.LocationVm)
	RequestToLeaveGame(gameIdVm string, playerIdVm string)
}

type gameAppServe struct {
	IntEventPublisher intevent.IntEventPublisher
	itemRepo          itemmodel.Repo
}

func NewGameAppService(IntEventPublisher intevent.IntEventPublisher, itemRepo itemmodel.Repo) GameAppService {
	return &gameAppServe{IntEventPublisher: IntEventPublisher, itemRepo: itemRepo}
}

func (gameAppServe *gameAppServe) SendErroredServerEvent(presenter Presenter, clientMessage string) {
	event := ErroredServerEvent{}
	event.Type = ErroredServerEventType
	event.Payload.ClientMessage = clientMessage
	presenter.OnSuccess(event)
}

func (gameAppServe *gameAppServe) SendGameJoinedServerEvent(presenter Presenter, myPlayerVm viewmodel.PlayerVm, otherPlayerVms []viewmodel.PlayerVm, mapSizeVm viewmodel.SizeVm, viewVm viewmodel.ViewVm) {
	event := GameJoinedServerEvent{}
	event.Type = GameJoinedServerEventType
	event.Payload.MyPlayer = myPlayerVm
	event.Payload.OtherPlayers = otherPlayerVms
	event.Payload.MapSize = mapSizeVm
	event.Payload.View = viewVm
	presenter.OnSuccess(event)
}

func (gameAppServe *gameAppServe) SendPlayersUpdatedServerEvent(presenter Presenter, myPlayerVm viewmodel.PlayerVm, otherPlayerVms []viewmodel.PlayerVm) {
	event := PlayersUpdatedServerEvent{}
	event.Type = PlayersUpdatedServerEventType
	event.Payload.MyPlayer = myPlayerVm
	event.Payload.OtherPlayers = otherPlayerVms
	presenter.OnSuccess(event)
}

func (gameAppServe *gameAppServe) SendViewUpdatedServerEvent(presenter Presenter, viewVm viewmodel.ViewVm) {
	event := ViewUpdatedServerEvent{}
	event.Type = ViewUpdatedServerEventType
	event.Payload.View = viewVm
	presenter.OnSuccess(event)
}

func (gameAppServe *gameAppServe) SendItemsUpdatedServerEvent(presenter Presenter) {
	items := gameAppServe.itemRepo.GetAll()
	itemVms := lo.Map(items, func(item itemmodel.ItemAgg, _ int) viewmodel.ItemVm {
		return viewmodel.NewItemVm(item)
	})

	event := ItemsUpdatedServerEvent{}
	event.Type = ItemsUpdatedServerEventType
	event.Payload.Items = itemVms
	presenter.OnSuccess(event)
}

func (gameAppServe *gameAppServe) RequestToJoinGame(gameIdVm string, playerIdVm string) {
	gameAppServe.IntEventPublisher.Publish(
		intevent.CreateGameAdminChannel(),
		jsonmarshaller.Marshal(intevent.NewJoinGameRequestedIntEvent(gameIdVm, playerIdVm)),
	)
}

func (gameAppServe *gameAppServe) RequestToMove(gameIdVm string, playerIdVm string, directionVm int8) {
	gameAppServe.IntEventPublisher.Publish(
		intevent.CreateGameAdminChannel(),
		jsonmarshaller.Marshal(intevent.NewMoveRequestedIntEvent(gameIdVm, playerIdVm, directionVm)),
	)
}

func (gameAppServe *gameAppServe) RequestToPlaceItem(gameIdVm string, playerIdVm string, locationVm viewmodel.LocationVm, itemIdVm int16) {
	gameAppServe.IntEventPublisher.Publish(
		intevent.CreateGameAdminChannel(),
		jsonmarshaller.Marshal(intevent.NewPlaceItemRequestedIntEvent(gameIdVm, playerIdVm, locationVm, itemIdVm)),
	)
}

func (gameAppServe *gameAppServe) RequestToDestroyItem(gameIdVm string, playerIdVm string, locationVm viewmodel.LocationVm) {
	gameAppServe.IntEventPublisher.Publish(
		intevent.CreateGameAdminChannel(),
		jsonmarshaller.Marshal(intevent.NewDestroyItemRequestedIntEvent(gameIdVm, playerIdVm, locationVm)),
	)
}

func (gameAppServe *gameAppServe) RequestToLeaveGame(gameIdVm string, playerIdVm string) {
	gameAppServe.IntEventPublisher.Publish(
		intevent.CreateGameAdminChannel(),
		jsonmarshaller.Marshal(intevent.NewLeaveGameRequestedIntEvent(gameIdVm, playerIdVm)),
	)
}
