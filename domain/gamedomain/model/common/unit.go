package common

type Unit struct {
	alive    bool
	itemType ItemType
}

func NewUnit(alive bool, itemType ItemType) Unit {
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

func (gu Unit) GetItemType() ItemType {
	return gu.itemType
}

func (gu Unit) SetItemType(itemType ItemType) Unit {
	return Unit{
		alive:    gu.alive,
		itemType: itemType,
	}
}
