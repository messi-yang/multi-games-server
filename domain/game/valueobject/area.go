package valueobject

type Area struct {
	from Coordinate
	to   Coordinate
}

func NewArea(from Coordinate, to Coordinate) Area {
	return Area{
		from: from,
		to:   to,
	}
}

func (a Area) GetFrom() Coordinate {
	return a.from
}

func (a Area) GetTo() Coordinate {
	return a.to
}
