package unitmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/worldmodel"
)

type Repository interface {
	Add(unit UnitAgg) error
	GetUnitAt(worldmodel.WorldIdVo, commonmodel.PositionVo) (unit UnitAgg, found bool, err error)
	GetUnitsInBound(worldmodel.WorldIdVo, commonmodel.BoundVo) ([]UnitAgg, error)
	Delete(worldmodel.WorldIdVo, commonmodel.PositionVo) error
}
