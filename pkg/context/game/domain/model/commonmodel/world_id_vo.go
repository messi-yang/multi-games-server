package commonmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/domainmodel"
	"github.com/google/uuid"
)

type WorldIdVo struct {
	id uuid.UUID
}

// Interface Implementation Check
var _ domainmodel.ValueObject[WorldIdVo] = (*WorldIdVo)(nil)

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
