package unitmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
)

type ViewVo struct {
	bound commonmodel.BoundVo
	units []UnitAgg
}

func NewViewVo(bound commonmodel.BoundVo, units []UnitAgg) ViewVo {
	return ViewVo{
		bound: bound,
		units: units,
	}
}

func (view ViewVo) GetUnits() []UnitAgg {
	return view.units
}

func (view ViewVo) GetBound() commonmodel.BoundVo {
	return view.bound
}
