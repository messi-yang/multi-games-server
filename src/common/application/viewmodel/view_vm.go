package viewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/tool"
)

type ViewVm struct {
	Bound BoundVm    `json:"bound"`
	Map   [][]UnitVm `json:"map"`
}

func NewViewVm(view livegamemodel.View) ViewVm {
	unitVmMatrix, _ := tool.MapMatrix(view.GetMap().GetUnitMatrix(), func(colIdx int, rowIdx int, unit commonmodel.Unit) (UnitVm, error) {
		return NewUnitVm(unit), nil
	})
	return ViewVm{
		Bound: NewBoundVm(view.GetBound()),
		Map:   unitVmMatrix,
	}
}
