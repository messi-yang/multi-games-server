package commonmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
)

type Unit struct {
	itemId itemmodel.ItemId
}

func NewUnit(itemId itemmodel.ItemId) Unit {
	return Unit{
		itemId: itemId,
	}
}

func (gu Unit) GetItemId() itemmodel.ItemId {
	return gu.itemId
}

func (gu Unit) SetItemId(itemId itemmodel.ItemId) Unit {
	return Unit{
		itemId: itemId,
	}
}
