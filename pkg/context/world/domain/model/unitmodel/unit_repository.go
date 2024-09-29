package unitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
)

type UnitRepo interface {
	Get(UnitId) (Unit, error)
	Find(globalcommonmodel.WorldId, worldcommonmodel.Position) (unit *Unit, err error)
	GetUnitsOfWorld(globalcommonmodel.WorldId) ([]Unit, error)
	GetUnitsInBlock(globalcommonmodel.WorldId, worldcommonmodel.Block) ([]Unit, error)
}
