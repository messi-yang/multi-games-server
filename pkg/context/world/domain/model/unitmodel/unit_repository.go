package unitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
)

type UnitRepo interface {
	Add(Unit) error
	Update(Unit) error
	Get(UnitId) (Unit, error)
	Delete(Unit) error
	Find(UnitId) (unit *Unit, err error)
	GetUnitsOfWorld(globalcommonmodel.WorldId) ([]Unit, error)
}
