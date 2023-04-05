package itemmodel

import "github.com/google/uuid"

type ItemIdVo struct {
	id uuid.UUID
}

func NewItemIdVo(uuid uuid.UUID) ItemIdVo {
	return ItemIdVo{
		id: uuid,
	}
}

func (vo ItemIdVo) IsEqual(anotherVo ItemIdVo) bool {
	return vo.id == anotherVo.id
}

func (vo ItemIdVo) Uuid() uuid.UUID {
	return vo.id
}
