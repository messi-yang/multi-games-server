package gamesocketappservice

import (
	"math/rand"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/dto"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/intevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/tool"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/playermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/unitmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/usermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/worldmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/service"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type Service interface {
	CreateWorld(userIdDto string) error
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
	worldRepo         worldmodel.Repo
	playerRepo        playermodel.Repo
	unitRepo          unitmodel.Repo
	itemRepo          itemmodel.Repo
	gameService       service.GameService
}

func NewService(IntEventPublisher intevent.IntEventPublisher, worldRepo worldmodel.Repo, playerRepo playermodel.Repo, unitRepo unitmodel.Repo, itemRepo itemmodel.Repo) Service {
	return &serve{
		IntEventPublisher: IntEventPublisher,
		worldRepo:         worldRepo,
		playerRepo:        playerRepo,
		unitRepo:          unitRepo,
		itemRepo:          itemRepo,
		gameService:       service.NewGameService(worldRepo, playerRepo, unitRepo, itemRepo),
	}
}

func (serve *serve) presentError(presenter Presenter, err error) {
	presenter.OnMessage(ErroredResponseDto{
		Type:          ErroredResponseDtoType,
		ClientMessage: err.Error(),
	})
}

func (serve *serve) publishUnitsUpdatedEventTo(worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo) {
	serve.IntEventPublisher.Publish(
		NewUnitsUpdatedIntEventChannel(worldId.String(), playerId.String()),
		UnitsUpdatedIntEvent{},
	)
}

func (serve *serve) publishUnitsUpdatedEventToNearPlayers(playerId playermodel.PlayerIdVo) {
	player, err := serve.playerRepo.Get(playerId)
	if err != nil {
		return
	}

	players, err := serve.playerRepo.GetPlayersAround(player.GetWorldId(), player.GetPosition())
	if err != nil {
		return
	}

	lo.ForEach(players, func(player playermodel.PlayerAgg, _ int) {
		serve.publishUnitsUpdatedEventTo(player.GetWorldId(), player.GetId())
	})
}

func (serve *serve) publishPlayersUpdatedEventToNearPlayers(worldId worldmodel.WorldIdVo, playerId playermodel.PlayerIdVo) {
	player, err := serve.playerRepo.Get(playerId)
	if err != nil {
		return
	}

	players, err := serve.playerRepo.GetPlayersAround(worldId, player.GetPosition())
	if err != nil {
		return
	}

	lo.ForEach(players, func(player playermodel.PlayerAgg, _ int) {
		serve.IntEventPublisher.Publish(
			NewPlayersUpdatedIntEventChannel(worldId.String(), player.GetId().String()),
			PlayersUpdatedIntEvent{},
		)
	})
}

func (serve *serve) GetPlayersAroundPlayer(presenter Presenter, query GetPlayersQuery) {
	worldId, playerId, err := query.Validate()
	if err != nil {
		serve.presentError(presenter, err)
	}

	unlocker := serve.worldRepo.LockAccess(worldId)
	defer unlocker()

	player, err := serve.playerRepo.Get(playerId)
	if err != nil {
		serve.presentError(presenter, err)
	}
	players, err := serve.playerRepo.GetPlayersAround(worldId, player.GetPosition())
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
	worldId, playerId, err := query.Validate()
	if err != nil {
		serve.presentError(presenter, err)
	}

	unlocker := serve.worldRepo.LockAccess(worldId)
	defer unlocker()

	player, err := serve.playerRepo.Get(playerId)
	if err != nil {
		serve.presentError(presenter, err)
	}

	playerVisionBound := player.GetVisionBound()
	units, err := serve.unitRepo.GetUnitsInBound(worldId, playerVisionBound)
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

func (serve *serve) CreateWorld(userIdDto string) error {
	userId, err := usermodel.ParseUserIdVo(userIdDto)
	if err != nil {
		return err
	}

	worldId, _ := worldmodel.ParseWorldIdVo(uuid.New().String())

	world, _ := serve.worldRepo.GetByUserId(userId)
	if world != nil {
		return nil
	}

	items := serve.itemRepo.GetAll()
	tool.RangeMatrix(100, 100, func(x int, z int) {
		randomInt := rand.Intn(100)
		position := commonmodel.NewPositionVo(x-50, z-50)
		if randomInt < 3 {
			newUnit := unitmodel.NewUnitAgg(worldId, position, items[randomInt].GetId())
			serve.unitRepo.Add(newUnit)
		}
	})

	newWorld := worldmodel.NewWorldAgg(worldId, userId)

	err = serve.worldRepo.Add(newWorld)
	if err != nil {
		return err
	}

	return nil
}

func (serve *serve) AddPlayer(presenter Presenter, command AddPlayerCommand) {
	worldId, playerId, err := command.Validate()
	if err != nil {
		serve.presentError(presenter, err)
	}

	unlocker := serve.worldRepo.LockAccess(worldId)
	defer unlocker()

	direction, _ := commonmodel.NewDirectionVo(2)
	newPlayer := playermodel.NewPlayerAgg(playerId, worldId, "Hello", commonmodel.NewPositionVo(0, 0), direction)

	err = serve.playerRepo.Add(newPlayer)
	if err != nil {
		serve.presentError(presenter, err)
	}

	items := serve.itemRepo.GetAll()
	itemDtos := lo.Map(items, func(item itemmodel.ItemAgg, _ int) dto.ItemDto {
		return dto.NewItemDto(item)
	})

	players, err := serve.playerRepo.GetPlayersAround(worldId, newPlayer.GetPosition())
	if err != nil {
		serve.presentError(presenter, err)
	}

	playerDtos := lo.Map(players, func(p playermodel.PlayerAgg, _ int) dto.PlayerDto {
		return dto.NewPlayerDto(p)
	})

	playerVisionBound := newPlayer.GetVisionBound()
	units, err := serve.unitRepo.GetUnitsInBound(worldId, playerVisionBound)
	if err != nil {
		serve.presentError(presenter, err)
	}

	presenter.OnMessage(GameJoinedResponseDto{
		Type:        GameJoinedResponseDtoType,
		Items:       itemDtos,
		PlayerId:    playerId.String(),
		Players:     playerDtos,
		VisionBound: dto.NewBoundDto(playerVisionBound),
		Units: lo.Map(units, func(unit unitmodel.UnitAgg, _ int) dto.UnitDto {
			return dto.NewUnitDto(unit)
		}),
	})

	serve.publishPlayersUpdatedEventToNearPlayers(worldId, playerId)
}

func (serve *serve) MovePlayer(presenter Presenter, command MovePlayerCommand) {
	worldId, playerId, direction, err := command.Validate()
	if err != nil {
		serve.presentError(presenter, err)
	}

	unlocker := serve.worldRepo.LockAccess(worldId)
	defer unlocker()

	err = serve.gameService.MovePlayer(worldId, playerId, direction)
	if err != nil {
		serve.presentError(presenter, err)
	}

	serve.publishPlayersUpdatedEventToNearPlayers(worldId, playerId)
	serve.publishUnitsUpdatedEventTo(worldId, playerId)
}

func (serve *serve) RemovePlayer(presenter Presenter, command RemovePlayerCommand) {
	worldId, playerId, err := command.Validate()
	if err != nil {
		serve.presentError(presenter, err)
	}

	unlocker := serve.worldRepo.LockAccess(worldId)
	defer unlocker()

	err = serve.playerRepo.Delete(playerId)
	if err != nil {
		serve.presentError(presenter, err)
	}

	serve.publishPlayersUpdatedEventToNearPlayers(worldId, playerId)
}

func (serve *serve) PlaceItem(presenter Presenter, command PlaceItemCommand) {
	worldId, playerId, itemId, err := command.Validate()
	if err != nil {
		serve.presentError(presenter, err)
	}

	unlocker := serve.worldRepo.LockAccess(worldId)
	defer unlocker()

	err = serve.gameService.PlaceItem(worldId, playerId, itemId)
	if err != nil {
		serve.presentError(presenter, err)
	}

	serve.publishUnitsUpdatedEventToNearPlayers(playerId)
}

func (serve *serve) DestroyItem(presenter Presenter, command DestroyItemCommand) {
	worldId, playerId, err := command.Validate()
	if err != nil {
		serve.presentError(presenter, err)
	}

	unlocker := serve.worldRepo.LockAccess(worldId)
	defer unlocker()

	err = serve.gameService.DestroyItem(worldId, playerId)
	if err != nil {
		serve.presentError(presenter, err)
	}

	serve.publishUnitsUpdatedEventToNearPlayers(playerId)
}
