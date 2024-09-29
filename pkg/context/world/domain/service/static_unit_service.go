package service

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/itemmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/staticunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
)

var (
	errItemIsNotForStaticUnit = fmt.Errorf("item is not for static unit")
)

type StaticUnitService interface {
	CreateStaticUnit(
		staticunitmodel.StaticUnitId,
		globalcommonmodel.WorldId,
		worldcommonmodel.ItemId,
		worldcommonmodel.Position,
		worldcommonmodel.Direction,
	) error
	RotateStaticUnit(staticunitmodel.StaticUnitId) error
	RemoveStaticUnit(staticunitmodel.StaticUnitId) error
}

type staticUnitServe struct {
	unitRepo       unitmodel.UnitRepo
	staticUnitRepo staticunitmodel.StaticUnitRepo
	itemRepo       itemmodel.ItemRepo
}

func NewStaticUnitService(
	unitRepo unitmodel.UnitRepo,
	staticUnitRepo staticunitmodel.StaticUnitRepo,
	itemRepo itemmodel.ItemRepo,
) StaticUnitService {
	return &staticUnitServe{
		unitRepo:       unitRepo,
		staticUnitRepo: staticUnitRepo,
		itemRepo:       itemRepo,
	}
}

func (staticUnitServe *staticUnitServe) CreateStaticUnit(
	id staticunitmodel.StaticUnitId,
	worldId globalcommonmodel.WorldId,
	itemId worldcommonmodel.ItemId,
	position worldcommonmodel.Position,
	direction worldcommonmodel.Direction,
) error {
	item, err := staticUnitServe.itemRepo.Get(itemId)
	if err != nil {
		return err
	}

	if !item.GetCompatibleUnitType().IsStatic() {
		return errItemIsNotForStaticUnit
	}

	if position.IsEqual(worldcommonmodel.NewPosition(0, 0)) {
		return errUnitCannotBeAtOriginPosition
	}

	unit, err := staticUnitServe.unitRepo.Find(worldId, position)
	if err != nil {
		return err
	}
	if unit != nil {
		return errPositionAlreadyHasUnit
	}

	newStaticUnit := staticunitmodel.NewStaticUnit(id, worldId, position, itemId, direction, item.GetDimension())
	return staticUnitServe.staticUnitRepo.Add(newStaticUnit)
}

func (staticUnitServe *staticUnitServe) RotateStaticUnit(id staticunitmodel.StaticUnitId) error {
	unit, err := staticUnitServe.staticUnitRepo.Get(id)
	if err != nil {
		return err
	}
	unit.Rotate()

	return staticUnitServe.staticUnitRepo.Update(unit)
}

func (staticUnitServe *staticUnitServe) RemoveStaticUnit(id staticunitmodel.StaticUnitId) error {
	unit, err := staticUnitServe.staticUnitRepo.Get(id)
	if err != nil {
		return err
	}

	unit.Delete()
	return staticUnitServe.staticUnitRepo.Delete(unit)
}
