package gamesocketservice

import (
	"math/rand"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/intevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/library/jsonmarshaller"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/library/tool"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/presenter"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/viewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/unitmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/service"
	"github.com/samber/lo"
)

type Service interface {
	SendErroredServerEvent(presenter presenter.SocketPresenter, clientMessage string)
	HandlePlayerUpdatedEvent(presenter presenter.SocketPresenter, intEvent PlayerUpdatedIntEvent)
	HandleUnitUpdatedEvent(presenter presenter.SocketPresenter, playerIdVm string, intEvent UnitUpdatedIntEvent)
	LoadGame(gameIdVm string)
	AddPlayer(presenter presenter.SocketPresenter, gameIdVm string, playerIdVm string)
	MovePlayer(presenter presenter.SocketPresenter, gameIdVm string, playerIdVm string, directionVm int8)
	RemovePlayer(gameIdVm string, playerIdVm string)
	PlaceItem(gameIdVm string, playerIdVm string, locationVm viewmodel.LocationVm, itemIdVm int16)
	DestroyItem(gameIdVm string, playerIdVm string, locationVm viewmodel.LocationVm)
}

type serve struct {
	IntEventPublisher intevent.IntEventPublisher
	gameRepo          gamemodel.Repo
	unitRepo          unitmodel.Repo
	itemRepo          itemmodel.Repo
	gameService       service.GameService
}

func NewService(IntEventPublisher intevent.IntEventPublisher, gameRepo gamemodel.Repo, unitRepo unitmodel.Repo, itemRepo itemmodel.Repo) Service {
	return &serve{
		IntEventPublisher: IntEventPublisher,
		gameRepo:          gameRepo,
		unitRepo:          unitRepo,
		itemRepo:          itemRepo,
		gameService:       service.NewGameService(gameRepo, unitRepo, itemRepo),
	}
}

func (serve *serve) SendErroredServerEvent(presenter presenter.SocketPresenter, clientMessage string) {
	presenter.OnMessage(ErroredServerEvent{
		Type:          ErroredServerEventType,
		ClientMessage: clientMessage,
	})
}

func (serve *serve) HandlePlayerUpdatedEvent(presenter presenter.SocketPresenter, intEvent PlayerUpdatedIntEvent) {
	gameId, err := gamemodel.NewGameIdVo(intEvent.GameId)
	if err != nil {
		return
	}
	game, err := serve.gameRepo.Get(gameId)
	if err != nil {
		return
	}
	players := game.GetPlayers()

	presenter.OnMessage(PlayersUpdatedServerEvent{
		Type: PlayersUpdatedServerEventType,
		Players: lo.Map(players, func(player gamemodel.PlayerEntity, _ int) viewmodel.PlayerVm {
			return viewmodel.NewPlayerVm(player)
		}),
	})
}

func (serve *serve) HandleUnitUpdatedEvent(presenter presenter.SocketPresenter, playerIdVm string, intEvent UnitUpdatedIntEvent) {
	gameId, err := gamemodel.NewGameIdVo(intEvent.GameId)
	if err != nil {
		return
	}
	playerId, err := gamemodel.NewPlayerIdVo(playerIdVm)
	if err != nil {
		return
	}
	game, err := serve.gameRepo.Get(gameId)
	if err != nil {
		return
	}
	location := commonmodel.NewLocationVo(intEvent.Location.X, intEvent.Location.Y)
	if !game.CanPlayerSeeAnyLocations(playerId, []commonmodel.LocationVo{location}) {
		return
	}

	// Delete this section later
	bound, _ := game.GetPlayerViewBound(playerId)
	units := serve.unitRepo.GetUnits(gameId, bound)
	view := unitmodel.NewViewVo(bound, units)
	// Delete this section later

	presenter.OnMessage(ViewUpdatedServerEvent{
		Type: ViewUpdatedServerEventType,
		View: viewmodel.NewViewVm(view),
	})
}

func (serve *serve) LoadGame(gameIdVm string) {
	gameId, err := gamemodel.NewGameIdVo(gameIdVm)
	if err != nil {
		return
	}

	items := serve.itemRepo.GetAll()
	tool.RangeMatrix(200, 200, func(x int, y int) {
		randomInt := rand.Intn(17)
		location := commonmodel.NewLocationVo(x, y)
		if randomInt < 2 {
			newUnit := unitmodel.NewUnitAgg(gameId, location, items[randomInt].GetId())
			serve.unitRepo.UpdateUnit(newUnit)
		}
	})

	newGame := gamemodel.NewGameAgg(gameId)

	serve.gameRepo.Add(newGame)
}

func (serve *serve) AddPlayer(presenter presenter.SocketPresenter, gameIdVm string, playerIdVm string) {
	gameId, err := gamemodel.NewGameIdVo(gameIdVm)
	if err != nil {
		return
	}
	playerId, err := gamemodel.NewPlayerIdVo(playerIdVm)
	if err != nil {
		return
	}

	unlocker := serve.gameRepo.LockAccess(gameId)
	defer unlocker()

	game, err := serve.gameRepo.Get(gameId)
	if err != nil {
		return
	}

	err = game.AddPlayer(playerId)
	if err != nil {
		return
	}

	serve.gameRepo.Update(gameId, game)

	items := serve.itemRepo.GetAll()
	itemVms := lo.Map(items, func(item itemmodel.ItemAgg, _ int) viewmodel.ItemVm {
		return viewmodel.NewItemVm(item)
	})

	players := game.GetPlayers()
	playerVms := lo.Map(players, func(p gamemodel.PlayerEntity, _ int) viewmodel.PlayerVm {
		return viewmodel.NewPlayerVm(p)
	})

	// Delete this section later
	bound, _ := game.GetPlayerViewBound(playerId)
	units := serve.unitRepo.GetUnits(gameId, bound)
	view := unitmodel.NewViewVo(bound, units)
	// Delete this section later

	presenter.OnMessage(GameJoinedServerEvent{
		Type:     GameJoinedServerEventType,
		Items:    itemVms,
		PlayerId: playerIdVm,
		Players:  playerVms,
		View:     viewmodel.NewViewVm(view),
	})

	serve.IntEventPublisher.Publish(
		CreateGameIntEventChannel(gameIdVm),
		jsonmarshaller.Marshal(NewPlayerUpdatedIntEvent(
			gameIdVm,
			playerIdVm,
		)))
}

func (serve *serve) MovePlayer(presenter presenter.SocketPresenter, gameIdVm string, playerIdVm string, directionVm int8) {
	gameId, err := gamemodel.NewGameIdVo(gameIdVm)
	if err != nil {
		return
	}
	playerId, err := gamemodel.NewPlayerIdVo(playerIdVm)
	if err != nil {
		return
	}
	direction, err := gamemodel.NewDirectionVo(directionVm)
	if err != nil {
		return
	}

	unlocker := serve.gameRepo.LockAccess(gameId)
	defer unlocker()

	err = serve.gameService.MovePlayer(gameId, playerId, direction)
	if err != nil {
		return
	}

	game, err := serve.gameRepo.Get(gameId)
	if err != nil {
		return
	}

	serve.IntEventPublisher.Publish(
		CreateGameIntEventChannel(gameIdVm),
		jsonmarshaller.Marshal(NewPlayerUpdatedIntEvent(
			gameIdVm,
			playerIdVm,
		)))

	// Delete this section later
	bound, _ := game.GetPlayerViewBound(playerId)
	units := serve.unitRepo.GetUnits(gameId, bound)
	view := unitmodel.NewViewVo(bound, units)
	// Delete this section later

	presenter.OnMessage(ViewUpdatedServerEvent{
		Type: ViewUpdatedServerEventType,
		View: viewmodel.NewViewVm(view),
	})
}

func (serve *serve) RemovePlayer(gameIdVm string, playerIdVm string) {
	gameId, err := gamemodel.NewGameIdVo(gameIdVm)
	if err != nil {
		return
	}
	playerId, err := gamemodel.NewPlayerIdVo(playerIdVm)
	if err != nil {
		return
	}

	unlocker := serve.gameRepo.LockAccess(gameId)
	defer unlocker()

	game, err := serve.gameRepo.Get(gameId)
	if err != nil {
		return
	}

	game.RemovePlayer(playerId)
	serve.gameRepo.Update(gameId, game)

	serve.IntEventPublisher.Publish(
		CreateGameIntEventChannel(gameIdVm),
		jsonmarshaller.Marshal(NewPlayerUpdatedIntEvent(
			gameIdVm,
			playerIdVm,
		)))
}

func (serve *serve) PlaceItem(gameIdVm string, playerIdVm string, locationVm viewmodel.LocationVm, itemIdVm int16) {
	gameId, err := gamemodel.NewGameIdVo(gameIdVm)
	if err != nil {
		return
	}
	playerId, err := gamemodel.NewPlayerIdVo(playerIdVm)
	if err != nil {
		return
	}
	itemId := itemmodel.NewItemIdVo(itemIdVm)
	location := commonmodel.NewLocationVo(locationVm.X, locationVm.Y)

	unlocker := serve.gameRepo.LockAccess(gameId)
	defer unlocker()

	err = serve.gameService.PlaceItem(gameId, playerId, itemId, location)
	if err != nil {
		return
	}

	serve.IntEventPublisher.Publish(
		CreateGameIntEventChannel(gameIdVm),
		jsonmarshaller.Marshal(NewUnitUpdatedIntEvent(
			gameIdVm,
			locationVm,
		)))
}

func (serve *serve) DestroyItem(gameIdVm string, playerIdVm string, locationVm viewmodel.LocationVm) {
	gameId, err := gamemodel.NewGameIdVo(gameIdVm)
	if err != nil {
		return
	}
	playerId, err := gamemodel.NewPlayerIdVo(playerIdVm)
	if err != nil {
		return
	}
	location := commonmodel.NewLocationVo(locationVm.X, locationVm.Y)

	unlocker := serve.gameRepo.LockAccess(gameId)
	defer unlocker()

	serve.gameService.DestroyItem(gameId, playerId, location)

	serve.IntEventPublisher.Publish(
		CreateGameIntEventChannel(gameIdVm),
		jsonmarshaller.Marshal(NewUnitUpdatedIntEvent(
			gameIdVm,
			locationVm,
		)))
}
