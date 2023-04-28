package unitmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
)

type Repo interface {
	Add(unit UnitAgg) error
	FindUnitAt(commonmodel.WorldIdVo, commonmodel.PositionVo) (unit UnitAgg, found bool, err error)
	QueryUnitsInBound(commonmodel.WorldIdVo, commonmodel.BoundVo) ([]UnitAgg, error)
	Delete(commonmodel.WorldIdVo, commonmodel.PositionVo) error
}
