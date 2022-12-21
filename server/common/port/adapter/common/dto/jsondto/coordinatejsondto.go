package jsondto

import "github.com/dum-dum-genius/game-of-liberty-computer/domain/model/commonmodel"

type CoordinateJsonDto struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func NewCoordinateJsonDto(coordinate commonmodel.Coordinate) CoordinateJsonDto {
	return CoordinateJsonDto{
		X: coordinate.GetX(),
		Y: coordinate.GetY(),
	}
}

func (dto CoordinateJsonDto) ToValueObject() (commonmodel.Coordinate, error) {
	return commonmodel.NewCoordinate(dto.X, dto.Y)
}
