package gamemodel

type ViewVo struct {
	map_  MapVo
	bound BoundVo
}

func NewViewVo(map_ MapVo, bound BoundVo) ViewVo {
	return ViewVo{
		map_:  map_,
		bound: bound,
	}
}

func (view ViewVo) GetMap() MapVo {
	return view.map_
}

func (view ViewVo) GetBound() BoundVo {
	return view.bound
}
