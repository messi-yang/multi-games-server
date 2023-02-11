package unitmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
)

type UnitAgg struct {
	gameId   gamemodel.GameIdVo
	location commonmodel.LocationVo
	itemId   itemmodel.ItemIdVo
}

func NewUnitAgg(
	gameId gamemodel.GameIdVo,
	location commonmodel.LocationVo,
	itemId itemmodel.ItemIdVo,
) UnitAgg {
	return UnitAgg{gameId: gameId, location: location, itemId: itemId}
}

func (ua *UnitAgg) GetGameId() gamemodel.GameIdVo {
	return ua.gameId
}

func (ua *UnitAgg) GetLocation() commonmodel.LocationVo {
	return ua.location
}

func (ua *UnitAgg) GetItemId() itemmodel.ItemIdVo {
	return ua.itemId
}
