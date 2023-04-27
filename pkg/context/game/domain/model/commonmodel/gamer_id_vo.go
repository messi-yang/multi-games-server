package commonmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/valueobject"
	"github.com/google/uuid"
)

type GamerIdVo struct {
	id uuid.UUID
}

// Interface Implementation Check
var _ valueobject.ValueObject[GamerIdVo] = (*GamerIdVo)(nil)

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
