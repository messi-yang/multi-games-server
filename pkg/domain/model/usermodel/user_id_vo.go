package usermodel

import "github.com/google/uuid"

type UserIdVo struct {
	id uuid.UUID
}

func ParseUserIdVo(uuidStr string) (UserIdVo, error) {
	id, err := uuid.Parse(uuidStr)
	if err != nil {
		return UserIdVo{}, err
	}

	return UserIdVo{
		id: id,
	}, nil
}

func NewUserIdVo(uuid uuid.UUID) UserIdVo {
	return UserIdVo{
		id: uuid,
	}
}

func (vo UserIdVo) IsEqual(otherUserId UserIdVo) bool {
	return vo.String() == otherUserId.String()
}

func (vo UserIdVo) String() string {
	return vo.id.String()
}

func (vo UserIdVo) Uuid() uuid.UUID {
	return vo.id
}
