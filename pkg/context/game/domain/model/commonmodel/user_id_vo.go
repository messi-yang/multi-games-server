package commonmodel

import "github.com/google/uuid"

// Referenced to the UserIdVo from Indentity Access context, you should avoid using it as much as you can.
type UserIdVo struct {
	id uuid.UUID
}

func NewUserIdVo(uuid uuid.UUID) UserIdVo {
	return UserIdVo{
		id: uuid,
	}
}

func (vo UserIdVo) Uuid() uuid.UUID {
	return vo.id
}
