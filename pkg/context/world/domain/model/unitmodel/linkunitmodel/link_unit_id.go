package linkunitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/google/uuid"
)

type LinkUnitId struct {
	id uuid.UUID
}

// Interface Implementation Check
var _ domain.ValueObject[LinkUnitId] = (*LinkUnitId)(nil)

func NewLinkUnitId(uuid uuid.UUID) LinkUnitId {
	return LinkUnitId{
		id: uuid,
	}
}

func (portalUnitId LinkUnitId) IsEqual(otherLinkUnitId LinkUnitId) bool {
	return portalUnitId.id == otherLinkUnitId.id
}

func (portalUnitId LinkUnitId) Uuid() uuid.UUID {
	return portalUnitId.id
}
