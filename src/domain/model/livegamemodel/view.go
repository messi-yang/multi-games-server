package livegamemodel

import "github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"

type View struct {
	map_   commonmodel.Map
	range_ commonmodel.Range
}

func NewView(map_ commonmodel.Map, range_ commonmodel.Range) View {
	return View{
		map_:   map_,
		range_: range_,
	}
}

func (view View) GetMap() commonmodel.Map {
	return view.map_
}

func (view View) GetRange() commonmodel.Range {
	return view.range_
}
