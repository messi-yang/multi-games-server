package dto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/gamecommonmodel"
)

type CoordinateDto struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func NewCoordinateDto(coordinate gamecommonmodel.Coordinate) CoordinateDto {
	return CoordinateDto{
		X: coordinate.GetX(),
		Y: coordinate.GetY(),
	}
}

func NewCoordinateDtos(coordinates []gamecommonmodel.Coordinate) []CoordinateDto {
	coordinateDtos := make([]CoordinateDto, 0)

	for _, coord := range coordinates {
		coordinate := NewCoordinateDto(coord)
		coordinateDtos = append(coordinateDtos, coordinate)
	}

	return coordinateDtos
}

func ParseCoordinateDtos(coordDtos []CoordinateDto) ([]gamecommonmodel.Coordinate, error) {
	coordinates := make([]gamecommonmodel.Coordinate, 0)

	for _, coord := range coordDtos {
		coordinate, err := coord.ToValueObject()
		if err != nil {
			return nil, err
		}
		coordinates = append(coordinates, coordinate)
	}

	return coordinates, nil
}

func (dto CoordinateDto) ToValueObject() (gamecommonmodel.Coordinate, error) {
	return gamecommonmodel.NewCoordinate(dto.X, dto.Y)
}
