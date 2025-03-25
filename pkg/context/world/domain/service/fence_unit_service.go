package service

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/itemmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/fenceunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
)

var (
	errItemIsNotForFenceUnit = fmt.Errorf("item is not for fence unit")
)

type FenceUnitService interface {
	CreateFenceUnit(
		fenceunitmodel.FenceUnitId,
		globalcommonmodel.WorldId,
		worldcommonmodel.ItemId,
		worldcommonmodel.Position,
		worldcommonmodel.Direction,
	) error
	RotateFenceUnit(fenceunitmodel.FenceUnitId) error
	RemoveFenceUnit(fenceunitmodel.FenceUnitId) error
}

type fenceUnitServe struct {
	unitRepo      unitmodel.UnitRepo
	fenceUnitRepo fenceunitmodel.FenceUnitRepo
	itemRepo      itemmodel.ItemRepo
}

func NewFenceUnitService(
	unitRepo unitmodel.UnitRepo,
	fenceUnitRepo fenceunitmodel.FenceUnitRepo,
	itemRepo itemmodel.ItemRepo,
) FenceUnitService {
	return &fenceUnitServe{
		unitRepo:      unitRepo,
		fenceUnitRepo: fenceUnitRepo,
		itemRepo:      itemRepo,
	}
}

func (fenceUnitServe *fenceUnitServe) CreateFenceUnit(
	id fenceunitmodel.FenceUnitId,
	worldId globalcommonmodel.WorldId,
	itemId worldcommonmodel.ItemId,
	position worldcommonmodel.Position,
	direction worldcommonmodel.Direction,
) error {
	item, err := fenceUnitServe.itemRepo.Get(itemId)
	if err != nil {
		return err
	}

	if !item.GetCompatibleUnitType().IsFence() {
		return errItemIsNotForFenceUnit
	}

	if position.IsEqual(worldcommonmodel.NewPosition(0, 0)) {
		return errUnitCannotBeAtOriginPosition
	}

	newFenceUnit := fenceunitmodel.NewFenceUnit(id, worldId, position, itemId, direction, item.GetDimension())

	hasUnitsInBound, err := fenceUnitServe.unitRepo.HasUnitsInBound(worldId, newFenceUnit.GetOccupiedBound())
	if err != nil {
		return err
	}
	if hasUnitsInBound {
		return errBoundAlreadyHasUnit
	}

	return fenceUnitServe.fenceUnitRepo.Add(newFenceUnit)
}

func (fenceUnitServe *fenceUnitServe) RotateFenceUnit(id fenceunitmodel.FenceUnitId) error {
	unit, err := fenceUnitServe.fenceUnitRepo.Get(id)
	if err != nil {
		return err
	}
	unit.Rotate()

	return fenceUnitServe.fenceUnitRepo.Update(unit)
}

func (fenceUnitServe *fenceUnitServe) RemoveFenceUnit(id fenceunitmodel.FenceUnitId) error {
	unit, err := fenceUnitServe.fenceUnitRepo.Get(id)
	if err != nil {
		return err
	}

	unit.Delete()
	return fenceUnitServe.fenceUnitRepo.Delete(unit)
}
