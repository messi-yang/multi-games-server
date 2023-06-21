package unitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
)

type UnitRepo interface {
	Add(Unit) error
	Get(UnitId) (Unit, error)
	Delete(Unit) error
	FindUnitAt(sharedkernelmodel.WorldId, commonmodel.Position) (unit Unit, found bool, err error)
	QueryUnitsInBound(sharedkernelmodel.WorldId, commonmodel.Bound) ([]Unit, error)
}
