package livegamemodel

import "github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"

type View struct {
	map_   commonmodel.Map
	bound_ commonmodel.Bound
}

func NewView(map_ commonmodel.Map, bound_ commonmodel.Bound) View {
	return View{
		map_:   map_,
		bound_: bound_,
	}
}

func (view View) GetMap() commonmodel.Map {
	return view.map_
}

func (view View) GetBound() commonmodel.Bound {
	return view.bound_
}
