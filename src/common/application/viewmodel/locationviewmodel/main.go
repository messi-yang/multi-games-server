package locationviewmodel

import "github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"

type ViewModel struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func New(location commonmodel.Location) ViewModel {
	return ViewModel{
		X: location.GetX(),
		Y: location.GetY(),
	}
}

func (dto ViewModel) ToValueObject() (commonmodel.Location, error) {
	return commonmodel.NewLocation(dto.X, dto.Y)
}
