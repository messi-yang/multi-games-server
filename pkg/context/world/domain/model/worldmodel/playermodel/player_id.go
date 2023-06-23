package playermodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain"
	"github.com/google/uuid"
)

type PlayerId struct {
	id uuid.UUID
}

// Interface Implementation Check
var _ domain.ValueObject[PlayerId] = (*PlayerId)(nil)

func NewPlayerId(uuid uuid.UUID) PlayerId {
	return PlayerId{
		id: uuid,
	}
}

func (playerId PlayerId) IsEqual(otherPlayerId PlayerId) bool {
	return playerId.id == otherPlayerId.id
}

func (playerId PlayerId) Uuid() uuid.UUID {
	return playerId.id
}
