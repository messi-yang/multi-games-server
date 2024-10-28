package colorunitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/google/uuid"
)

type ColorUnitId struct {
	id uuid.UUID
}

// Interface Implementation Check
var _ domain.ValueObject[ColorUnitId] = (*ColorUnitId)(nil)

func NewColorUnitId(uuid uuid.UUID) ColorUnitId {
	return ColorUnitId{
		id: uuid,
	}
}

func (colorUnitId ColorUnitId) IsEqual(otherColorUnitId ColorUnitId) bool {
	return colorUnitId.id == otherColorUnitId.id
}

func (colorUnitId ColorUnitId) Uuid() uuid.UUID {
	return colorUnitId.id
}
