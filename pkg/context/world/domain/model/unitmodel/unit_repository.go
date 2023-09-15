package unitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
)

type UnitRepo interface {
	Get(UnitId) (Unit, error)
	Find(UnitId) (unit *Unit, err error)
	GetUnitsOfWorld(globalcommonmodel.WorldId) ([]Unit, error)
}
