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
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/usermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/service"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type Service interface {
	CreateGame(userIdDto string) error
	GetPlayersAroundPlayer(presenter Presenter, query GetPlayersQuery)
	GetUnitsVisibleByPlayer(presenter Presenter, query GetUnitsVisibleByPlayerQuery)
	AddPlayer(presenter Presenter, command AddPlayerCommand)
	MovePlayer(presenter Presenter, command MovePlayerCommand)
	RemovePlayer(presenter Presenter, command RemovePlayerCommand)
	PlaceItem(presenter Presenter, command PlaceItemCommand)
	DestroyItem(presenter Presenter, command DestroyItemCommand)
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

func (serve *serve) presentError(presenter Presenter, err error) {
	presenter.OnMessage(ErroredResponseDto{
		Type:          ErroredResponseDtoType,
		ClientMessage: err.Error(),
	})
}

func (serve *serve) publishUnitsUpdatedEventTo(gameId gamemodel.GameIdVo, playerId playermodel.PlayerIdVo) {
	serve.IntEventPublisher.Publish(
		NewUnitsUpdatedIntEventChannel(gameId.ToString(), playerId.ToString()),
		UnitsUpdatedIntEvent{},
	)
}

func (serve *serve) publishUnitsUpdatedEventToNearPlayers(playerId playermodel.PlayerIdVo) {
	player, err := serve.playerRepo.Get(playerId)
	if err != nil {
		return
	}

	players, err := serve.playerRepo.GetPlayersAround(player.GetGameId(), player.GetPosition())
	if err != nil {
		return
	}

	lo.ForEach(players, func(player playermodel.PlayerAgg, _ int) {
		serve.publishUnitsUpdatedEventTo(player.GetGameId(), player.GetId())
	})
}

func (serve *serve) publishPlayersUpdatedEventToNearPlayers(gameId gamemodel.GameIdVo, playerId playermodel.PlayerIdVo) {
	player, err := serve.playerRepo.Get(playerId)
	if err != nil {
		return
	}

	players, err := serve.playerRepo.GetPlayersAround(gameId, player.GetPosition())
	if err != nil {
		return
	}

	lo.ForEach(players, func(player playermodel.PlayerAgg, _ int) {
		serve.IntEventPublisher.Publish(
			NewPlayersUpdatedIntEventChannel(gameId.ToString(), player.GetId().ToString()),
			PlayersUpdatedIntEvent{},
		)
	})
}

func (serve *serve) GetPlayersAroundPlayer(presenter Presenter, query GetPlayersQuery) {
	gameId, playerId, err := query.Validate()
	if err != nil {
		serve.presentError(presenter, err)
	}

	unlocker := serve.gameRepo.LockAccess(gameId)
	defer unlocker()

	player, err := serve.playerRepo.Get(playerId)
	if err != nil {
		serve.presentError(presenter, err)
	}
	players, err := serve.playerRepo.GetPlayersAround(gameId, player.GetPosition())
	if err != nil {
		serve.presentError(presenter, err)
	}

	presenter.OnMessage(PlayersUpdatedResponseDto{
		Type: PlayersUpdatedResponseDtoType,
		Players: lo.Map(players, func(player playermodel.PlayerAgg, _ int) dto.PlayerDto {
			return dto.NewPlayerDto(player)
		}),
	})
}

func (serve *serve) GetUnitsVisibleByPlayer(presenter Presenter, query GetUnitsVisibleByPlayerQuery) {
	gameId, playerId, err := query.Validate()
	if err != nil {
		serve.presentError(presenter, err)
	}

	unlocker := serve.gameRepo.LockAccess(gameId)
	defer unlocker()

	player, err := serve.playerRepo.Get(playerId)
	if err != nil {
		serve.presentError(presenter, err)
	}

	playerVisionBound := player.GetVisionBound()
	units, err := serve.unitRepo.GetUnitsInBound(gameId, playerVisionBound)
	if err != nil {
		serve.presentError(presenter, err)
	}

	presenter.OnMessage(UnitsUpdatedResponseDto{
		Type:        UnitsUpdatedResponseDtoType,
		VisionBound: dto.NewBoundDto(playerVisionBound),
		Units: lo.Map(units, func(unit unitmodel.UnitAgg, _ int) dto.UnitDto {
			return dto.NewUnitDto(unit)
		}),
	})
}

func (serve *serve) CreateGame(userIdDto string) error {
	userId, err := usermodel.ParseUserIdVo(userIdDto)
	if err != nil {
		return err
	}

	gameId, _ := gamemodel.NewGameIdVo(uuid.New().String())

	game, _ := serve.gameRepo.GetByUserId(userId)
	if game != nil {
		return nil
	}

	items := serve.itemRepo.GetAll()
	tool.RangeMatrix(100, 100, func(x int, z int) {
		randomInt := rand.Intn(100)
		position := commonmodel.NewPositionVo(x-50, z-50)
		if randomInt < 3 {
			newUnit := unitmodel.NewUnitAgg(gameId, position, items[randomInt].GetId())
			serve.unitRepo.Add(newUnit)
		}
	})

	newGame := gamemodel.NewGameAgg(gameId, userId)

	err = serve.gameRepo.Add(newGame)
	if err != nil {
		return err
	}

	return nil
}

func (serve *serve) AddPlayer(presenter Presenter, command AddPlayerCommand) {
	gameId, playerId, err := command.Validate()
	if err != nil {
		serve.presentError(presenter, err)
	}

	unlocker := serve.gameRepo.LockAccess(gameId)
	defer unlocker()

	direction, _ := commonmodel.NewDirectionVo(2)
	newPlayer := playermodel.NewPlayerAgg(playerId, gameId, "Hello", commonmodel.NewPositionVo(0, 0), direction)

	err = serve.playerRepo.Add(newPlayer)
	if err != nil {
		serve.presentError(presenter, err)
	}

	items := serve.itemRepo.GetAll()
	itemDtos := lo.Map(items, func(item itemmodel.ItemAgg, _ int) dto.ItemDto {
		return dto.NewItemDto(item)
	})

	players, err := serve.playerRepo.GetPlayersAround(gameId, newPlayer.GetPosition())
	if err != nil {
		serve.presentError(presenter, err)
	}

	playerDtos := lo.Map(players, func(p playermodel.PlayerAgg, _ int) dto.PlayerDto {
		return dto.NewPlayerDto(p)
	})

	playerVisionBound := newPlayer.GetVisionBound()
	units, err := serve.unitRepo.GetUnitsInBound(gameId, playerVisionBound)
	if err != nil {
		serve.presentError(presenter, err)
	}

	presenter.OnMessage(GameJoinedResponseDto{
		Type:        GameJoinedResponseDtoType,
		Items:       itemDtos,
		PlayerId:    playerId.ToString(),
		Players:     playerDtos,
		VisionBound: dto.NewBoundDto(playerVisionBound),
		Units: lo.Map(units, func(unit unitmodel.UnitAgg, _ int) dto.UnitDto {
			return dto.NewUnitDto(unit)
		}),
	})

	serve.publishPlayersUpdatedEventToNearPlayers(gameId, playerId)
}

func (serve *serve) MovePlayer(presenter Presenter, command MovePlayerCommand) {
	gameId, playerId, direction, err := command.Validate()
	if err != nil {
		serve.presentError(presenter, err)
	}

	unlocker := serve.gameRepo.LockAccess(gameId)
	defer unlocker()

	err = serve.gameService.MovePlayer(gameId, playerId, direction)
	if err != nil {
		serve.presentError(presenter, err)
	}

	serve.publishPlayersUpdatedEventToNearPlayers(gameId, playerId)
	serve.publishUnitsUpdatedEventTo(gameId, playerId)
}

func (serve *serve) RemovePlayer(presenter Presenter, command RemovePlayerCommand) {
	gameId, playerId, err := command.Validate()
	if err != nil {
		serve.presentError(presenter, err)
	}

	unlocker := serve.gameRepo.LockAccess(gameId)
	defer unlocker()

	err = serve.playerRepo.Delete(playerId)
	if err != nil {
		serve.presentError(presenter, err)
	}

	serve.publishPlayersUpdatedEventToNearPlayers(gameId, playerId)
}

func (serve *serve) PlaceItem(presenter Presenter, command PlaceItemCommand) {
	gameId, playerId, itemId, err := command.Validate()
	if err != nil {
		serve.presentError(presenter, err)
	}

	unlocker := serve.gameRepo.LockAccess(gameId)
	defer unlocker()

	err = serve.gameService.PlaceItem(gameId, playerId, itemId)
	if err != nil {
		serve.presentError(presenter, err)
	}

	serve.publishUnitsUpdatedEventToNearPlayers(playerId)
}

func (serve *serve) DestroyItem(presenter Presenter, command DestroyItemCommand) {
	gameId, playerId, err := command.Validate()
	if err != nil {
		serve.presentError(presenter, err)
	}

	unlocker := serve.gameRepo.LockAccess(gameId)
	defer unlocker()

	err = serve.gameService.DestroyItem(gameId, playerId)
	if err != nil {
		serve.presentError(presenter, err)
	}

	serve.publishUnitsUpdatedEventToNearPlayers(playerId)
}
