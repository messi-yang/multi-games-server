package unitappsrv

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/itemmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/service"
	"github.com/samber/lo"
)

type Service interface {
	GetUnits(GetUnitsQuery) (unitDtos []dto.UnitDto, err error)
	GetUnit(GetUnitQuery) (dto.UnitDto, error)

	RotateUnit(RotateUnitCommand) error

	CreateStaticUnit(CreateStaticUnitCommand) error
	RemoveStaticUnit(RemoveStaticUnitCommand) error

	CreatePortalUnit(CreatePortalUnitCommand) error
	RemovePortalUnit(RemovePortalUnitCommand) error
}

type serve struct {
	worldRepo         worldmodel.WorldRepo
	unitRepo          unitmodel.UnitRepo
	itemRepo          itemmodel.ItemRepo
	staticUnitService service.StaticUnitService
	portalUnitService service.PortalUnitService
}

func NewService(
	worldRepo worldmodel.WorldRepo,
	unitRepo unitmodel.UnitRepo,
	itemRepo itemmodel.ItemRepo,
	staticUnitService service.StaticUnitService,
	portalUnitService service.PortalUnitService,
) Service {
	return &serve{
		worldRepo:         worldRepo,
		unitRepo:          unitRepo,
		itemRepo:          itemRepo,
		staticUnitService: staticUnitService,
		portalUnitService: portalUnitService,
	}
}

func (serve *serve) GetUnits(query GetUnitsQuery) (
	unitDtos []dto.UnitDto, err error,
) {
	units, err := serve.unitRepo.GetUnitsOfWorld(globalcommonmodel.NewWorldId(query.WorldId))
	if err != nil {
		return unitDtos, err
	}
	unitDtos = lo.Map(units, func(unit unitmodel.Unit, _ int) dto.UnitDto {
		return dto.NewUnitDto(unit)
	})

	return unitDtos, err
}

func (serve *serve) GetUnit(query GetUnitQuery) (unitDto dto.UnitDto, err error) {
	worldId := globalcommonmodel.NewWorldId(query.WorldId)
	position := worldcommonmodel.NewPosition(query.Position.X, query.Position.Z)
	unit, err := serve.unitRepo.Get(unitmodel.NewUnitId(worldId, position))
	if err != nil {
		return unitDto, err
	}
	return dto.NewUnitDto(unit), nil
}

func (serve *serve) CreateStaticUnit(command CreateStaticUnitCommand) error {
	worldId := globalcommonmodel.NewWorldId(command.WorldId)
	position := worldcommonmodel.NewPosition(command.Position.X, command.Position.Z)

	return serve.staticUnitService.CreateStaticUnit(
		worldId,
		worldcommonmodel.NewItemId(command.ItemId),
		position,
		worldcommonmodel.NewDirection(command.Direction),
	)
}

func (serve *serve) RemoveStaticUnit(command RemoveStaticUnitCommand) error {
	worldId := globalcommonmodel.NewWorldId(command.WorldId)
	position := worldcommonmodel.NewPosition(command.Position.X, command.Position.Z)
	unitId := unitmodel.NewUnitId(worldId, position)

	return serve.staticUnitService.RemoveStaticUnit(unitId)
}

func (serve *serve) CreatePortalUnit(command CreatePortalUnitCommand) error {
	worldId := globalcommonmodel.NewWorldId(command.WorldId)
	position := worldcommonmodel.NewPosition(command.Position.X, command.Position.Z)

	return serve.portalUnitService.CreatePortalUnit(
		worldId,
		worldcommonmodel.NewItemId(command.ItemId),
		position,
		worldcommonmodel.NewDirection(command.Direction),
	)
}

func (serve *serve) RemovePortalUnit(command RemovePortalUnitCommand) error {
	worldId := globalcommonmodel.NewWorldId(command.WorldId)
	position := worldcommonmodel.NewPosition(command.Position.X, command.Position.Z)
	unitId := unitmodel.NewUnitId(worldId, position)

	return serve.portalUnitService.RemovePortalUnit(unitId)
}

func (serve *serve) RotateUnit(command RotateUnitCommand) error {
	worldId := globalcommonmodel.NewWorldId(command.WorldId)
	position := worldcommonmodel.NewPosition(command.Position.X, command.Position.Z)
	unitId := unitmodel.NewUnitId(worldId, position)
	unit, err := serve.unitRepo.Get(unitId)
	if err != nil {
		return err
	}

	if unit.GetType().IsEqual(worldcommonmodel.NewPortalUnitType()) {
		return serve.portalUnitService.RotatePortalUnit(unitId)
	} else if unit.GetType().IsEqual(worldcommonmodel.NewStaticUnitType()) {
		return serve.staticUnitService.RotateStaticUnit(unitId)
	}

	return nil
}
