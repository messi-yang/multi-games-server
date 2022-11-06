package valueobject

type Item struct {
	itemType ItemType
}

func NewUnit(itemType ItemType) Item {
	return Item{
		itemType: itemType,
	}
}

func (item Item) GetType() ItemType {
	return item.itemType
}
