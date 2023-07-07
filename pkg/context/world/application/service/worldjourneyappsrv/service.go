package worldjourneyappsrv

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/commonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/itemmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel/playermodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/service"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

var (
	errPositionHasPlayers     = fmt.Errorf("this position has players")
	errPlayerExceededBoundary = fmt.Errorf("player exceeded the boundary of the world")
)

type Service interface {
	GetPlayers(GetPlayersQuery) (playerDtos []dto.PlayerDto, err error)
	GetUnits(GetUnitsQuery) (unitDtos []dto.UnitDto, err error)
	GetUnit(GetUnitQuery) (dto.UnitDto, error)
	GetPlayer(GetPlayerQuery) (dto.PlayerDto, error)
	EnterWorld(EnterWorldCommand) (playerId uuid.UUID, err error)
	Move(MoveCommand) error
	LeaveWorld(LeaveWorldCommand) error
	ChangeHeldItem(ChangeHeldItemCommand) error
	CreateUnit(CreateUnitCommand) error
	RemoveUnit(RemoveUnitCommand) error
}

type serve struct {
	worldRepo   worldmodel.WorldRepo
	playerRepo  playermodel.PlayerRepo
	unitRepo    unitmodel.UnitRepo
	itemRepo    itemmodel.ItemRepo
	unitService service.UnitService
}

func NewService(
	worldRepo worldmodel.WorldRepo,
	playerRepo playermodel.PlayerRepo,
	unitRepo unitmodel.UnitRepo,
	itemRepo itemmodel.ItemRepo,
	unitService service.UnitService,
) Service {
	return &serve{
		worldRepo:   worldRepo,
		playerRepo:  playerRepo,
		unitRepo:    unitRepo,
		itemRepo:    itemRepo,
		unitService: unitService,
	}
}

func (serve *serve) GetPlayers(query GetPlayersQuery) (
	playerDtos []dto.PlayerDto, err error,
) {
	players, err := serve.playerRepo.GetPlayersOfWorld(sharedkernelmodel.NewWorldId(query.WorldId))
	if err != nil {
		return playerDtos, err
	}
	playerDtos = lo.Map(players, func(_player playermodel.Player, _ int) dto.PlayerDto {
		return dto.NewPlayerDto(_player)
	})

	return playerDtos, nil
}

func (serve *serve) GetUnits(query GetUnitsQuery) (
	unitDtos []dto.UnitDto, err error,
) {
	units, err := serve.unitRepo.GetUnitsOfWorld(sharedkernelmodel.NewWorldId(query.WorldId))
	if err != nil {
		return unitDtos, err
	}
	unitDtos = lo.Map(units, func(unit unitmodel.Unit, _ int) dto.UnitDto {
		return dto.NewUnitDto(unit)
	})

	return unitDtos, err
}

func (serve *serve) GetUnit(query GetUnitQuery) (unitDto dto.UnitDto, err error) {
	worldId := sharedkernelmodel.NewWorldId(query.WorldId)
	position := commonmodel.NewPosition(query.Position.X, query.Position.Z)
	unit, err := serve.unitRepo.Get(unitmodel.NewUnitId(worldId, position))
	if err != nil {
		return unitDto, err
	}
	return dto.NewUnitDto(unit), nil
}

func (serve *serve) GetPlayer(query GetPlayerQuery) (playerDto dto.PlayerDto, err error) {
	player, err := serve.playerRepo.Get(sharedkernelmodel.NewWorldId(query.WorldId), playermodel.NewPlayerId(query.PlayerId))
	if err != nil {
		return playerDto, err
	}
	return dto.NewPlayerDto(player), nil
}

func (serve *serve) EnterWorld(command EnterWorldCommand) (plyaerIdDto uuid.UUID, err error) {
	worldId := sharedkernelmodel.NewWorldId(command.WorldId)

	firstItem, err := serve.itemRepo.GetFirstItem()
	if err != nil {
		return plyaerIdDto, err
	}
	firstItemId := firstItem.GetId()

	direction := commonmodel.NewDownDirection()
	newPlayer := playermodel.NewPlayer(
		playermodel.NewPlayerId(uuid.New()),
		worldId,
		"Hello",
		commonmodel.NewPosition(0, 0),
		direction,
		&firstItemId,
	)

	if err = serve.playerRepo.Add(newPlayer); err != nil {
		return plyaerIdDto, err
	}
	return newPlayer.GetId().Uuid(), nil
}

func (serve *serve) Move(command MoveCommand) error {
	worldId := sharedkernelmodel.NewWorldId(command.WorldId)
	playerId := playermodel.NewPlayerId(command.PlayerId)
	direction := commonmodel.NewDirection(command.Direction)

	player, err := serve.playerRepo.Get(worldId, playerId)
	if err != nil {
		return err
	}

	if !direction.IsEqual(player.GetDirection()) {
		player.Move(player.GetPosition(), direction)
		return serve.playerRepo.Update(player)
	}

	newItemPos := player.GetPositionOneStepFoward()

	player.Move(newItemPos, direction)

	return serve.playerRepo.Update(player)
}

func (serve *serve) LeaveWorld(command LeaveWorldCommand) error {
	worldId := sharedkernelmodel.NewWorldId(command.WorldId)
	playerId := playermodel.NewPlayerId(command.PlayerId)

	player, err := serve.playerRepo.Get(worldId, playerId)
	if err != nil {
		return err
	}
	return serve.playerRepo.Delete(player)
}

func (serve *serve) ChangeHeldItem(command ChangeHeldItemCommand) error {
	worldId := sharedkernelmodel.NewWorldId(command.WorldId)
	playerId := playermodel.NewPlayerId(command.PlayerId)
	itemId := commonmodel.NewItemId(command.ItemId)

	player, err := serve.playerRepo.Get(worldId, playerId)
	if err != nil {
		return err
	}

	player.ChangeHeldItem(itemId)
	return serve.playerRepo.Update(player)
}

func (serve *serve) CreateUnit(command CreateUnitCommand) error {
	worldId := sharedkernelmodel.NewWorldId(command.WorldId)
	position := commonmodel.NewPosition(command.Position.X, command.Position.Z)
	players, err := serve.playerRepo.GetPlayersAt(worldId, position)
	if err != nil {
		return err
	}

	if len(players) > 0 {
		return errPositionHasPlayers
	}

	return serve.unitService.CreateUnit(
		worldId,
		commonmodel.NewItemId(command.ItemId),
		position,
		commonmodel.NewDirection(command.Direction),
	)
}

func (serve *serve) RemoveUnit(command RemoveUnitCommand) error {
	return serve.unitService.RemoveUnit(
		sharedkernelmodel.NewWorldId(command.WorldId),
		commonmodel.NewPosition(command.Position.X, command.Position.Z),
	)
}
