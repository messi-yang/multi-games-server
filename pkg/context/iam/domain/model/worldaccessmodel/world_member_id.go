package worldaccessmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/google/uuid"
)

type WorldMemberId struct {
	id uuid.UUID
}

// Interface Implementation Check
var _ domain.ValueObject[WorldMemberId] = (*WorldMemberId)(nil)

func NewWorldMemberId(uuid uuid.UUID) WorldMemberId {
	return WorldMemberId{
		id: uuid,
	}
}

func (itemId WorldMemberId) IsEqual(otherWorldMemberId WorldMemberId) bool {
	return itemId.id == otherWorldMemberId.id
}

func (itemId WorldMemberId) Uuid() uuid.UUID {
	return itemId.id
}
