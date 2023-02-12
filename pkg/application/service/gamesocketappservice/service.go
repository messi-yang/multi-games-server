package gamesocketappservice

import (
	"math/rand"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/intevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/library/jsonmarshaller"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/library/tool"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/viewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/unitmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/service"
	"github.com/samber/lo"
)

type Service interface {
	CreateGame(gameIdVm string)
	GetError(presenter Presenter, errorMessage string)
	GetPlayers(presenter Presenter, query GetPlayersQuery) error
	GetView(presenter Presenter, query GetViewQuery) error
	AddPlayer(presenter Presenter, command AddPlayerCommand) error
	MovePlayer(presenter Presenter, command MovePlayerCommand) error
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

func (serve *serve) GetError(presenter Presenter, errorMessage string) {
	presenter.OnMessage(ErroredResponseDto{
		Type:          ErroredResponseDtoType,
		ClientMessage: errorMessage,
	})
}

func (serve *serve) publishViewUpdatedEventTo(gameId gamemodel.GameIdVo, playerId gamemodel.PlayerIdVo) {
	serve.IntEventPublisher.Publish(
		CreateGamePlayerChannel(gameId.ToString(), playerId.ToString()),
		jsonmarshaller.Marshal(NewViewUpdatedIntEvent(
			gameId.ToString(),
		)))
}

func (serve *serve) publishViewUpdatedEventToNearbyPlayersOfLocation(gameId gamemodel.GameIdVo, location commonmodel.LocationVo) {
	players, err := serve.gameService.GetNearbyPlayersOfLocation(gameId, location)
	if err != nil {
		return
	}

	lo.ForEach(players, func(player gamemodel.PlayerEntity, _ int) {
		serve.publishViewUpdatedEventTo(gameId, player.GetId())
	})
}

func (serve *serve) publishPlayersUpdatedEventToNearbyPlayersOfPlayer(gameId gamemodel.GameIdVo, playerId gamemodel.PlayerIdVo) {
	players, err := serve.gameService.GetNearbyPlayersOfPlayer(gameId, playerId)
	if err != nil {
		return
	}

	lo.ForEach(players, func(player gamemodel.PlayerEntity, _ int) {
		serve.IntEventPublisher.Publish(
			CreateGamePlayerChannel(gameId.ToString(), player.GetId().ToString()),
			jsonmarshaller.Marshal(NewPlayersUpdatedIntEvent(
				gameId.ToString(),
			)))
	})
}

func (serve *serve) GetPlayers(presenter Presenter, query GetPlayersQuery) error {
	unlocker := serve.gameRepo.LockAccess(query.GameId)
	defer unlocker()

	game, err := serve.gameRepo.Get(query.GameId)
	if err != nil {
		return err
	}
	players := game.GetPlayers()

	presenter.OnMessage(PlayersUpdatedResponseDto{
		Type: PlayersUpdatedResponseDtoType,
		Players: lo.Map(players, func(player gamemodel.PlayerEntity, _ int) viewmodel.PlayerVm {
			return viewmodel.NewPlayerVm(player)
		}),
	})

	return nil
}

func (serve *serve) GetView(presenter Presenter, query GetViewQuery) error {
	unlocker := serve.gameRepo.LockAccess(query.GameId)
	defer unlocker()

	game, err := serve.gameRepo.Get(query.GameId)
	if err != nil {
		return err
	}

	// Delete this section later
	bound, _ := game.GetPlayerViewBound(query.PlayerId)
	units := serve.unitRepo.GetUnits(query.GameId, bound)
	view := unitmodel.NewViewVo(bound, units)
	// Delete this section later

	presenter.OnMessage(ViewUpdatedResponseDto{
		Type: ViewUpdatedResponseDtoType,
		View: viewmodel.NewViewVm(view),
	})

	return nil
}

func (serve *serve) CreateGame(gameIdVm string) {
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

func (serve *serve) AddPlayer(presenter Presenter, command AddPlayerCommand) error {
	unlocker := serve.gameRepo.LockAccess(command.GameId)
	defer unlocker()

	err := serve.gameService.AddPlayer(command.GameId, command.PlayerId)
	if err != nil {
		return err
	}

	game, err := serve.gameRepo.Get(command.GameId)
	if err != nil {
		return err
	}

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

	serve.publishPlayersUpdatedEventToNearbyPlayersOfPlayer(command.GameId, command.PlayerId)

	return nil
}

func (serve *serve) MovePlayer(presenter Presenter, command MovePlayerCommand) error {
	unlocker := serve.gameRepo.LockAccess(command.GameId)
	defer unlocker()

	err := serve.gameService.MovePlayer(command.GameId, command.PlayerId, command.Direction)
	if err != nil {
		return err
	}

	serve.publishPlayersUpdatedEventToNearbyPlayersOfPlayer(command.GameId, command.PlayerId)
	serve.publishViewUpdatedEventTo(command.GameId, command.PlayerId)

	return nil
}

func (serve *serve) RemovePlayer(command RemovePlayerCommand) error {
	unlocker := serve.gameRepo.LockAccess(command.GameId)
	defer unlocker()

	err := serve.gameService.RemovePlayer(command.GameId, command.PlayerId)
	if err != nil {
		return err
	}

	serve.publishPlayersUpdatedEventToNearbyPlayersOfPlayer(command.GameId, command.PlayerId)

	return nil
}

func (serve *serve) PlaceItem(command PlaceItemCommand) error {
	unlocker := serve.gameRepo.LockAccess(command.GameId)
	defer unlocker()

	err := serve.gameService.PlaceItem(command.GameId, command.PlayerId, command.ItemId, command.Location)
	if err != nil {
		return err
	}

	serve.publishViewUpdatedEventToNearbyPlayersOfLocation(command.GameId, command.Location)

	return nil
}

func (serve *serve) DestroyItem(command DestroyItemCommand) error {
	unlocker := serve.gameRepo.LockAccess(command.GameId)
	defer unlocker()

	err := serve.gameService.DestroyItem(command.GameId, command.PlayerId, command.Location)
	if err != nil {
		return err
	}

	serve.publishViewUpdatedEventToNearbyPlayersOfLocation(command.GameId, command.Location)

	return nil
}
