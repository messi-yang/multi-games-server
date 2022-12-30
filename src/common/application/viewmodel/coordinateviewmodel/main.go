package coordinateviewmodel

import "github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"

type ViewModel struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func New(coordinate commonmodel.Coordinate) ViewModel {
	return ViewModel{
		X: coordinate.GetX(),
		Y: coordinate.GetY(),
	}
}

func (dto ViewModel) ToValueObject() (commonmodel.Coordinate, error) {
	return commonmodel.NewCoordinate(dto.X, dto.Y)
}
