package livegamemodel

type View struct {
	map_  Map
	bound Bound
}

func NewView(map_ Map, bound Bound) View {
	return View{
		map_:  map_,
		bound: bound,
	}
}

func (view View) GetMap() Map {
	return view.map_
}

func (view View) GetBound() Bound {
	return view.bound
}
