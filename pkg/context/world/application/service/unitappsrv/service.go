package unitappsrv

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/itemmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/portalunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/staticunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/service"
	"github.com/samber/lo"
)

type Service interface {
	GetUnits(GetUnitsQuery) (unitDtos []dto.UnitDto, err error)
	GetUnit(GetUnitQuery) (dto.UnitDto, error)
	RotateUnit(RotateUnitCommand) error
	RemoveUnit(RemoveUnitCommand) error

	CreateStaticUnit(CreateStaticUnitCommand) error
	CreatePortalUnit(CreatePortalUnitCommand) error

	HandlePortalUnitCreatedDomainEvent(portalunitmodel.PortalUnitCreated) error
	HandlePortalUnitUpdatedDomainEvent(portalunitmodel.PortalUnitUpdated) error
	HandlePortalUnitDeletedDomainEvent(portalunitmodel.PortalUnitDeleted) error

	HandleStaticUnitCreatedDomainEvent(staticunitmodel.StaticUnitCreated) error
	HandleStaticUnitUpdatedDomainEvent(staticunitmodel.StaticUnitUpdated) error
	HandleStaticUnitDeletedDomainEvent(staticunitmodel.StaticUnitDeleted) error
}

type serve struct {
	worldRepo         worldmodel.WorldRepo
	unitRepo          unitmodel.UnitRepo
	itemRepo          itemmodel.ItemRepo
	unitService       service.UnitService
	staticUnitService service.StaticUnitService
	portalUnitService service.PortalUnitService
}

func NewService(
	worldRepo worldmodel.WorldRepo,
	unitRepo unitmodel.UnitRepo,
	itemRepo itemmodel.ItemRepo,
	unitService service.UnitService,
	staticUnitService service.StaticUnitService,
	portalUnitService service.PortalUnitService,
) Service {
	return &serve{
		worldRepo:         worldRepo,
		unitRepo:          unitRepo,
		itemRepo:          itemRepo,
		unitService:       unitService,
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

func (serve *serve) RemoveUnit(command RemoveUnitCommand) error {
	worldId := globalcommonmodel.NewWorldId(command.WorldId)
	position := worldcommonmodel.NewPosition(command.Position.X, command.Position.Z)
	unitId := unitmodel.NewUnitId(worldId, position)
	unit, err := serve.unitRepo.Get(unitId)
	if err != nil {
		return err
	}

	if unit.GetType().IsEqual(worldcommonmodel.NewPortalUnitType()) {
		return serve.portalUnitService.RemovePortalUnit(unitId)
	} else if unit.GetType().IsEqual(worldcommonmodel.NewStaticUnitType()) {
		return serve.staticUnitService.RemoveStaticUnit(unitId)
	}

	return nil
}

func (serve *serve) HandlePortalUnitCreatedDomainEvent(portalUnitCreated portalunitmodel.PortalUnitCreated) error {
	portalUnit := portalUnitCreated.GetPortalUnit()

	return serve.unitService.CreateUnit(
		portalUnit.GetWorldId(),
		portalUnit.GetItemId(),
		portalUnit.GetPosition(),
		portalUnit.GetDirection(),
		worldcommonmodel.NewPortalUnitType(),
	)
}

func (serve *serve) HandlePortalUnitUpdatedDomainEvent(portalUnitUpdated portalunitmodel.PortalUnitUpdated) error {
	portalUnit := portalUnitUpdated.GetPortalUnit()

	return serve.unitService.UpdateUnit(
		portalUnit.GetId(),
		portalUnit.GetDirection(),
	)
}

func (serve *serve) HandlePortalUnitDeletedDomainEvent(portalUnitDeleted portalunitmodel.PortalUnitDeleted) error {
	portalUnit := portalUnitDeleted.GetPortalUnit()

	return serve.unitService.RemoveUnit(portalUnit.GetId())
}

func (serve *serve) HandleStaticUnitCreatedDomainEvent(staticCreated staticunitmodel.StaticUnitCreated) error {
	staticUnit := staticCreated.GetStaticUnit()

	return serve.unitService.CreateUnit(
		staticUnit.GetWorldId(),
		staticUnit.GetItemId(),
		staticUnit.GetPosition(),
		staticUnit.GetDirection(),
		worldcommonmodel.NewStaticUnitType(),
	)
}

func (serve *serve) HandleStaticUnitUpdatedDomainEvent(staticUnitUpdated staticunitmodel.StaticUnitUpdated) error {
	portalUnit := staticUnitUpdated.GetStaticUnit()

	return serve.unitService.UpdateUnit(
		portalUnit.GetId(),
		portalUnit.GetDirection(),
	)
}

func (serve *serve) HandleStaticUnitDeletedDomainEvent(staticDeleted staticunitmodel.StaticUnitDeleted) error {
	staticUnit := staticDeleted.GetStaticUnit()

	return serve.unitService.RemoveUnit(staticUnit.GetId())
}
