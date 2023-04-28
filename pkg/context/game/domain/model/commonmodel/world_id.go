package commonmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/domainmodel"
	"github.com/google/uuid"
)

type WorldId struct {
	id uuid.UUID
}

// Interface Implementation Check
var _ domainmodel.ValueObject[WorldId] = (*WorldId)(nil)

func NewWorldId(uuid uuid.UUID) WorldId {
	return WorldId{
		id: uuid,
	}
}

func (worldId WorldId) IsEqual(otherWorldId WorldId) bool {
	return worldId.id == otherWorldId.id
}

func (worldId WorldId) Uuid() uuid.UUID {
	return worldId.id
}
