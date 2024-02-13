package fenceunitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/google/uuid"
)

type FenceUnitId struct {
	id uuid.UUID
}

// Interface Implementation Check
var _ domain.ValueObject[FenceUnitId] = (*FenceUnitId)(nil)

func NewFenceUnitId(uuid uuid.UUID) FenceUnitId {
	return FenceUnitId{
		id: uuid,
	}
}

func (unitId FenceUnitId) IsEqual(otherFenceUnitId FenceUnitId) bool {
	return unitId.id == otherFenceUnitId.id
}

func (unitId FenceUnitId) Uuid() uuid.UUID {
	return unitId.id
}
