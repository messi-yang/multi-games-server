package itemmodel

import "github.com/google/uuid"

type ItemIdVo struct {
	id uuid.UUID
}

func NewItemIdVo(uuidStr string) (ItemIdVo, error) {
	id, err := uuid.Parse(uuidStr)
	if err != nil {
		return ItemIdVo{}, err
	}

	return ItemIdVo{
		id: id,
	}, nil
}

func (id ItemIdVo) IsEqual(anotherItem ItemIdVo) bool {
	return id.id.String() == anotherItem.ToString()
}

func (id ItemIdVo) IsEmpty() bool {
	return id.id.String() == uuid.Nil.String()
}

func (id ItemIdVo) ToString() string {
	return id.id.String()
}
