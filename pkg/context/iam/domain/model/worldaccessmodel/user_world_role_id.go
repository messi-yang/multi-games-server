package worldaccessmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain"
	"github.com/google/uuid"
)

type UserWorldRoleId struct {
	id uuid.UUID
}

// Interface Implementation Check
var _ domain.ValueObject[UserWorldRoleId] = (*UserWorldRoleId)(nil)

func NewUserWorldRoleId(uuid uuid.UUID) UserWorldRoleId {
	return UserWorldRoleId{
		id: uuid,
	}
}

func (itemId UserWorldRoleId) IsEqual(otherUserWorldRoleId UserWorldRoleId) bool {
	return itemId.id == otherUserWorldRoleId.id
}

func (itemId UserWorldRoleId) Uuid() uuid.UUID {
	return itemId.id
}
