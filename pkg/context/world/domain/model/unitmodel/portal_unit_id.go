package unitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/google/uuid"
)

type PortalUnitId struct {
	id uuid.UUID
}

// Interface Implementation Check
var _ domain.ValueObject[PortalUnitId] = (*PortalUnitId)(nil)

func NewPortalUnitId(uuid uuid.UUID) PortalUnitId {
	return PortalUnitId{
		id: uuid,
	}
}

func (portalUnitid PortalUnitId) IsEqual(otherPortalUnitId PortalUnitId) bool {
	return portalUnitid.id == otherPortalUnitId.id
}

func (portalUnitid PortalUnitId) Uuid() uuid.UUID {
	return portalUnitid.id
}
