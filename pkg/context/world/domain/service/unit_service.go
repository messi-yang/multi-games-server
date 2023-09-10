package service

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/itemmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
)

var (
	errUnitExceededBoundary = fmt.Errorf("unit exceeded the boundary of the world")
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
}

type unitServe struct {
	worldRepo worldmodel.WorldRepo
	unitRepo  unitmodel.UnitRepo
	itemRepo  itemmodel.ItemRepo
}

func NewUnitService(
	worldRepo worldmodel.WorldRepo,
	unitRepo unitmodel.UnitRepo,
	itemRepo itemmodel.ItemRepo,
) UnitService {
	return &unitServe{
		worldRepo: worldRepo,
		unitRepo:  unitRepo,
		itemRepo:  itemRepo,
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
