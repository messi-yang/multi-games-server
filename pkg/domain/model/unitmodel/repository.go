package unitmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/gamemodel"
)

type Repo interface {
	GetUnitAt(gameId gamemodel.GameIdVo, location commonmodel.LocationVo) (UnitAgg, bool)
	GetUnitsInBound(gameId gamemodel.GameIdVo, bound commonmodel.BoundVo) []UnitAgg
	Update(unit UnitAgg)
	Delete(gameId gamemodel.GameIdVo, location commonmodel.LocationVo)
}
