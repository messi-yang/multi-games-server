package portalunitmodel

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

func (portalUnitId PortalUnitId) IsEqual(otherPortalUnitId PortalUnitId) bool {
	return portalUnitId.id == otherPortalUnitId.id
}

func (portalUnitId PortalUnitId) Uuid() uuid.UUID {
	return portalUnitId.id
}
