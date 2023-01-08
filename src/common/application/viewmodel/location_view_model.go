package viewmodel

import "github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"

type LocationViewModel struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func NewLocationViewModel(location commonmodel.Location) LocationViewModel {
	return LocationViewModel{
		X: location.GetX(),
		Y: location.GetY(),
	}
}

func (dto LocationViewModel) ToValueObject() (commonmodel.Location, error) {
	return commonmodel.NewLocation(dto.X, dto.Y)
}
