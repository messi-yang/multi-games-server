package item

import "github.com/google/uuid"

type Item struct {
	id   uuid.UUID
	name string
}

func NewItem(uuid uuid.UUID, name string) Item {
	return Item{id: uuid, name: name}
}

func (item *Item) GetId() uuid.UUID {
	return item.id
}

func (item *Item) GetName() string {
	return item.name
}
