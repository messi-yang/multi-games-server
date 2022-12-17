package itemmodel

import "github.com/google/uuid"

type ItemId struct {
	id uuid.UUID
}

func NewItemId(id uuid.UUID) ItemId {
	return ItemId{
		id: id,
	}
}

func (id ItemId) GetId() uuid.UUID {
	return id.id
}
