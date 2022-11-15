package uidto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/gamecommonmodel"
)

type CoordinateUiDto struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func NewCoordinateUiDto(coordinate gamecommonmodel.Coordinate) CoordinateUiDto {
	return CoordinateUiDto{
		X: coordinate.GetX(),
		Y: coordinate.GetY(),
	}
}

func NewCoordinateUiDtos(coordinates []gamecommonmodel.Coordinate) []CoordinateUiDto {
	coordinateUiDtos := make([]CoordinateUiDto, 0)

	for _, coord := range coordinates {
		coordinate := NewCoordinateUiDto(coord)
		coordinateUiDtos = append(coordinateUiDtos, coordinate)
	}

	return coordinateUiDtos
}

func ParseCoordinateUiDtos(coordUiDtos []CoordinateUiDto) ([]gamecommonmodel.Coordinate, error) {
	coordinates := make([]gamecommonmodel.Coordinate, 0)

	for _, coord := range coordUiDtos {
		coordinate, err := coord.ToValueObject()
		if err != nil {
			return nil, err
		}
		coordinates = append(coordinates, coordinate)
	}

	return coordinates, nil
}

func (dto CoordinateUiDto) ToValueObject() (gamecommonmodel.Coordinate, error) {
	return gamecommonmodel.NewCoordinate(dto.X, dto.Y)
}
