package unitappsrv

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/itemmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/fenceunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/linkunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/portalunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/staticunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/service"
	"github.com/samber/lo"
)

type Service interface {
	GetUnits(GetUnitsQuery) (unitDtos []dto.UnitDto, err error)

	RotateUnit(RotateUnitCommand) error

	CreateStaticUnit(CreateStaticUnitCommand) error
	RemoveStaticUnit(RemoveStaticUnitCommand) error

	CreateFenceUnit(CreateFenceUnitCommand) error
	RemoveFenceUnit(RemoveFenceUnitCommand) error

	CreatePortalUnit(CreatePortalUnitCommand) error
	RemovePortalUnit(RemovePortalUnitCommand) error
}

type serve struct {
	worldRepo         worldmodel.WorldRepo
	unitRepo          unitmodel.UnitRepo
	itemRepo          itemmodel.ItemRepo
	staticUnitService service.StaticUnitService
	fenceUnitService  service.FenceUnitService
	portalUnitService service.PortalUnitService
	linkUnitService   service.LinkUnitService
}

func NewService(
	worldRepo worldmodel.WorldRepo,
	unitRepo unitmodel.UnitRepo,
	itemRepo itemmodel.ItemRepo,
	staticUnitService service.StaticUnitService,
	fenceUnitService service.FenceUnitService,
	portalUnitService service.PortalUnitService,
	linkUnitService service.LinkUnitService,
) Service {
	return &serve{
		worldRepo:         worldRepo,
		unitRepo:          unitRepo,
		itemRepo:          itemRepo,
		staticUnitService: staticUnitService,
		fenceUnitService:  fenceUnitService,
		portalUnitService: portalUnitService,
		linkUnitService:   linkUnitService,
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

func (serve *serve) CreateStaticUnit(command CreateStaticUnitCommand) error {
	return serve.staticUnitService.CreateStaticUnit(
		staticunitmodel.NewStaticUnitId(command.Id),
		globalcommonmodel.NewWorldId(command.WorldId),
		worldcommonmodel.NewItemId(command.ItemId),
		worldcommonmodel.NewPosition(command.Position.X, command.Position.Z),
		worldcommonmodel.NewDirection(command.Direction),
	)
}

func (serve *serve) RemoveStaticUnit(command RemoveStaticUnitCommand) error {
	return serve.staticUnitService.RemoveStaticUnit(staticunitmodel.NewStaticUnitId(command.Id))
}

func (serve *serve) CreateFenceUnit(command CreateFenceUnitCommand) error {
	return serve.fenceUnitService.CreateFenceUnit(
		fenceunitmodel.NewFenceUnitId(command.Id),
		globalcommonmodel.NewWorldId(command.WorldId),
		worldcommonmodel.NewItemId(command.ItemId),
		worldcommonmodel.NewPosition(command.Position.X, command.Position.Z),
		worldcommonmodel.NewDirection(command.Direction),
	)
}

func (serve *serve) RemoveFenceUnit(command RemoveFenceUnitCommand) error {
	return serve.fenceUnitService.RemoveFenceUnit(fenceunitmodel.NewFenceUnitId(command.Id))
}

func (serve *serve) CreatePortalUnit(command CreatePortalUnitCommand) error {
	return serve.portalUnitService.CreatePortalUnit(
		portalunitmodel.NewPortalUnitId(command.Id),
		globalcommonmodel.NewWorldId(command.WorldId),
		worldcommonmodel.NewItemId(command.ItemId),
		worldcommonmodel.NewPosition(command.Position.X, command.Position.Z),
		worldcommonmodel.NewDirection(command.Direction),
	)
}

func (serve *serve) RemovePortalUnit(command RemovePortalUnitCommand) error {
	return serve.portalUnitService.RemovePortalUnit(portalunitmodel.NewPortalUnitId(command.Id))
}

func (serve *serve) RotateUnit(command RotateUnitCommand) error {
	unitId := unitmodel.NewUnitId(command.Id)
	unit, err := serve.unitRepo.Get(unitId)
	if err != nil {
		return err
	}

	if unit.GetType().IsEqual(worldcommonmodel.NewPortalUnitType()) {
		return serve.portalUnitService.RotatePortalUnit(portalunitmodel.NewPortalUnitId(command.Id))
	} else if unit.GetType().IsEqual(worldcommonmodel.NewStaticUnitType()) {
		return serve.staticUnitService.RotateStaticUnit(staticunitmodel.NewStaticUnitId(command.Id))
	} else if unit.GetType().IsEqual(worldcommonmodel.NewFenceUnitType()) {
		return serve.fenceUnitService.RotateFenceUnit(fenceunitmodel.NewFenceUnitId(command.Id))
	} else if unit.GetType().IsEqual(worldcommonmodel.NewLinkUnitType()) {
		return serve.linkUnitService.RotateLinkUnit(linkunitmodel.NewLinkUnitId(command.Id))
	}

	return nil
}
