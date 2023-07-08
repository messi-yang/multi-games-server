package worldjourneyappsrv

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/commonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/itemmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/service"
	"github.com/samber/lo"
)

type Service interface {
	GetUnits(GetUnitsQuery) (unitDtos []dto.UnitDto, err error)
	GetUnit(GetUnitQuery) (dto.UnitDto, error)
	CreateUnit(CreateUnitCommand) error
	RemoveUnit(RemoveUnitCommand) error
}

type serve struct {
	worldRepo   worldmodel.WorldRepo
	unitRepo    unitmodel.UnitRepo
	itemRepo    itemmodel.ItemRepo
	unitService service.UnitService
}

func NewService(
	worldRepo worldmodel.WorldRepo,
	unitRepo unitmodel.UnitRepo,
	itemRepo itemmodel.ItemRepo,
	unitService service.UnitService,
) Service {
	return &serve{
		worldRepo:   worldRepo,
		unitRepo:    unitRepo,
		itemRepo:    itemRepo,
		unitService: unitService,
	}
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

func (serve *serve) CreateUnit(command CreateUnitCommand) error {
	worldId := sharedkernelmodel.NewWorldId(command.WorldId)
	position := commonmodel.NewPosition(command.Position.X, command.Position.Z)

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
