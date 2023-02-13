package gamesocketappservice

import (
	"math/rand"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/dto"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/intevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/tool"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/unitmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/service"
	"github.com/samber/lo"
)

type Service interface {
	CreateGame(gameIdDto string)
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
		NewViewUpdatedIntEventChannel(gameId.ToString(), playerId.ToString()),
		ViewUpdatedIntEvent{},
	)
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
			NewPlayersUpdatedIntEventChannel(gameId.ToString(), player.GetId().ToString()),
			PlayersUpdatedIntEvent{},
		)
	})
}

func (serve *serve) GetPlayers(presenter Presenter, query GetPlayersQuery) error {
	gameId, _, err := query.Validate()
	if err != nil {
		return err
	}

	unlocker := serve.gameRepo.LockAccess(gameId)
	defer unlocker()

	game, err := serve.gameRepo.Get(gameId)
	if err != nil {
		return err
	}
	players := game.GetPlayers()

	presenter.OnMessage(PlayersUpdatedResponseDto{
		Type: PlayersUpdatedResponseDtoType,
		Players: lo.Map(players, func(player gamemodel.PlayerEntity, _ int) dto.PlayerDto {
			return dto.NewPlayerDto(player)
		}),
	})

	return nil
}

func (serve *serve) GetView(presenter Presenter, query GetViewQuery) error {
	gameId, playerId, err := query.Validate()
	if err != nil {
		return err
	}

	unlocker := serve.gameRepo.LockAccess(gameId)
	defer unlocker()

	game, err := serve.gameRepo.Get(gameId)
	if err != nil {
		return err
	}

	// Delete this section later
	bound, _ := game.GetPlayerViewBound(playerId)
	units := serve.unitRepo.GetUnits(gameId, bound)
	view := unitmodel.NewViewVo(bound, units)
	// Delete this section later

	presenter.OnMessage(ViewUpdatedResponseDto{
		Type: ViewUpdatedResponseDtoType,
		View: dto.NewViewDto(view),
	})

	return nil
}

func (serve *serve) CreateGame(gameIdDto string) {
	gameId, err := gamemodel.NewGameIdVo(gameIdDto)
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
	gameId, playerId, err := command.Validate()
	if err != nil {
		return err
	}

	unlocker := serve.gameRepo.LockAccess(gameId)
	defer unlocker()

	err = serve.gameService.AddPlayer(gameId, playerId)
	if err != nil {
		return err
	}

	game, err := serve.gameRepo.Get(gameId)
	if err != nil {
		return err
	}

	items := serve.itemRepo.GetAll()
	itemDtos := lo.Map(items, func(item itemmodel.ItemAgg, _ int) dto.ItemDto {
		return dto.NewItemDto(item)
	})

	players, _ := serve.gameService.GetNearbyPlayersOfPlayer(gameId, playerId)
	playerDtos := lo.Map(players, func(p gamemodel.PlayerEntity, _ int) dto.PlayerDto {
		return dto.NewPlayerDto(p)
	})

	// Delete this section later
	bound, _ := game.GetPlayerViewBound(playerId)
	units := serve.unitRepo.GetUnits(gameId, bound)
	view := unitmodel.NewViewVo(bound, units)
	// Delete this section later

	presenter.OnMessage(GameJoinedResponseDto{
		Type:     GameJoinedResponseDtoType,
		Items:    itemDtos,
		PlayerId: playerId.ToString(),
		Players:  playerDtos,
		View:     dto.NewViewDto(view),
	})

	serve.publishPlayersUpdatedEventToNearbyPlayersOfPlayer(gameId, playerId)

	return nil
}

func (serve *serve) MovePlayer(presenter Presenter, command MovePlayerCommand) error {
	gameId, playerId, direction, err := command.Validate()
	if err != nil {
		return err
	}

	unlocker := serve.gameRepo.LockAccess(gameId)
	defer unlocker()

	err = serve.gameService.MovePlayer(gameId, playerId, direction)
	if err != nil {
		return err
	}

	serve.publishPlayersUpdatedEventToNearbyPlayersOfPlayer(gameId, playerId)
	serve.publishViewUpdatedEventTo(gameId, playerId)

	return nil
}

func (serve *serve) RemovePlayer(command RemovePlayerCommand) error {
	gameId, playerId, err := command.Validate()
	if err != nil {
		return err
	}

	unlocker := serve.gameRepo.LockAccess(gameId)
	defer unlocker()

	err = serve.gameService.RemovePlayer(gameId, playerId)
	if err != nil {
		return err
	}

	serve.publishPlayersUpdatedEventToNearbyPlayersOfPlayer(gameId, playerId)

	return nil
}

func (serve *serve) PlaceItem(command PlaceItemCommand) error {
	gameId, playerId, itemId, location, err := command.Validate()
	if err != nil {
		return err
	}

	unlocker := serve.gameRepo.LockAccess(gameId)
	defer unlocker()

	err = serve.gameService.PlaceItem(gameId, playerId, itemId, location)
	if err != nil {
		return err
	}

	serve.publishViewUpdatedEventToNearbyPlayersOfLocation(gameId, location)

	return nil
}

func (serve *serve) DestroyItem(command DestroyItemCommand) error {
	gameId, playerId, location, err := command.Validate()
	if err != nil {
		return err
	}

	unlocker := serve.gameRepo.LockAccess(gameId)
	defer unlocker()

	err = serve.gameService.DestroyItem(gameId, playerId, location)
	if err != nil {
		return err
	}

	serve.publishViewUpdatedEventToNearbyPlayersOfLocation(gameId, location)

	return nil
}
