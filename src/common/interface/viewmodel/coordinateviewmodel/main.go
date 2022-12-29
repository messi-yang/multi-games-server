package coordinateviewmodel

import "github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"

type CoordinateViewModel struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func New(coordinate commonmodel.Coordinate) CoordinateViewModel {
	return CoordinateViewModel{
		X: coordinate.GetX(),
		Y: coordinate.GetY(),
	}
}

func (dto CoordinateViewModel) ToValueObject() (commonmodel.Coordinate, error) {
	return commonmodel.NewCoordinate(dto.X, dto.Y)
}
