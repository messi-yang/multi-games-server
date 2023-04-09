package sharedkernelmodel

import "github.com/google/uuid"

type UserIdVo struct {
	id uuid.UUID
}

func NewUserIdVo(uuid uuid.UUID) UserIdVo {
	return UserIdVo{
		id: uuid,
	}
}

func (vo UserIdVo) IsEqual(otherUserId UserIdVo) bool {
	return vo.id == otherUserId.id
}

func (vo UserIdVo) Uuid() uuid.UUID {
	return vo.id
}
