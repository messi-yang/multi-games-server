package viewmodel

import "github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/livegamemodel"

type ViewVm struct {
	Center LocationVm `json:"center"`
}

func NewViewVm(view livegamemodel.View) ViewVm {
	return ViewVm{
		Center: NewLocationVm(view.GetCenter()),
	}
}
