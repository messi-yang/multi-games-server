package unitmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/blockmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
)

type UnitRepo interface {
	Get(UnitId) (Unit, error)
	Update(Unit) error
	HasUnitsInBound(globalcommonmodel.WorldId, worldcommonmodel.Bound) (bool, error)
	GetUnitsOfWorld(globalcommonmodel.WorldId) ([]Unit, error)
	GetUnitsInBlock(globalcommonmodel.WorldId, blockmodel.Block) ([]Unit, error)
}
