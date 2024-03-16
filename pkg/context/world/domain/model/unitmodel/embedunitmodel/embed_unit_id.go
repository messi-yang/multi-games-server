package embedunitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/google/uuid"
)

type EmbedUnitId struct {
	id uuid.UUID
}

// Interface Implementation Check
var _ domain.ValueObject[EmbedUnitId] = (*EmbedUnitId)(nil)

func NewEmbedUnitId(uuid uuid.UUID) EmbedUnitId {
	return EmbedUnitId{
		id: uuid,
	}
}

func (portalUnitId EmbedUnitId) IsEqual(otherEmbedUnitId EmbedUnitId) bool {
	return portalUnitId.id == otherEmbedUnitId.id
}

func (portalUnitId EmbedUnitId) Uuid() uuid.UUID {
	return portalUnitId.id
}
