package viewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/tool"
)

type ViewVm struct {
	Bound BoundVm    `json:"bound"`
	Map   [][]UnitVm `json:"map"`
}

func NewViewVm(view gamemodel.ViewVo) ViewVm {
	unitVmMatrix, _ := tool.MapMatrix(view.GetMap().GetUnitMatrix(), func(colIdx int, rowIdx int, unit commonmodel.UnitVo) (UnitVm, error) {
		return NewUnitVm(unit), nil
	})
	return ViewVm{
		Bound: NewBoundVm(view.GetBound()),
		Map:   unitVmMatrix,
	}
}
