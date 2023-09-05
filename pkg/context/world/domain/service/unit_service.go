package service

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/itemmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
	"github.com/samber/lo"
)

var (
	errUnitExceededBoundary   = fmt.Errorf("unit exceeded the boundary of the world")
	errItemIsNotForStaticUnit = fmt.Errorf("item is not for static unit")
	errItemIsNotForPortalUnit = fmt.Errorf("item is not for portal unit")
	errUnitIsNotStatic        = fmt.Errorf("unit is not static")
	// errTargetPositionHasNoPortalUnit = fmt.Errorf("target position has no portal unit")
)

type UnitService interface {
	CreateStaticUnit(globalcommonmodel.WorldId, worldcommonmodel.ItemId, worldcommonmodel.Position, worldcommonmodel.Direction) error
	CreatePortalUnit(
		worldId globalcommonmodel.WorldId,
		itemId worldcommonmodel.ItemId,
		position worldcommonmodel.Position,
		direction worldcommonmodel.Direction,
	) error
	RemovePortalUnit(unitmodel.UnitId) error
	RemoveStaticUnit(unitmodel.UnitId) error
}

type unitServe struct {
	worldRepo      worldmodel.WorldRepo
	unitRepo       unitmodel.UnitRepo
	portalUnitRepo unitmodel.PortalUnitRepo
	itemRepo       itemmodel.ItemRepo
}

func NewUnitService(
	worldRepo worldmodel.WorldRepo,
	unitRepo unitmodel.UnitRepo,
	portalUnitRepo unitmodel.PortalUnitRepo,
	itemRepo itemmodel.ItemRepo,
) UnitService {
	return &unitServe{
		worldRepo:      worldRepo,
		unitRepo:       unitRepo,
		portalUnitRepo: portalUnitRepo,
		itemRepo:       itemRepo,
	}
}

func (unitServe *unitServe) CreateStaticUnit(
	worldId globalcommonmodel.WorldId,
	itemId worldcommonmodel.ItemId,
	position worldcommonmodel.Position,
	direction worldcommonmodel.Direction,
) error {
	world, err := unitServe.worldRepo.Get(worldId)
	if err != nil {
		return err
	}

	item, err := unitServe.itemRepo.Get(itemId)
	if err != nil {
		return err
	}

	if !item.GetCompatibleUnitType().IsStatic() {
		return errItemIsNotForStaticUnit
	}

	if !world.GetBound().CoversPosition(position) {
		return errUnitExceededBoundary
	}

	if position.IsEqual(worldcommonmodel.NewPosition(0, 0)) {
		return nil
	}

	unit, err := unitServe.unitRepo.Find(unitmodel.NewUnitId(worldId, position))
	if err != nil {
		return err
	}
	if unit != nil {
		return nil
	}

	newUnit := unitmodel.NewUnit(worldId, position, itemId, direction, item.GetCompatibleUnitType())
	return unitServe.unitRepo.Add(newUnit)
}

func (unitServe *unitServe) CreatePortalUnit(
	worldId globalcommonmodel.WorldId,
	itemId worldcommonmodel.ItemId,
	position worldcommonmodel.Position,
	direction worldcommonmodel.Direction,
) error {
	world, err := unitServe.worldRepo.Get(worldId)
	if err != nil {
		return err
	}

	item, err := unitServe.itemRepo.Get(itemId)
	if err != nil {
		return err
	}

	if !item.GetCompatibleUnitType().IsPortal() {
		return errItemIsNotForPortalUnit
	}

	if !world.GetBound().CoversPosition(position) {
		return errUnitExceededBoundary
	}

	if position.IsEqual(worldcommonmodel.NewPosition(0, 0)) {
		return nil
	}

	existingUnit, err := unitServe.unitRepo.Find(unitmodel.NewUnitId(worldId, position))
	if err != nil {
		return err
	}
	if existingUnit != nil {
		return nil
	}

	randomPortalUnit, err := unitServe.portalUnitRepo.GetRandomPortalUnit(worldId)
	if err != nil {
		return err
	}

	newPortalUnit := unitmodel.NewPortalUnit(
		worldId,
		position,
		itemId,
		direction,
		lo.TernaryF[*worldcommonmodel.Position](
			randomPortalUnit == nil,
			func() *worldcommonmodel.Position { return nil },
			func() *worldcommonmodel.Position { return commonutil.ToPointer(randomPortalUnit.GetPosition()) },
		),
	)

	return unitServe.portalUnitRepo.Add(newPortalUnit)
}

func (unitServe *unitServe) RemoveStaticUnit(unitId unitmodel.UnitId) error {
	unit, err := unitServe.unitRepo.Find(unitId)
	if err != nil {
		return err
	}
	if unit == nil {
		return nil
	}

	if !unit.GetType().IsEqual(worldcommonmodel.NewStaticUnitType()) {
		return errUnitIsNotStatic
	}

	return unitServe.unitRepo.Delete(*unit)
}

func (unitServe *unitServe) RemovePortalUnit(unitId unitmodel.UnitId) error {
	portalUnit, err := unitServe.portalUnitRepo.Get(unitId)
	if err != nil {
		return err
	}

	portalUnit.Delete()
	return unitServe.portalUnitRepo.Delete(portalUnit)
}
