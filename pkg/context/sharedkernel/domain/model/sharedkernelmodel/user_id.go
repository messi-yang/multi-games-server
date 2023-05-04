package sharedkernelmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/common/domain"
	"github.com/google/uuid"
)

type UserId struct {
	id uuid.UUID
}

// Interface Implementation Check
var _ domain.ValueObject[UserId] = (*UserId)(nil)

func NewUserId(uuid uuid.UUID) UserId {
	return UserId{
		id: uuid,
	}
}

func (userId UserId) IsEqual(otherUserId UserId) bool {
	return userId.id == otherUserId.id
}

func (userId UserId) Uuid() uuid.UUID {
	return userId.id
}
