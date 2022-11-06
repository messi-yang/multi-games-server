package valueobject

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/common/valueobject"
)

type Unit struct {
	alive    bool
	itemType valueobject.ItemType
}

func NewUnit(alive bool, itemType valueobject.ItemType) Unit {
	return Unit{
		alive:    alive,
		itemType: itemType,
	}
}

func (gu Unit) GetAlive() bool {
	return gu.alive
}

func (gu Unit) SetAlive(alive bool) Unit {
	return Unit{
		alive:    alive,
		itemType: gu.itemType,
	}
}

func (gu Unit) GetItemType() valueobject.ItemType {
	return gu.itemType
}

func (gu Unit) SetItemType(itemType valueobject.ItemType) Unit {
	return Unit{
		alive:    gu.alive,
		itemType: itemType,
	}
}
