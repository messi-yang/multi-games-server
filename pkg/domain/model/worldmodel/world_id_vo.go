package worldmodel

import "github.com/google/uuid"

type WorldIdVo struct {
	id uuid.UUID
}

func ParseWorldIdVo(uuidStr string) (WorldIdVo, error) {
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

func (vo WorldIdVo) String() string {
	return vo.id.String()
}

func (vo WorldIdVo) Uuid() uuid.UUID {
	return vo.id
}
