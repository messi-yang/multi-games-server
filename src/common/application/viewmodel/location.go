package viewmodel

import "github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"

type Location struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func NewLocation(location commonmodel.Location) Location {
	return Location{
		X: location.GetX(),
		Y: location.GetY(),
	}
}

func (dto Location) ToValueObject() (commonmodel.Location, error) {
	return commonmodel.NewLocation(dto.X, dto.Y)
}
