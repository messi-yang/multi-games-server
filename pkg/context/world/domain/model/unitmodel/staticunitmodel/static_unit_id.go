package staticunitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/google/uuid"
)

type StaticUnitId struct {
	id uuid.UUID
}

// Interface Implementation Check
var _ domain.ValueObject[StaticUnitId] = (*StaticUnitId)(nil)

func NewStaticUnitId(uuid uuid.UUID) StaticUnitId {
	return StaticUnitId{
		id: uuid,
	}
}

func (unitId StaticUnitId) IsEqual(otherStaticUnitId StaticUnitId) bool {
	return unitId.id == otherStaticUnitId.id
}

func (unitId StaticUnitId) Uuid() uuid.UUID {
	return unitId.id
}
