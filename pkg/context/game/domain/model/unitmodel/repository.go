package unitmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
)

type Repository interface {
	Add(unit UnitAgg) error
	GetUnitAt(commonmodel.WorldIdVo, commonmodel.PositionVo) (unit UnitAgg, found bool, err error)
	GetUnitsInBound(commonmodel.WorldIdVo, commonmodel.BoundVo) ([]UnitAgg, error)
	Delete(commonmodel.WorldIdVo, commonmodel.PositionVo) error
}
