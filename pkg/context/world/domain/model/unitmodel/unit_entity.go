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
	_type     worldcommonmodel.UnitType
	info      any
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
	_type worldcommonmodel.UnitType,
	info any,
) UnitEntity {
	return UnitEntity{
		id:        id,
		worldId:   worldId,
		position:  position,
		itemId:    itemId,
		direction: direction,
		dimension: dimension,
		label:     label,
		_type:     _type,
		info:      info,
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
	_type worldcommonmodel.UnitType,
	info any,
) UnitEntity {
	return UnitEntity{
		id:        id,
		worldId:   worldId,
		position:  position,
		itemId:    itemId,
		direction: direction,
		dimension: dimension,
		label:     label,
		_type:     _type,
		info:      info,
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

func (unit *UnitEntity) GetInfo() any {
	return unit.info
}

func (unit *UnitEntity) GetType() worldcommonmodel.UnitType {
	return unit._type
}

func (unit *UnitEntity) getOccupiedBound() worldcommonmodel.Bound {
	var occupiedBoundFromPos worldcommonmodel.Position
	var occupiedBoundToPos worldcommonmodel.Position

	dimensionWidth := int(unit.dimension.GetWidth())
	dimensionDepth := int(unit.dimension.GetDepth())

	if unit.direction.IsDown() {
		occupiedBoundFromPos = unit.position
		occupiedBoundToPos = occupiedBoundFromPos.Shift(dimensionWidth-1, dimensionDepth-1)
	} else if unit.direction.IsRight() {
		occupiedBoundFromPos = unit.position.Shift(0, -dimensionWidth+1)
		occupiedBoundToPos = occupiedBoundFromPos.Shift(dimensionDepth-1, dimensionWidth-1)
	} else if unit.direction.IsUp() {
		occupiedBoundFromPos = unit.position.Shift(-dimensionWidth+1, -dimensionDepth+1)
		occupiedBoundToPos = occupiedBoundFromPos.Shift(dimensionWidth-1, dimensionDepth-1)
	} else {
		occupiedBoundFromPos = unit.position.Shift(-dimensionDepth+1, 0)
		occupiedBoundToPos = occupiedBoundFromPos.Shift(dimensionDepth-1, dimensionWidth-1)
	}

	occupiedBound, _ := worldcommonmodel.NewBound(occupiedBoundFromPos, occupiedBoundToPos)
	return occupiedBound
}

// Get all the positions occupied by the unit
func (unit *UnitEntity) GetOccupiedPositions() []worldcommonmodel.Position {
	occupiedPositions := make([]worldcommonmodel.Position, 0)
	unit.getOccupiedBound().Iterate(func(position worldcommonmodel.Position) {
		occupiedPositions = append(occupiedPositions, position)
	})

	return occupiedPositions
}

// Rotate the unit without chaning its occupied bound.
// When you rotate a non-symmetric unit, it does a flip to make sure the occupied bound doesn't change.
func (unit *UnitEntity) Rotate() {
	dimension := unit.dimension
	direction := unit.direction

	occupiedBound := unit.getOccupiedBound()
	if dimension.IsSymmetric() {
		if direction.IsDown() {
			unit.direction = worldcommonmodel.NewRightDirection()
			unit.position = occupiedBound.GetLeftDown()
		} else if direction.IsRight() {
			unit.direction = worldcommonmodel.NewUpDirection()
			unit.position = occupiedBound.GetTo()
		} else if direction.IsUp() {
			unit.direction = worldcommonmodel.NewLeftDirection()
			unit.position = occupiedBound.GetRightUp()
		} else {
			unit.direction = worldcommonmodel.NewDownDirection()
			unit.position = occupiedBound.GetFrom()
		}
	} else {
		if direction.IsDown() {
			unit.direction = worldcommonmodel.NewUpDirection()
			unit.position = occupiedBound.GetTo()
		} else if direction.IsRight() {
			unit.direction = worldcommonmodel.NewLeftDirection()
			unit.position = occupiedBound.GetRightUp()
		} else if direction.IsUp() {
			unit.direction = worldcommonmodel.NewDownDirection()
			unit.position = occupiedBound.GetFrom()
		} else {
			unit.direction = worldcommonmodel.NewRightDirection()
			unit.position = occupiedBound.GetLeftDown()
		}
	}
}
