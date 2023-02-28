package unitmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/worldmodel"
)

type UnitAgg struct {
	worldId  worldmodel.WorldIdVo
	position commonmodel.PositionVo
	itemId   itemmodel.ItemIdVo
}

func NewUnitAgg(
	worldId worldmodel.WorldIdVo,
	position commonmodel.PositionVo,
	itemId itemmodel.ItemIdVo,
) UnitAgg {
	return UnitAgg{worldId: worldId, position: position, itemId: itemId}
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
