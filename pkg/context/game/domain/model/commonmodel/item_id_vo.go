package commonmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/domainmodel"
	"github.com/google/uuid"
)

type ItemIdVo struct {
	id uuid.UUID
}

// Interface Implementation Check
var _ domainmodel.ValueObject[ItemIdVo] = (*ItemIdVo)(nil)

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
