package gamesocketappservice

import (
	"math/rand"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/dto"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/intevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/tool"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/playermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/unitmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/service"
	"github.com/samber/lo"
)

type Service interface {
	CreateGame(gameIdDto string)
	GetError(presenter Presenter, errorMessage string)
	GetPlayers(presenter Presenter, query GetPlayersQuery) error
	GetUnitsAroundPlayer(presenter Presenter, query GetUnitsAroundPlayerQuery) error
	AddPlayer(presenter Presenter, command AddPlayerCommand) error
	MovePlayer(presenter Presenter, command MovePlayerCommand) error
	RemovePlayer(command RemovePlayerCommand) error
	PlaceItem(command PlaceItemCommand) error
	DestroyItem(command DestroyItemCommand) error
}

type serve struct {
	IntEventPublisher intevent.IntEventPublisher
	gameRepo          gamemodel.Repo
	playerRepo        playermodel.Repo
	unitRepo          unitmodel.Repo
	itemRepo          itemmodel.Repo
	gameService       service.GameService
}

func NewService(IntEventPublisher intevent.IntEventPublisher, gameRepo gamemodel.Repo, playerRepo playermodel.Repo, unitRepo unitmodel.Repo, itemRepo itemmodel.Repo) Service {
	return &serve{
		IntEventPublisher: IntEventPublisher,
		gameRepo:          gameRepo,
		playerRepo:        playerRepo,
		unitRepo:          unitRepo,
		itemRepo:          itemRepo,
		gameService:       service.NewGameService(gameRepo, playerRepo, unitRepo, itemRepo),
	}
}

func (serve *serve) GetError(presenter Presenter, errorMessage string) {
	presenter.OnMessage(ErroredResponseDto{
		Type:          ErroredResponseDtoType,
		ClientMessage: errorMessage,
	})
}

func (serve *serve) publishUnitsUpdatedEventTo(gameId gamemodel.GameIdVo, playerId playermodel.PlayerIdVo) {
	serve.IntEventPublisher.Publish(
		NewUnitsUpdatedIntEventChannel(gameId.ToString(), playerId.ToString()),
		UnitsUpdatedIntEvent{},
	)
}

func (serve *serve) publishUnitsUpdatedEventToNearPlayers(gameId gamemodel.GameIdVo, location commonmodel.LocationVo) {
	players := serve.playerRepo.GetPlayersAround(gameId, location)

	lo.ForEach(players, func(player playermodel.PlayerAgg, _ int) {
		serve.publishUnitsUpdatedEventTo(gameId, player.GetId())
	})
}

func (serve *serve) publishPlayersUpdatedEventToNearPlayers(gameId gamemodel.GameIdVo, playerId playermodel.PlayerIdVo) {
	player, err := serve.playerRepo.Get(playerId)
	if err != nil {
		return
	}

	players := serve.playerRepo.GetPlayersAround(gameId, player.GetLocation())

	lo.ForEach(players, func(player playermodel.PlayerAgg, _ int) {
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

	players := serve.playerRepo.GetAll(gameId)

	presenter.OnMessage(PlayersUpdatedResponseDto{
		Type: PlayersUpdatedResponseDtoType,
		Players: lo.Map(players, func(player playermodel.PlayerAgg, _ int) dto.PlayerDto {
			return dto.NewPlayerDto(player)
		}),
	})

	return nil
}

func (serve *serve) GetUnitsAroundPlayer(presenter Presenter, query GetUnitsAroundPlayerQuery) error {
	gameId, playerId, err := query.Validate()
	if err != nil {
		return err
	}

	unlocker := serve.gameRepo.LockAccess(gameId)
	defer unlocker()

	player, err := serve.playerRepo.Get(playerId)
	if err != nil {
		return err
	}

	bound := player.GetVisionBound()
	units := serve.unitRepo.GetUnits(gameId, bound)

	presenter.OnMessage(UnitsUpdatedResponseDto{
		Type:  UnitsUpdatedResponseDtoType,
		Bound: dto.NewBoundDto(bound),
		Units: lo.Map(units, func(unit unitmodel.UnitAgg, _ int) dto.UnitDto {
			return dto.NewUnitDto(unit)
		}),
	})

	return nil
}

func (serve *serve) CreateGame(gameIdDto string) {
	gameId, err := gamemodel.NewGameIdVo(gameIdDto)
	if err != nil {
		return
	}

	items := serve.itemRepo.GetAll()
	tool.RangeMatrix(200, 200, func(x int, z int) {
		randomInt := rand.Intn(17)
		location := commonmodel.NewLocationVo(x, z)
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

	newPlayer := playermodel.NewPlayerAgg(playerId, gameId, "Hello", commonmodel.NewLocationVo(0, 0))

	err = serve.playerRepo.Add(newPlayer)
	if err != nil {
		return err
	}

	items := serve.itemRepo.GetAll()
	itemDtos := lo.Map(items, func(item itemmodel.ItemAgg, _ int) dto.ItemDto {
		return dto.NewItemDto(item)
	})

	players := serve.playerRepo.GetPlayersAround(gameId, newPlayer.GetLocation())
	playerDtos := lo.Map(players, func(p playermodel.PlayerAgg, _ int) dto.PlayerDto {
		return dto.NewPlayerDto(p)
	})

	bound := newPlayer.GetVisionBound()
	units := serve.unitRepo.GetUnits(gameId, bound)

	presenter.OnMessage(GameJoinedResponseDto{
		Type:     GameJoinedResponseDtoType,
		Items:    itemDtos,
		PlayerId: playerId.ToString(),
		Players:  playerDtos,
		Bound:    dto.NewBoundDto(bound),
		Units: lo.Map(units, func(unit unitmodel.UnitAgg, _ int) dto.UnitDto {
			return dto.NewUnitDto(unit)
		}),
	})

	serve.publishPlayersUpdatedEventToNearPlayers(gameId, playerId)

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

	serve.publishPlayersUpdatedEventToNearPlayers(gameId, playerId)
	serve.publishUnitsUpdatedEventTo(gameId, playerId)

	return nil
}

func (serve *serve) RemovePlayer(command RemovePlayerCommand) error {
	gameId, playerId, err := command.Validate()
	if err != nil {
		return err
	}

	unlocker := serve.gameRepo.LockAccess(gameId)
	defer unlocker()

	serve.playerRepo.Delete(playerId)

	serve.publishPlayersUpdatedEventToNearPlayers(gameId, playerId)

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

	serve.publishUnitsUpdatedEventToNearPlayers(gameId, location)

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

	serve.publishUnitsUpdatedEventToNearPlayers(gameId, location)

	return nil
}
