package fenceunitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
)

type FenceUnitRepo interface {
	Add(FenceUnit) error
	Update(FenceUnit) error
	Get(unitmodel.UnitId) (FenceUnit, error)
	Delete(FenceUnit) error
}
