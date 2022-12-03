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

func NewCoordinateJsonDtos(coordinates []gamecommonmodel.Coordinate) []CoordinateJsonDto {
	coordinateJsonDtos := make([]CoordinateJsonDto, 0)

	for _, coord := range coordinates {
		coordinate := NewCoordinateJsonDto(coord)
		coordinateJsonDtos = append(coordinateJsonDtos, coordinate)
	}

	return coordinateJsonDtos
}

func ParseCoordinateJsonDtos(coordJsonDtos []CoordinateJsonDto) ([]gamecommonmodel.Coordinate, error) {
	coordinates := make([]gamecommonmodel.Coordinate, 0)

	for _, coord := range coordJsonDtos {
		coordinate, err := coord.ToValueObject()
		if err != nil {
			return nil, err
		}
		coordinates = append(coordinates, coordinate)
	}

	return coordinates, nil
}

func (dto CoordinateJsonDto) ToValueObject() (gamecommonmodel.Coordinate, error) {
	return gamecommonmodel.NewCoordinate(dto.X, dto.Y)
}
