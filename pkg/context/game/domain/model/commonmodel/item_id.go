package commonmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/domainmodel"
	"github.com/google/uuid"
)

type ItemId struct {
	id uuid.UUID
}

// Interface Implementation Check
var _ domainmodel.ValueObject[ItemId] = (*ItemId)(nil)

func NewItemId(uuid uuid.UUID) ItemId {
	return ItemId{
		id: uuid,
	}
}

func (itemId ItemId) IsEqual(otherItemId ItemId) bool {
	return itemId.id == otherItemId.id
}

func (itemId ItemId) Uuid() uuid.UUID {
	return itemId.id
}
