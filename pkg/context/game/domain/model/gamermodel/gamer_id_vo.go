package gamermodel

import "github.com/google/uuid"

type GamerIdVo struct {
	id uuid.UUID
}

func NewGamerIdVo(uuid uuid.UUID) GamerIdVo {
	return GamerIdVo{
		id: uuid,
	}
}

func (vo GamerIdVo) IsEqual(otherGamerId GamerIdVo) bool {
	return vo.id == otherGamerId.id
}

func (vo GamerIdVo) Uuid() uuid.UUID {
	return vo.id
}
