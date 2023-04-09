package unitmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
)

type UnitAgg struct {
	worldId   commonmodel.WorldIdVo
	position  commonmodel.PositionVo
	itemId    commonmodel.ItemIdVo
	direction commonmodel.DirectionVo
}

func NewUnitAgg(
	worldId commonmodel.WorldIdVo,
	position commonmodel.PositionVo,
	itemId commonmodel.ItemIdVo,
	direction commonmodel.DirectionVo,
) UnitAgg {
	return UnitAgg{worldId: worldId, position: position, itemId: itemId, direction: direction}
}

func (ua *UnitAgg) GetWorldId() commonmodel.WorldIdVo {
	return ua.worldId
}

func (ua *UnitAgg) GetPosition() commonmodel.PositionVo {
	return ua.position
}

func (ua *UnitAgg) GetItemId() commonmodel.ItemIdVo {
	return ua.itemId
}

func (ua *UnitAgg) GetDirection() commonmodel.DirectionVo {
	return ua.direction
}
