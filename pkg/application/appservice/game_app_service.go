package appservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/intevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/viewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/unitmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/library/jsonmarshaller"
	"github.com/samber/lo"
)

type GameAppService interface {
	SendErroredServerEvent(presenter Presenter, clientMessage string)
	SendItemsUpdatedServerEvent(presenter Presenter)
	SendPlayersUpdatedServerEvent(presenter Presenter, myPlayerVm viewmodel.PlayerVm, otherPlayerVms []viewmodel.PlayerVm)
	SendViewUpdatedServerEvent(presenter Presenter, viewVm viewmodel.ViewVm)
	JoinGame(presenter Presenter, gameIdVm string, playerIdVm string)
	RequestToMove(gameIdVm string, playerIdVm string, directionVm int8)
	RequestToPlaceItem(gameIdVm string, playerIdVm string, locationVm viewmodel.LocationVm, itemIdVm int16)
	RequestToDestroyItem(gameIdVm string, playerIdVm string, locationVm viewmodel.LocationVm)
	RequestToLeaveGame(gameIdVm string, playerIdVm string)
}

type gameAppServe struct {
	IntEventPublisher intevent.IntEventPublisher
	gameRepo          gamemodel.Repo
	unitRepo          unitmodel.Repo
	itemRepo          itemmodel.Repo
}

func NewGameAppService(IntEventPublisher intevent.IntEventPublisher, gameRepo gamemodel.Repo, unitRepo unitmodel.Repo, itemRepo itemmodel.Repo) GameAppService {
	return &gameAppServe{IntEventPublisher: IntEventPublisher, gameRepo: gameRepo, unitRepo: unitRepo, itemRepo: itemRepo}
}

func (gameAppServe *gameAppServe) SendErroredServerEvent(presenter Presenter, clientMessage string) {
	event := ErroredServerEvent{}
	event.Type = ErroredServerEventType
	event.Payload.ClientMessage = clientMessage
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

func (gameAppServe *gameAppServe) JoinGame(presenter Presenter, gameIdVm string, playerIdVm string) {
	gameId, err := gamemodel.NewGameIdVo(gameIdVm)
	if err != nil {
		return
	}
	playerId, err := gamemodel.NewPlayerIdVo(playerIdVm)
	if err != nil {
		return
	}

	unlocker := gameAppServe.gameRepo.LockAccess(gameId)
	defer unlocker()

	game, err := gameAppServe.gameRepo.Get(gameId)
	if err != nil {
		return
	}

	err = game.AddPlayer(playerId)
	if err != nil {
		return
	}

	gameAppServe.gameRepo.Update(gameId, game)

	players := game.GetPlayers()
	playerVms := lo.Map(players, func(p gamemodel.PlayerEntity, _ int) viewmodel.PlayerVm {
		return viewmodel.NewPlayerVm(p)
	})
	myPlayerVm, exists := lo.Find(playerVms, func(playerVm viewmodel.PlayerVm) bool {
		return playerVm.Id == playerIdVm
	})
	if !exists {
		return
	}

	otherPlayerVms := lo.Filter(playerVms, func(playerVm viewmodel.PlayerVm, _ int) bool {
		return playerVm.Id != playerIdVm
	})

	// Delete this section later
	bound, _ := game.GetPlayerViewBound(playerId)
	units := gameAppServe.unitRepo.GetUnits(gameId, bound)
	view := unitmodel.NewViewVo(bound, units)
	// Delete this section later

	event := GameJoinedServerEvent{}
	event.Type = GameJoinedServerEventType
	event.Payload.MyPlayer = myPlayerVm
	event.Payload.OtherPlayers = otherPlayerVms
	event.Payload.MapSize = viewmodel.NewSizeVm(game.GetMapSize())
	event.Payload.View = viewmodel.NewViewVm(view)
	presenter.OnSuccess(event)
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
