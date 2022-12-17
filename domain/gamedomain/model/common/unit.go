package common

import "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/itemmodel"

type Unit struct {
	alive  bool
	itemId itemmodel.ItemId
}

func NewUnit(alive bool, itemId itemmodel.ItemId) Unit {
	return Unit{
		alive:  alive,
		itemId: itemId,
	}
}

func (gu Unit) GetAlive() bool {
	return gu.alive
}

func (gu Unit) SetAlive(alive bool) Unit {
	return Unit{
		alive:  alive,
		itemId: gu.itemId,
	}
}

func (gu Unit) GetItemId() itemmodel.ItemId {
	return gu.itemId
}

func (gu Unit) SetItemId(itemId itemmodel.ItemId) Unit {
	return Unit{
		alive:  gu.alive,
		itemId: itemId,
	}
}
