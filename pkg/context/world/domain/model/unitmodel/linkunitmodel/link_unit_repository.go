package linkunitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
)

type LinkUnitRepo interface {
	Add(LinkUnit) error
	Get(unitmodel.UnitId) (LinkUnit, error)
	Update(LinkUnit) error
	Delete(LinkUnit) error
}
