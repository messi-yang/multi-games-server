package commonmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
)

type GameMapUnit struct {
	itemId itemmodel.ItemId
}

func NewGameMapUnit(itemId itemmodel.ItemId) GameMapUnit {
	return GameMapUnit{
		itemId: itemId,
	}
}

func (gu GameMapUnit) GetItemId() itemmodel.ItemId {
	return gu.itemId
}

func (gu GameMapUnit) SetItemId(itemId itemmodel.ItemId) GameMapUnit {
	return GameMapUnit{
		itemId: itemId,
	}
}
