package worldmodel

import "github.com/google/uuid"

type WorldIdVo struct {
	id uuid.UUID
}

func NewWorldIdVo(uuid uuid.UUID) WorldIdVo {
	return WorldIdVo{
		id: uuid,
	}
}

func (vo WorldIdVo) IsEqual(otherWorldId WorldIdVo) bool {
	return vo.id == otherWorldId.id
}

func (vo WorldIdVo) Uuid() uuid.UUID {
	return vo.id
}
