package unitmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/worldmodel"
)

type UnitAgg struct {
	worldId   worldmodel.WorldIdVo
	position  commonmodel.PositionVo
	itemId    itemmodel.ItemIdVo
	direction commonmodel.DirectionVo
}

func NewUnitAgg(
	worldId worldmodel.WorldIdVo,
	position commonmodel.PositionVo,
	itemId itemmodel.ItemIdVo,
	direction commonmodel.DirectionVo,
) UnitAgg {
	return UnitAgg{worldId: worldId, position: position, itemId: itemId, direction: direction}
}

func (ua *UnitAgg) GetWorldId() worldmodel.WorldIdVo {
	return ua.worldId
}

func (ua *UnitAgg) GetPosition() commonmodel.PositionVo {
	return ua.position
}

func (ua *UnitAgg) GetItemId() itemmodel.ItemIdVo {
	return ua.itemId
}

func (ua *UnitAgg) GetDirection() commonmodel.DirectionVo {
	return ua.direction
}
