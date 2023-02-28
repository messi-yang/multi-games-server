package worldmodel

import "github.com/google/uuid"

type WorldIdVo struct {
	id uuid.UUID
}

func NewWorldIdVo(uuidStr string) (WorldIdVo, error) {
	id, err := uuid.Parse(uuidStr)
	if err != nil {
		return WorldIdVo{}, err
	}

	return WorldIdVo{
		id: id,
	}, nil
}

func (vo WorldIdVo) IsEqual(otherWorldId WorldIdVo) bool {
	return vo.id == otherWorldId.id
}

func (vo WorldIdVo) ToString() string {
	return vo.id.String()
}

func (vo WorldIdVo) Uuid() uuid.UUID {
	return vo.id
}
