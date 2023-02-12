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
	SendErroredResponseDto(presenter presenter.SocketPresenter, clientMessage string)
	HandlePlayerUpdatedEvent(presenter presenter.SocketPresenter, intEvent PlayerUpdatedIntEvent)
	HandleUnitUpdatedEvent(presenter presenter.SocketPresenter, playerIdVm string, intEvent UnitUpdatedIntEvent)
	LoadGame(gameIdVm string)
	AddPlayer(presenter presenter.SocketPresenter, command AddPlayerCommand) error
	MovePlayer(presenter presenter.SocketPresenter, command MovePlayerCommand) error
	RemovePlayer(command RemovePlayerCommand) error
	PlaceItem(command PlaceItemCommand) error
	DestroyItem(command DestroyItemCommand) error
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

func (serve *serve) SendErroredResponseDto(presenter presenter.SocketPresenter, clientMessage string) {
	presenter.OnMessage(ErroredResponseDto{
		Type:          ErroredResponseDtoType,
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

	presenter.OnMessage(PlayersUpdatedResponseDto{
		Type: PlayersUpdatedResponseDtoType,
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

	presenter.OnMessage(ViewUpdatedResponseDto{
		Type: ViewUpdatedResponseDtoType,
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

func (serve *serve) AddPlayer(presenter presenter.SocketPresenter, command AddPlayerCommand) error {
	unlocker := serve.gameRepo.LockAccess(command.GameId)
	defer unlocker()

	game, err := serve.gameRepo.Get(command.GameId)
	if err != nil {
		return err
	}

	err = game.AddPlayer(command.PlayerId)
	if err != nil {
		return err
	}

	serve.gameRepo.Update(command.GameId, game)

	items := serve.itemRepo.GetAll()
	itemVms := lo.Map(items, func(item itemmodel.ItemAgg, _ int) viewmodel.ItemVm {
		return viewmodel.NewItemVm(item)
	})

	players := game.GetPlayers()
	playerVms := lo.Map(players, func(p gamemodel.PlayerEntity, _ int) viewmodel.PlayerVm {
		return viewmodel.NewPlayerVm(p)
	})

	// Delete this section later
	bound, _ := game.GetPlayerViewBound(command.PlayerId)
	units := serve.unitRepo.GetUnits(command.GameId, bound)
	view := unitmodel.NewViewVo(bound, units)
	// Delete this section later

	presenter.OnMessage(GameJoinedResponseDto{
		Type:     GameJoinedResponseDtoType,
		Items:    itemVms,
		PlayerId: command.PlayerId.ToString(),
		Players:  playerVms,
		View:     viewmodel.NewViewVm(view),
	})

	serve.IntEventPublisher.Publish(
		CreateGameIntEventChannel(command.GameId.ToString()),
		jsonmarshaller.Marshal(NewPlayerUpdatedIntEvent(
			command.GameId.ToString(),
			command.PlayerId.ToString(),
		)))

	return nil
}

func (serve *serve) MovePlayer(presenter presenter.SocketPresenter, command MovePlayerCommand) error {
	unlocker := serve.gameRepo.LockAccess(command.GameId)
	defer unlocker()

	err := serve.gameService.MovePlayer(command.GameId, command.PlayerId, command.Direction)
	if err != nil {
		return err
	}

	game, err := serve.gameRepo.Get(command.GameId)
	if err != nil {
		return err
	}

	serve.IntEventPublisher.Publish(
		CreateGameIntEventChannel(command.GameId.ToString()),
		jsonmarshaller.Marshal(NewPlayerUpdatedIntEvent(
			command.GameId.ToString(),
			command.PlayerId.ToString(),
		)))

	// Delete this section later
	bound, _ := game.GetPlayerViewBound(command.PlayerId)
	units := serve.unitRepo.GetUnits(command.GameId, bound)
	view := unitmodel.NewViewVo(bound, units)
	// Delete this section later

	presenter.OnMessage(ViewUpdatedResponseDto{
		Type: ViewUpdatedResponseDtoType,
		View: viewmodel.NewViewVm(view),
	})

	return nil
}

func (serve *serve) RemovePlayer(command RemovePlayerCommand) error {
	unlocker := serve.gameRepo.LockAccess(command.GameId)
	defer unlocker()

	game, err := serve.gameRepo.Get(command.GameId)
	if err != nil {
		return err
	}

	game.RemovePlayer(command.PlayerId)
	serve.gameRepo.Update(command.GameId, game)

	serve.IntEventPublisher.Publish(
		CreateGameIntEventChannel(command.GameId.ToString()),
		jsonmarshaller.Marshal(NewPlayerUpdatedIntEvent(
			command.GameId.ToString(),
			command.PlayerId.ToString(),
		)))

	return nil
}

func (serve *serve) PlaceItem(command PlaceItemCommand) error {
	unlocker := serve.gameRepo.LockAccess(command.GameId)
	defer unlocker()

	err := serve.gameService.PlaceItem(command.GameId, command.PlayerId, command.ItemId, command.Location)
	if err != nil {
		return err
	}

	serve.IntEventPublisher.Publish(
		CreateGameIntEventChannel(command.GameId.ToString()),
		jsonmarshaller.Marshal(NewUnitUpdatedIntEvent(
			command.GameId.ToString(),
			viewmodel.NewLocationVm(command.Location),
		)))

	return nil
}

func (serve *serve) DestroyItem(command DestroyItemCommand) error {
	unlocker := serve.gameRepo.LockAccess(command.GameId)
	defer unlocker()

	err := serve.gameService.DestroyItem(command.GameId, command.PlayerId, command.Location)
	if err != nil {
		return err
	}

	serve.IntEventPublisher.Publish(
		CreateGameIntEventChannel(command.GameId.ToString()),
		jsonmarshaller.Marshal(NewUnitUpdatedIntEvent(
			command.GameId.ToString(),
			viewmodel.NewLocationVm(command.Location),
		)))

	return nil
}
