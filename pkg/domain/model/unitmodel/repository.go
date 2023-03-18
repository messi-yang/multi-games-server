package unitmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/worldmodel"
)

type Repository interface {
	Add(unit UnitAgg) error
	GetUnitAt(worldId worldmodel.WorldIdVo, position commonmodel.PositionVo) (unit UnitAgg, found bool, err error)
	GetUnitsInBound(worldId worldmodel.WorldIdVo, bound commonmodel.BoundVo) ([]UnitAgg, error)
	Delete(worldId worldmodel.WorldIdVo, position commonmodel.PositionVo) error
}
