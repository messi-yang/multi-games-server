package itemmodel

import "github.com/google/uuid"

type ItemIdVo struct {
	id uuid.UUID
}

func ParseItemIdVo(uuidStr string) (ItemIdVo, error) {
	id, err := uuid.Parse(uuidStr)
	if err != nil {
		return ItemIdVo{}, err
	}

	return ItemIdVo{
		id: id,
	}, nil
}

func NewItemIdVo(uuid uuid.UUID) ItemIdVo {
	return ItemIdVo{
		id: uuid,
	}
}

func (vo ItemIdVo) IsEqual(anotherVo ItemIdVo) bool {
	return vo.id == anotherVo.id
}

func (vo ItemIdVo) String() string {
	return vo.id.String()
}

func (vo ItemIdVo) Uuid() uuid.UUID {
	return vo.id
}
