package worldrolemodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain"
	"github.com/google/uuid"
)

type WorldRoleId struct {
	id uuid.UUID
}

// Interface Implementation Check
var _ domain.ValueObject[WorldRoleId] = (*WorldRoleId)(nil)

func NewWorldRoleId(uuid uuid.UUID) WorldRoleId {
	return WorldRoleId{
		id: uuid,
	}
}

func (itemId WorldRoleId) IsEqual(otherWorldRoleId WorldRoleId) bool {
	return itemId.id == otherWorldRoleId.id
}

func (itemId WorldRoleId) Uuid() uuid.UUID {
	return itemId.id
}
