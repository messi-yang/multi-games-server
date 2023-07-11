package service

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/commonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel/unitmodel"
)

var (
	errUnitExceededBoundary = fmt.Errorf("unit exceeded the boundary of the world")
)

type UnitService interface {
	CreateUnit(sharedkernelmodel.WorldId, commonmodel.ItemId, commonmodel.Position, commonmodel.Direction) error
	RemoveUnit(sharedkernelmodel.WorldId, commonmodel.Position) error
}

type unitServe struct {
	worldRepo worldmodel.WorldRepo
	unitRepo  unitmodel.UnitRepo
}

func NewUnitService(
	worldRepo worldmodel.WorldRepo,
	unitRepo unitmodel.UnitRepo,
) UnitService {
	return &unitServe{
		worldRepo: worldRepo,
		unitRepo:  unitRepo,
	}
}

func (unitServe *unitServe) CreateUnit(
	worldId sharedkernelmodel.WorldId,
	itemId commonmodel.ItemId,
	position commonmodel.Position,
	direction commonmodel.Direction,
) error {
	world, err := unitServe.worldRepo.Get(worldId)
	if err != nil {
		return err
	}

	if !world.GetBound().CoversPosition(position) {
		return errUnitExceededBoundary
	}

	if position.IsEqual(commonmodel.NewPosition(0, 0)) {
		return nil
	}

	unit, err := unitServe.unitRepo.GetUnitAt(worldId, position)
	if err != nil {
		return err
	}
	if unit != nil {
		return nil
	}

	newUnit := unitmodel.NewUnit(unitmodel.NewUnitId(worldId, position), worldId, position, itemId, direction)
	return unitServe.unitRepo.Add(newUnit)
}

func (unitServe *unitServe) RemoveUnit(worldId sharedkernelmodel.WorldId, position commonmodel.Position) error {
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
	return unitServe.unitRepo.Delete(*unit)
}
