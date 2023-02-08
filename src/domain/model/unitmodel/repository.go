package unitmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/gamemodel"
)

type Repo interface {
	GetUnit(gameId gamemodel.GameIdVo, location commonmodel.LocationVo) (UnitAgg, error)
	GetUnits(gameId gamemodel.GameIdVo, bound commonmodel.BoundVo) []UnitAgg
	UpdateUnit(unit UnitAgg)
	DeleteUnit(gameId gamemodel.GameIdVo, location commonmodel.LocationVo)
}
