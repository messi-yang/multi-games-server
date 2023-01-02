package commonmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
)

type MapUnit struct {
	itemId itemmodel.ItemId
}

func NewMapUnit(itemId itemmodel.ItemId) MapUnit {
	return MapUnit{
		itemId: itemId,
	}
}

func (gu MapUnit) GetItemId() itemmodel.ItemId {
	return gu.itemId
}

func (gu MapUnit) SetItemId(itemId itemmodel.ItemId) MapUnit {
	return MapUnit{
		itemId: itemId,
	}
}
