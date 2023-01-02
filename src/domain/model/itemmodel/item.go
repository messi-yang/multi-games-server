package itemmodel

type Item struct {
	id   ItemId
	name string
}

func New(id ItemId, name string) Item {
	return Item{
		id:   id,
		name: name,
	}
}

func (item *Item) GetId() ItemId {
	return item.id
}

func (item *Item) GetName() string {
	return item.name
}
