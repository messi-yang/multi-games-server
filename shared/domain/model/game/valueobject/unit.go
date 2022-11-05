package valueobject

import "github.com/google/uuid"

type Unit struct {
	alive  bool
	itemId uuid.UUID
}

func NewUnit(alive bool, itemId uuid.UUID) Unit {
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

func (gu Unit) GetItemId() uuid.UUID {
	return gu.itemId
}

func (gu Unit) SetItemId(itemId uuid.UUID) Unit {
	return Unit{
		alive:  gu.alive,
		itemId: itemId,
	}
}
