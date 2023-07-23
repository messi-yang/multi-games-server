package unitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
)

type UnitRepo interface {
	Add(Unit) error
	Get(UnitId) (Unit, error)
	Delete(Unit) error
	GetUnitAt(globalcommonmodel.WorldId, worldcommonmodel.Position) (unit *Unit, err error)
	GetUnitsOfWorld(globalcommonmodel.WorldId) ([]Unit, error)
}
