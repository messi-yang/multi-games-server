package viewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/unitmodel"
	"github.com/samber/lo"
)

type ViewVm struct {
	Bound BoundVm  `json:"bound"`
	Units []UnitVm `json:"units"`
}

func NewViewVm(view unitmodel.ViewVo) ViewVm {
	unitVms := lo.Map(view.GetUnits(), func(unit unitmodel.UnitAgg, _ int) UnitVm {
		return NewUnitVm(unit)
	})
	return ViewVm{
		Bound: NewBoundVm(view.GetBound()),
		Units: unitVms,
	}
}
