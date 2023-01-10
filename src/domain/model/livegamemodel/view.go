package livegamemodel

import "github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"

type View struct {
	center commonmodel.Location
}

func NewView(center commonmodel.Location) View {
	return View{
		center: center,
	}
}

func (view View) GetCenter() commonmodel.Location {
	return view.center
}
