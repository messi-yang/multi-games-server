package unitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/commonmodel"
)

type UnitRepo interface {
	Add(Unit) error
	Get(UnitId) (Unit, error)
	Delete(Unit) error
	FindUnitAt(sharedkernelmodel.WorldId, commonmodel.Position) (unit Unit, found bool, err error)
	GetUnitsOfWorld(sharedkernelmodel.WorldId) ([]Unit, error)
}
