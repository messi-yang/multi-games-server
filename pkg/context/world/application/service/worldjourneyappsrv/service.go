package worldjourneyappsrv

import (
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

type Service interface {
	GetPlayers(GetPlayersQuery) (playerDtos []dto.PlayerDto, err error)
	GetUnits(GetUnitsQuery) (unitDtos []dto.UnitDto, err error)
	GetUnit(GetUnitQuery) (dto.UnitDto, error)
	GetPlayer(GetPlayerQuery) (dto.PlayerDto, error)
	EnterWorld(EnterWorldCommand) (playerId uuid.UUID, err error)
	Move(MoveCommand) error
	LeaveWorld(LeaveWorldCommand) error
	ChangeHeldItem(ChangeHeldItemCommand) error
	PlaceUnit(PlaceUnitCommand) error
	RemoveUnit(RemoveUnitCommand) error
}

type serve struct {
	worldRepo           worldmodel.WorldRepo
	playerRepo          playermodel.PlayerRepo
	unitRepo            unitmodel.UnitRepo
	itemRepo            itemmodel.ItemRepo
	worldJourneyService service.WorldJourneyService
}

func NewService(
	worldRepo worldmodel.WorldRepo,
	playerRepo playermodel.PlayerRepo,
	unitRepo unitmodel.UnitRepo,
	itemRepo itemmodel.ItemRepo,
	worldJourneyService service.WorldJourneyService,
) Service {
	return &serve{
		worldRepo:           worldRepo,
		playerRepo:          playerRepo,
		unitRepo:            unitRepo,
		itemRepo:            itemRepo,
		worldJourneyService: worldJourneyService,
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
	player, err := serve.playerRepo.Get(playermodel.NewPlayerId(query.PlayerId))
	if err != nil {
		return playerDto, err
	}
	return dto.NewPlayerDto(player), nil
}

func (serve *serve) EnterWorld(command EnterWorldCommand) (plyaerIdDto uuid.UUID, err error) {
	playerId, err := serve.worldJourneyService.EnterWorld(sharedkernelmodel.NewWorldId(command.WorldId))
	if err != nil {
		return plyaerIdDto, err
	}
	return playerId.Uuid(), nil
}

func (serve *serve) Move(command MoveCommand) error {
	return serve.worldJourneyService.Move(sharedkernelmodel.NewWorldId(command.WorldId), playermodel.NewPlayerId(command.PlayerId), commonmodel.NewDirection(command.Direction))
}

func (serve *serve) LeaveWorld(command LeaveWorldCommand) error {
	return serve.worldJourneyService.LeaveWorld(sharedkernelmodel.NewWorldId(command.WorldId), playermodel.NewPlayerId(command.PlayerId))
}

func (serve *serve) PlaceUnit(command PlaceUnitCommand) error {
	return serve.worldJourneyService.PlaceUnit(sharedkernelmodel.NewWorldId(command.WorldId), playermodel.NewPlayerId(command.PlayerId))
}

func (serve *serve) ChangeHeldItem(command ChangeHeldItemCommand) error {
	return serve.worldJourneyService.ChangeHeldItem(sharedkernelmodel.NewWorldId(command.WorldId), playermodel.NewPlayerId(command.PlayerId), commonmodel.NewItemId(command.ItemId))
}

func (serve *serve) RemoveUnit(command RemoveUnitCommand) error {
	return serve.worldJourneyService.RemoveUnit(sharedkernelmodel.NewWorldId(command.WorldId), playermodel.NewPlayerId(command.PlayerId))
}
