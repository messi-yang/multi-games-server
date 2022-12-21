package jsondto

import (
	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
)

type CoordinateJsonDto struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func NewCoordinateJsonDto(coordinate gamecommonmodel.Coordinate) CoordinateJsonDto {
	return CoordinateJsonDto{
		X: coordinate.GetX(),
		Y: coordinate.GetY(),
	}
}

func (dto CoordinateJsonDto) ToValueObject() (gamecommonmodel.Coordinate, error) {
	return gamecommonmodel.NewCoordinate(dto.X, dto.Y)
}
