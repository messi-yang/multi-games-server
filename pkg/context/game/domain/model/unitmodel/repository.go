package unitmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
)

type Repo interface {
	Add(unit Unit) error
	Delete(commonmodel.WorldId, commonmodel.Position) error
	FindUnitAt(commonmodel.WorldId, commonmodel.Position) (unit Unit, found bool, err error)
	QueryUnitsInBound(commonmodel.WorldId, commonmodel.Bound) ([]Unit, error)
}
