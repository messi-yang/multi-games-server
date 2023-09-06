package staticunitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
)

type StaticUnitRepo interface {
	Add(StaticUnit) error
	Get(unitmodel.UnitId) (StaticUnit, error)
	Delete(StaticUnit) error
}
