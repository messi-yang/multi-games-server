package fenceunitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
)

type FenceUnit struct {
	id        FenceUnitId
	worldId   globalcommonmodel.WorldId
	position  worldcommonmodel.Position
	itemId    worldcommonmodel.ItemId
	direction worldcommonmodel.Direction
}

// Interface Implementation Check
var _ domain.Aggregate[FenceUnitId] = (*FenceUnit)(nil)

func NewFenceUnit(
	id FenceUnitId,
	worldId globalcommonmodel.WorldId,
	position worldcommonmodel.Position,
	itemId worldcommonmodel.ItemId,
	direction worldcommonmodel.Direction,
) FenceUnit {
	return FenceUnit{
		id:        id,
		worldId:   worldId,
		position:  position,
		itemId:    itemId,
		direction: direction,
	}
}

func LoadFenceUnit(
	id FenceUnitId,
	worldId globalcommonmodel.WorldId,
	position worldcommonmodel.Position,
	itemId worldcommonmodel.ItemId,
	direction worldcommonmodel.Direction,
) FenceUnit {
	return FenceUnit{
		id:        id,
		worldId:   worldId,
		position:  position,
		itemId:    itemId,
		direction: direction,
	}
}

func (unit *FenceUnit) GetId() FenceUnitId {
	return unit.id
}

func (unit *FenceUnit) GetWorldId() globalcommonmodel.WorldId {
	return unit.worldId
}

func (unit *FenceUnit) GetPosition() worldcommonmodel.Position {
	return unit.position
}

func (unit *FenceUnit) GetItemId() worldcommonmodel.ItemId {
	return unit.itemId
}

func (unit *FenceUnit) GetDirection() worldcommonmodel.Direction {
	return unit.direction
}

func (unit *FenceUnit) Rotate() {
	unit.direction = unit.direction.Rotate()
}

func (unit *FenceUnit) Delete() {
}
