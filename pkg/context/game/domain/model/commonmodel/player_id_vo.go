package commonmodel

import "github.com/google/uuid"

type PlayerIdVo struct {
	id uuid.UUID
}

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
