package service

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/itemmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/fenceunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
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
	RotateFenceUnit(unitmodel.UnitId) error
	RemoveFenceUnit(unitmodel.UnitId) error
}

type fenceUnitServe struct {
	worldRepo     worldmodel.WorldRepo
	unitRepo      unitmodel.UnitRepo
	fenceUnitRepo fenceunitmodel.FenceUnitRepo
	itemRepo      itemmodel.ItemRepo
}

func NewFenceUnitService(
	worldRepo worldmodel.WorldRepo,
	unitRepo unitmodel.UnitRepo,
	fenceUnitRepo fenceunitmodel.FenceUnitRepo,
	itemRepo itemmodel.ItemRepo,
) FenceUnitService {
	return &fenceUnitServe{
		worldRepo:     worldRepo,
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
	world, err := fenceUnitServe.worldRepo.Get(worldId)
	if err != nil {
		return err
	}

	item, err := fenceUnitServe.itemRepo.Get(itemId)
	if err != nil {
		return err
	}

	if !item.GetCompatibleUnitType().IsFence() {
		return errItemIsNotForFenceUnit
	}

	if !world.GetBound().CoversPosition(position) {
		return errUnitExceededBoundary
	}

	if position.IsEqual(worldcommonmodel.NewPosition(0, 0)) {
		return nil
	}

	unit, err := fenceUnitServe.unitRepo.Find(worldId, position)
	if err != nil {
		return err
	}
	if unit != nil {
		return nil
	}

	newFenceUnit := fenceunitmodel.NewFenceUnit(id, worldId, position, itemId, direction)
	return fenceUnitServe.fenceUnitRepo.Add(newFenceUnit)
}

func (fenceUnitServe *fenceUnitServe) RotateFenceUnit(unitId unitmodel.UnitId) error {
	unit, err := fenceUnitServe.fenceUnitRepo.Get(unitId)
	if err != nil {
		return err
	}
	unit.Rotate()

	return fenceUnitServe.fenceUnitRepo.Update(unit)
}

func (fenceUnitServe *fenceUnitServe) RemoveFenceUnit(unitId unitmodel.UnitId) error {
	unit, err := fenceUnitServe.fenceUnitRepo.Get(unitId)
	if err != nil {
		return err
	}

	unit.Delete()
	return fenceUnitServe.fenceUnitRepo.Delete(unit)
}
