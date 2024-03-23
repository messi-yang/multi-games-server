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

func (unit *UnitEntity) Rotate() {
	unit.direction = unit.direction.Rotate()
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
