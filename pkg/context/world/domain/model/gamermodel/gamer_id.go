package gamermodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain"
	"github.com/google/uuid"
)

type GamerId struct {
	id uuid.UUID
}

// Interface Implementation Check
var _ domain.ValueObject[GamerId] = (*GamerId)(nil)

func NewGamerId(uuid uuid.UUID) GamerId {
	return GamerId{
		id: uuid,
	}
}

func (gamerId GamerId) IsEqual(otherGamerId GamerId) bool {
	return gamerId.id == otherGamerId.id
}

func (gamerId GamerId) Uuid() uuid.UUID {
	return gamerId.id
}
