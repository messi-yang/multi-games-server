package commonmodel

import "fmt"

type ErrInvalidArea struct {
	from Location
	to   Location
}

func (e *ErrInvalidArea) Error() string {
	return fmt.Sprintf("from location (%+v) cannot exceed to location (%+v)", e.from, e.to)
}

type Area struct {
	from Location
	to   Location
}

func NewArea(from Location, to Location) (Area, error) {
	if from.x > to.x || from.y > to.y {
		return Area{}, &ErrInvalidArea{from: from, to: to}
	}

	return Area{
		from: from,
		to:   to,
	}, nil
}

func (a Area) GetFrom() Location {
	return a.from
}

func (a Area) GetTo() Location {
	return a.to
}

func (a Area) GetWidth() int {
	return a.to.x - a.from.x + 1
}

func (a Area) GetHeight() int {
	return a.to.y - a.from.y + 1
}

func (a Area) IncludesLocation(location Location) bool {
	return location.x >= a.from.x && location.x <= a.to.x && location.y >= a.from.y && location.y <= a.to.y
}

func (a Area) IncludesAnyLocations(locations []Location) bool {
	locationsInArea := make([]Location, 0)
	for _, location := range locations {
		if a.IncludesLocation(location) {
			locationsInArea = append(locationsInArea, location)
		}
	}

	return len(locationsInArea) > 0
}
