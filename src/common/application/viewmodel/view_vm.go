package viewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/livegamemodel"
)

type ViewVm struct {
	Bound BoundVm `json:"bound"`
	Map   MapVm   `json:"map"`
}

func NewViewVm(view livegamemodel.View) ViewVm {
	return ViewVm{
		Bound: NewBoundVm(view.GetBound()),
		Map:   NewMapVm(view.GetMap()),
	}
}
