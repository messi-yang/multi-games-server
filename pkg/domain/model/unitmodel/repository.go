package unitmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/gamemodel"
)

type Repo interface {
	GetAt(gameId gamemodel.GameIdVo, location commonmodel.LocationVo) (UnitAgg, bool)
	GetUnits(gameId gamemodel.GameIdVo, bound commonmodel.BoundVo) []UnitAgg
	UpdateUnit(unit UnitAgg)
	DeleteUnit(gameId gamemodel.GameIdVo, location commonmodel.LocationVo)
}
