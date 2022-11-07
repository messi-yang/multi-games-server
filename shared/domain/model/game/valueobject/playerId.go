package valueobject

import "github.com/google/uuid"

type PlayerId struct {
	id uuid.UUID
}

func NewPlayerId(id uuid.UUID) PlayerId {
	return PlayerId{
		id: id,
	}
}

func (playerId PlayerId) GetId() uuid.UUID {
	return playerId.id
}
