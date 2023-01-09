package itemmodel

import "github.com/google/uuid"

type ItemId struct {
	id uuid.UUID
}

func NewItemId(uuidStr string) (ItemId, error) {
	id, err := uuid.Parse(uuidStr)
	if err != nil {
		return ItemId{}, err
	}

	return ItemId{
		id: id,
	}, nil
}

func (id ItemId) IsEmpty() bool {
	return id.id.String() == uuid.Nil.String()
}

func (id ItemId) ToString() string {
	return id.id.String()
}
