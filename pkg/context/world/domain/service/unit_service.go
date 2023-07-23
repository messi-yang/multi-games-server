package service

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel/unitmodel"
)

var (
	errUnitExceededBoundary = fmt.Errorf("unit exceeded the boundary of the world")
)

type UnitService interface {
	CreateUnit(globalcommonmodel.WorldId, worldcommonmodel.ItemId, worldcommonmodel.Position, worldcommonmodel.Direction) error
	RemoveUnit(globalcommonmodel.WorldId, worldcommonmodel.Position) error
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
	worldId globalcommonmodel.WorldId,
	itemId worldcommonmodel.ItemId,
	position worldcommonmodel.Position,
	direction worldcommonmodel.Direction,
) error {
	world, err := unitServe.worldRepo.Get(worldId)
	if err != nil {
		return err
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

	newUnit := unitmodel.NewUnit(unitmodel.NewUnitId(worldId, position), worldId, position, itemId, direction)
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
	return unitServe.unitRepo.Delete(*unit)
}
