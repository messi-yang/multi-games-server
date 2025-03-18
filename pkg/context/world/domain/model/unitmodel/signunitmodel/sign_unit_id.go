package signunitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/google/uuid"
)

type SignUnitId struct {
	id uuid.UUID
}

// Interface Implementation Check
var _ domain.ValueObject[SignUnitId] = (*SignUnitId)(nil)

func NewSignUnitId(uuid uuid.UUID) SignUnitId {
	return SignUnitId{
		id: uuid,
	}
}

func (signUnitId SignUnitId) IsEqual(otherSignUnitId SignUnitId) bool {
	return signUnitId.id == otherSignUnitId.id
}

func (signUnitId SignUnitId) Uuid() uuid.UUID {
	return signUnitId.id
}
