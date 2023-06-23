package worldaccountmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain"
	"github.com/google/uuid"
)

type WorldAccountId struct {
	id uuid.UUID
}

// Interface Implementation Check
var _ domain.ValueObject[WorldAccountId] = (*WorldAccountId)(nil)

func NewWorldAccountId(uuid uuid.UUID) WorldAccountId {
	return WorldAccountId{
		id: uuid,
	}
}

func (worldAccountId WorldAccountId) IsEqual(otherWorldAccountId WorldAccountId) bool {
	return worldAccountId.id == otherWorldAccountId.id
}

func (worldAccountId WorldAccountId) Uuid() uuid.UUID {
	return worldAccountId.id
}
