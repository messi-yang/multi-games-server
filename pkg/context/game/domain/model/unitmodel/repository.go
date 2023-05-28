package unitmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
)

type Repo interface {
	Add(Unit) error
	Delete(Unit) error
	FindUnitAt(sharedkernelmodel.WorldId, commonmodel.Position) (unit Unit, found bool, err error)
	QueryUnitsInBound(sharedkernelmodel.WorldId, commonmodel.Bound) ([]Unit, error)
}
