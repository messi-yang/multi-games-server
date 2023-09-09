package service

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/itemmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/portalunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/staticunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
)

var (
	errUnitExceededBoundary   = fmt.Errorf("unit exceeded the boundary of the world")
	errItemIsNotForStaticUnit = fmt.Errorf("item is not for static unit")
	errItemIsNotForPortalUnit = fmt.Errorf("item is not for portal unit")
	errUnitIsNotStatic        = fmt.Errorf("unit is not static")
	// errTargetPositionHasNoPortalUnit = fmt.Errorf("target position has no portal unit")
)

type UnitService interface {
	CreateUnit(
		globalcommonmodel.WorldId,
		worldcommonmodel.ItemId,
		worldcommonmodel.Position,
		worldcommonmodel.Direction,
		worldcommonmodel.UnitType,
	) error
	UpdateUnit(
		unitmodel.UnitId,
		worldcommonmodel.Direction,
	) error
	RemoveUnit(unitmodel.UnitId) error
	CreateStaticUnit(
		globalcommonmodel.WorldId,
		worldcommonmodel.ItemId,
		worldcommonmodel.Position,
		worldcommonmodel.Direction,
	) error
	RotateStaticUnit(unitmodel.UnitId) error
	RemoveStaticUnit(unitmodel.UnitId) error
	CreatePortalUnit(
		worldId globalcommonmodel.WorldId,
		itemId worldcommonmodel.ItemId,
		position worldcommonmodel.Position,
		direction worldcommonmodel.Direction,
	) error
	RotatePortalUnit(unitmodel.UnitId) error
	RemovePortalUnit(unitmodel.UnitId) error
}

type unitServe struct {
	worldRepo      worldmodel.WorldRepo
	unitRepo       unitmodel.UnitRepo
	staticUnitRepo staticunitmodel.StaticUnitRepo
	portalUnitRepo portalunitmodel.PortalUnitRepo
	itemRepo       itemmodel.ItemRepo
}

func NewUnitService(
	worldRepo worldmodel.WorldRepo,
	unitRepo unitmodel.UnitRepo,
	staticUnitRepo staticunitmodel.StaticUnitRepo,
	portalUnitRepo portalunitmodel.PortalUnitRepo,
	itemRepo itemmodel.ItemRepo,
) UnitService {
	return &unitServe{
		worldRepo:      worldRepo,
		unitRepo:       unitRepo,
		staticUnitRepo: staticUnitRepo,
		portalUnitRepo: portalUnitRepo,
		itemRepo:       itemRepo,
	}
}

func (unitServe *unitServe) CreateUnit(
	worldId globalcommonmodel.WorldId,
	itemId worldcommonmodel.ItemId,
	position worldcommonmodel.Position,
	direction worldcommonmodel.Direction,
	_type worldcommonmodel.UnitType,
) error {
	return unitServe.unitRepo.Add(unitmodel.NewUnit(worldId, position, itemId, direction, _type))
}

func (unitServe *unitServe) UpdateUnit(unitId unitmodel.UnitId, direction worldcommonmodel.Direction) error {
	unit, err := unitServe.unitRepo.Get(unitId)
	if err != nil {
		return err
	}
	unit.UpdateDirection(direction)
	return unitServe.unitRepo.Update(unit)
}

func (unitServe *unitServe) RemoveUnit(unitId unitmodel.UnitId) error {
	unit, err := unitServe.unitRepo.Get(unitId)
	if err != nil {
		return err
	}
	return unitServe.unitRepo.Delete(unit)
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

	newStaticUnit := staticunitmodel.NewStaticUnit(worldId, position, itemId, direction)
	return unitServe.staticUnitRepo.Add(newStaticUnit)
}

func (unitServe *unitServe) RotateStaticUnit(unitId unitmodel.UnitId) error {
	unit, err := unitServe.staticUnitRepo.Get(unitId)
	if err != nil {
		return err
	}
	unit.Rotate()

	return unitServe.staticUnitRepo.Update(unit)
}

func (unitServe *unitServe) RemoveStaticUnit(unitId unitmodel.UnitId) error {
	unit, err := unitServe.staticUnitRepo.Get(unitId)
	if err != nil {
		return err
	}

	unit.Delete()
	return unitServe.staticUnitRepo.Delete(unit)
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

	portalUnitWithNoTarget, err := unitServe.portalUnitRepo.GetFirstPortalUnitWithNoTarget(worldId)
	if err != nil {
		return err
	}

	newPortalUnit := portalunitmodel.NewPortalUnit(
		worldId,
		position,
		itemId,
		direction,
		nil,
	)

	if portalUnitWithNoTarget != nil {
		newPortalUnit.UpdateTargetPosition(commonutil.ToPointer(portalUnitWithNoTarget.GetPosition()))
		portalUnitWithNoTarget.UpdateTargetPosition(&position)
		if err = unitServe.portalUnitRepo.Update(*portalUnitWithNoTarget); err != nil {
			return err
		}
	}

	return unitServe.portalUnitRepo.Add(newPortalUnit)
}

func (unitServe *unitServe) RotatePortalUnit(unitId unitmodel.UnitId) error {
	unit, err := unitServe.portalUnitRepo.Get(unitId)
	if err != nil {
		return err
	}
	unit.Rotate()

	return unitServe.portalUnitRepo.Update(unit)
}

func (unitServe *unitServe) RemovePortalUnit(unitId unitmodel.UnitId) error {
	unit, err := unitServe.portalUnitRepo.Get(unitId)
	if err != nil {
		return err
	}

	targetPosition := unit.GetTargetPosition()
	if targetPosition != nil {
		unitAtTargetPosition, err := unitServe.portalUnitRepo.Get(unitmodel.NewUnitId(
			unit.GetWorldId(),
			*targetPosition,
		))
		if err != nil {
			return err
		}
		unitAtTargetPosition.UpdateTargetPosition(nil)
		if err = unitServe.portalUnitRepo.Update(unitAtTargetPosition); err != nil {
			return err
		}
	}

	unit.Delete()
	return unitServe.portalUnitRepo.Delete(unit)
}
