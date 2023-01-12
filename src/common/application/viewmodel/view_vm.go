package viewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/livegamemodel"
)

type ViewVm struct {
	Range RangeVm `json:"range"`
	Map   MapVm   `json:"map"`
}

func NewViewVm(view livegamemodel.View) ViewVm {
	return ViewVm{
		Range: NewRangeVm(view.GetRange()),
		Map:   NewMapVm(view.GetMap()),
	}
}
