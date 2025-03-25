package unitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
)

type UnitEntity struct {
	id        UnitId
	worldId   globalcommonmodel.WorldId
	position  worldcommonmodel.Position
	itemId    worldcommonmodel.ItemId
	direction worldcommonmodel.Direction
	dimension worldcommonmodel.Dimension
	label     *string
	color     *globalcommonmodel.Color
	_type     worldcommonmodel.UnitType
}

// Interface Implementation Check
var _ domain.Entity[UnitId] = (*UnitEntity)(nil)

func NewUnitEntity(
	id UnitId,
	worldId globalcommonmodel.WorldId,
	position worldcommonmodel.Position,
	itemId worldcommonmodel.ItemId,
	direction worldcommonmodel.Direction,
	dimension worldcommonmodel.Dimension,
	label *string,
	color *globalcommonmodel.Color,
	_type worldcommonmodel.UnitType,
) UnitEntity {
	return UnitEntity{
		id:        id,
		worldId:   worldId,
		position:  position,
		itemId:    itemId,
		direction: direction,
		dimension: dimension,
		label:     label,
		color:     color,
		_type:     _type,
	}
}

func LoadUnitEntity(
	id UnitId,
	worldId globalcommonmodel.WorldId,
	position worldcommonmodel.Position,
	itemId worldcommonmodel.ItemId,
	direction worldcommonmodel.Direction,
	dimension worldcommonmodel.Dimension,
	label *string,
	color *globalcommonmodel.Color,
	_type worldcommonmodel.UnitType,
) UnitEntity {
	return UnitEntity{
		id:        id,
		worldId:   worldId,
		position:  position,
		itemId:    itemId,
		direction: direction,
		dimension: dimension,
		label:     label,
		color:     color,
		_type:     _type,
	}
}

func (unit *UnitEntity) GetId() UnitId {
	return unit.id
}

func (unit *UnitEntity) GetWorldId() globalcommonmodel.WorldId {
	return unit.worldId
}

func (unit *UnitEntity) GetPosition() worldcommonmodel.Position {
	return unit.position
}

func (unit *UnitEntity) GetItemId() worldcommonmodel.ItemId {
	return unit.itemId
}

func (unit *UnitEntity) GetDirection() worldcommonmodel.Direction {
	return unit.direction
}

func (unit *UnitEntity) GetDimension() worldcommonmodel.Dimension {
	return unit.dimension
}

func (unit *UnitEntity) GetLabel() *string {
	return unit.label
}

func (unit *UnitEntity) GetColor() *globalcommonmodel.Color {
	return unit.color
}

func (unit *UnitEntity) GetType() worldcommonmodel.UnitType {
	return unit._type
}

func (unit *UnitEntity) GetOccupiedBound() worldcommonmodel.Bound {
	occupiedBoundFromPos := unit.position
	var occupiedBoundToPos worldcommonmodel.Position

	dimensionWidth := int(unit.dimension.GetWidth())
	dimensionDepth := int(unit.dimension.GetDepth())

	if unit.direction.IsDown() {
		occupiedBoundToPos = occupiedBoundFromPos.Shift(dimensionWidth-1, dimensionDepth-1)
	} else if unit.direction.IsRight() {
		occupiedBoundToPos = occupiedBoundFromPos.Shift(dimensionDepth-1, dimensionWidth-1)
	} else if unit.direction.IsUp() {
		occupiedBoundToPos = occupiedBoundFromPos.Shift(dimensionWidth-1, dimensionDepth-1)
	} else {
		occupiedBoundToPos = occupiedBoundFromPos.Shift(dimensionDepth-1, dimensionWidth-1)
	}

	occupiedBound, _ := worldcommonmodel.NewBound(occupiedBoundFromPos, occupiedBoundToPos)
	return occupiedBound
}

// Get all the positions occupied by the unit
func (unit *UnitEntity) GetOccupiedPositions() []worldcommonmodel.Position {
	occupiedPositions := make([]worldcommonmodel.Position, 0)
	unit.GetOccupiedBound().Iterate(func(position worldcommonmodel.Position) {
		occupiedPositions = append(occupiedPositions, position)
	})

	return occupiedPositions
}

func (unit *UnitEntity) Rotate() {
	dimension := unit.dimension

	if dimension.IsSymmetric() {
		unit.direction = unit.direction.Rotate()
	} else {
		unit.direction = unit.direction.Rotate().Rotate()
	}
}

func (unit *UnitEntity) Move(position worldcommonmodel.Position) {
	unit.position = position
}
