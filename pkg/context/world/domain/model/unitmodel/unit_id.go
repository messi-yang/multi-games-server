package unitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/google/uuid"
)

type UnitId struct {
	id uuid.UUID
}

// Interface Implementation Check
var _ domain.ValueObject[UnitId] = (*UnitId)(nil)

func NewUnitId(id uuid.UUID) UnitId {
	return UnitId{
		id: id,
	}
}

func (unitId UnitId) IsEqual(otherUnitId UnitId) bool {
	return unitId.id == otherUnitId.id
}

func (unitId UnitId) Uuid() uuid.UUID {
	return unitId.id
}
