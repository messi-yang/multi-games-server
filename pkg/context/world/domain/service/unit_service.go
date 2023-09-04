package service

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/itemmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
)

var (
	errUnitExceededBoundary          = fmt.Errorf("unit exceeded the boundary of the world")
	errItemIsNotForStaticUnit        = fmt.Errorf("item is not for static unit")
	errItemIsNotForPortalUnit        = fmt.Errorf("item is not for portal unit")
	errTargetPositionHasNoPortalUnit = fmt.Errorf("target position has no portal unit")
)

type UnitService interface {
	CreateStaticUnit(globalcommonmodel.WorldId, worldcommonmodel.ItemId, worldcommonmodel.Position, worldcommonmodel.Direction) error
	CreatePortalUnit(
		worldId globalcommonmodel.WorldId,
		itemId worldcommonmodel.ItemId,
		position worldcommonmodel.Position,
		direction worldcommonmodel.Direction,
	) error
	RemoveUnit(globalcommonmodel.WorldId, worldcommonmodel.Position) error
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

	unit, err := unitServe.unitRepo.GetUnitAt(worldId, position)
	if err != nil {
		return err
	}
	if unit != nil {
		return nil
	}

	newUnit := unitmodel.NewUnit(worldId, position, itemId, direction, item.GetCompatibleUnitType(), nil)
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

	unit, err := unitServe.unitRepo.GetUnitAt(worldId, position)
	if err != nil {
		return err
	}
	if unit != nil {
		return nil
	}

	portalUnit := unitmodel.NewPortalUnit(nil)

	randomPortalUnit, err := unitServe.unitRepo.GetRandomPortalUnit(worldId)
	if err != nil {
		return err
	}

	if randomPortalUnit != nil {
		portalUnit.UpdateTargetPosition(commonutil.ToPointer(randomPortalUnit.GetPosition()))
	}
	if err = unitServe.portalUnitRepo.Add(portalUnit); err != nil {
		return err
	}

	newUnit := unitmodel.NewUnit(
		worldId,
		position,
		itemId,
		direction,
		item.GetCompatibleUnitType(),
		commonutil.ToPointer(portalUnit.GetId().Uuid()),
	)
	return unitServe.unitRepo.Add(newUnit)
}

func (unitServe *unitServe) RemoveUnit(worldId globalcommonmodel.WorldId, position worldcommonmodel.Position) error {
	if _, err := unitServe.worldRepo.Get(worldId); err != nil {
		return err
	}

	unit, err := unitServe.unitRepo.GetUnitAt(worldId, position)
	if err != nil {
		return err
	}
	if unit == nil {
		return nil
	}

	if unit.GetType().IsPortal() {
		portalUnit, err := unitServe.portalUnitRepo.Get(unitmodel.NewPortalUnitId(*unit.GetLinkedUnitId()))
		if err != nil {
			return err
		}
		if err = unitServe.portalUnitRepo.Delete(portalUnit); err != nil {
			return err
		}
	}

	return unitServe.unitRepo.Delete(*unit)
}
