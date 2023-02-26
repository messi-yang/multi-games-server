package unitmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
)

type UnitAgg struct {
	gameId   gamemodel.GameIdVo
	position commonmodel.PositionVo
	itemId   itemmodel.ItemIdVo
}

func NewUnitAgg(
	gameId gamemodel.GameIdVo,
	position commonmodel.PositionVo,
	itemId itemmodel.ItemIdVo,
) UnitAgg {
	return UnitAgg{gameId: gameId, position: position, itemId: itemId}
}

func (ua *UnitAgg) GetGameId() gamemodel.GameIdVo {
	return ua.gameId
}

func (ua *UnitAgg) GetPosition() commonmodel.PositionVo {
	return ua.position
}

func (ua *UnitAgg) GetItemId() itemmodel.ItemIdVo {
	return ua.itemId
}
