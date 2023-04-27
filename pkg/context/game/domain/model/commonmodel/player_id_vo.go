package commonmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/valueobject"
	"github.com/google/uuid"
)

type PlayerIdVo struct {
	id uuid.UUID
}

// Interface Implementation Check
var _ valueobject.ValueObject[PlayerIdVo] = (*PlayerIdVo)(nil)

func NewPlayerIdVo(uuid uuid.UUID) PlayerIdVo {
	return PlayerIdVo{
		id: uuid,
	}
}

func (playerId PlayerIdVo) IsEqual(otherPlayerId PlayerIdVo) bool {
	return playerId.id == otherPlayerId.id
}

func (playerId PlayerIdVo) Uuid() uuid.UUID {
	return playerId.id
}
