package sharedkernelmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/valueobject"
	"github.com/google/uuid"
)

type UserIdVo struct {
	id uuid.UUID
}

// Interface Implementation Check
var _ valueobject.ValueObject[UserIdVo] = (*UserIdVo)(nil)

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
