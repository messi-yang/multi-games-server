package unitmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/gamemodel"
)

type Repo interface {
	Add(unit UnitAgg) error
	GetUnitAt(gameId gamemodel.GameIdVo, position commonmodel.PositionVo) (UnitAgg, bool, error)
	GetUnitsInBound(gameId gamemodel.GameIdVo, bound commonmodel.BoundVo) ([]UnitAgg, error)
	Delete(gameId gamemodel.GameIdVo, position commonmodel.PositionVo) error
}
