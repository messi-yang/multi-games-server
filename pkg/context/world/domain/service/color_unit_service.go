package service

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/itemmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/colorunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
)

var (
	errItemIsNotForColorUnit = fmt.Errorf("item is not for color unit")
)

type ColorUnitService interface {
	CreateColorUnit(
		colorunitmodel.ColorUnitId,
		globalcommonmodel.WorldId,
		worldcommonmodel.ItemId,
		worldcommonmodel.Position,
		worldcommonmodel.Direction,
		*string,
		globalcommonmodel.Color,
	) error
	RotateColorUnit(colorunitmodel.ColorUnitId) error
	RemoveColorUnit(colorunitmodel.ColorUnitId) error
}

type colorUnitServe struct {
	unitRepo      unitmodel.UnitRepo
	colorUnitRepo colorunitmodel.ColorUnitRepo
	itemRepo      itemmodel.ItemRepo
}

func NewColorUnitService(
	unitRepo unitmodel.UnitRepo,
	colorUnitRepo colorunitmodel.ColorUnitRepo,
	itemRepo itemmodel.ItemRepo,
) ColorUnitService {
	return &colorUnitServe{
		unitRepo:      unitRepo,
		colorUnitRepo: colorUnitRepo,
		itemRepo:      itemRepo,
	}
}

func (colorUnitServe *colorUnitServe) CreateColorUnit(
	id colorunitmodel.ColorUnitId,
	worldId globalcommonmodel.WorldId,
	itemId worldcommonmodel.ItemId,
	position worldcommonmodel.Position,
	direction worldcommonmodel.Direction,
	label *string,
	color globalcommonmodel.Color,
) error {
	item, err := colorUnitServe.itemRepo.Get(itemId)
	if err != nil {
		return err
	}

	if !item.GetCompatibleUnitType().IsColor() {
		return errItemIsNotForColorUnit
	}

	if position.IsEqual(worldcommonmodel.NewPosition(0, 0)) {
		return nil
	}

	newColorUnit := colorunitmodel.NewColorUnit(id, worldId, position, itemId, direction, item.GetDimension(), label, &color)

	hasUnitsInBound, err := colorUnitServe.unitRepo.HasUnitsInBound(worldId, newColorUnit.GetOccupiedBound())
	if err != nil {
		return err
	}
	if hasUnitsInBound {
		return errBoundAlreadyHasUnit
	}

	return colorUnitServe.colorUnitRepo.Add(newColorUnit)
}

func (colorUnitServe *colorUnitServe) RotateColorUnit(id colorunitmodel.ColorUnitId) error {
	unit, err := colorUnitServe.colorUnitRepo.Get(id)
	if err != nil {
		return err
	}
	unit.Rotate()

	return colorUnitServe.colorUnitRepo.Update(unit)
}

func (colorUnitServe *colorUnitServe) RemoveColorUnit(id colorunitmodel.ColorUnitId) error {
	unit, err := colorUnitServe.colorUnitRepo.Get(id)
	if err != nil {
		return err
	}

	unit.Delete()
	return colorUnitServe.colorUnitRepo.Delete(unit)
}
