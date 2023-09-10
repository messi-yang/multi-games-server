package service

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/itemmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/staticunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
)

var (
	errItemIsNotForStaticUnit = fmt.Errorf("item is not for static unit")
)

type StaticUnitService interface {
	CreateStaticUnit(
		globalcommonmodel.WorldId,
		worldcommonmodel.ItemId,
		worldcommonmodel.Position,
		worldcommonmodel.Direction,
	) error
	RotateStaticUnit(unitmodel.UnitId) error
	RemoveStaticUnit(unitmodel.UnitId) error
}

type staticUnitServe struct {
	worldRepo      worldmodel.WorldRepo
	unitRepo       unitmodel.UnitRepo
	staticUnitRepo staticunitmodel.StaticUnitRepo
	itemRepo       itemmodel.ItemRepo
}

func NewStaticUnitService(
	worldRepo worldmodel.WorldRepo,
	unitRepo unitmodel.UnitRepo,
	staticUnitRepo staticunitmodel.StaticUnitRepo,
	itemRepo itemmodel.ItemRepo,
) StaticUnitService {
	return &staticUnitServe{
		worldRepo:      worldRepo,
		unitRepo:       unitRepo,
		staticUnitRepo: staticUnitRepo,
		itemRepo:       itemRepo,
	}
}

func (staticUnitServe *staticUnitServe) CreateStaticUnit(
	worldId globalcommonmodel.WorldId,
	itemId worldcommonmodel.ItemId,
	position worldcommonmodel.Position,
	direction worldcommonmodel.Direction,
) error {
	world, err := staticUnitServe.worldRepo.Get(worldId)
	if err != nil {
		return err
	}

	item, err := staticUnitServe.itemRepo.Get(itemId)
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

	unit, err := staticUnitServe.unitRepo.Find(unitmodel.NewUnitId(worldId, position))
	if err != nil {
		return err
	}
	if unit != nil {
		return nil
	}

	newStaticUnit := staticunitmodel.NewStaticUnit(worldId, position, itemId, direction)
	return staticUnitServe.staticUnitRepo.Add(newStaticUnit)
}

func (staticUnitServe *staticUnitServe) RotateStaticUnit(unitId unitmodel.UnitId) error {
	unit, err := staticUnitServe.staticUnitRepo.Get(unitId)
	if err != nil {
		return err
	}
	unit.Rotate()

	return staticUnitServe.staticUnitRepo.Update(unit)
}

func (staticUnitServe *staticUnitServe) RemoveStaticUnit(unitId unitmodel.UnitId) error {
	unit, err := staticUnitServe.staticUnitRepo.Get(unitId)
	if err != nil {
		return err
	}

	unit.Delete()
	return staticUnitServe.staticUnitRepo.Delete(unit)
}
