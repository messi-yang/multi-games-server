package presenterdto

import (
	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
)

type CoordinatePresenterDto struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func NewCoordinatePresenterDto(coordinate gamecommonmodel.Coordinate) CoordinatePresenterDto {
	return CoordinatePresenterDto{
		X: coordinate.GetX(),
		Y: coordinate.GetY(),
	}
}

func NewCoordinatePresenterDtos(coordinates []gamecommonmodel.Coordinate) []CoordinatePresenterDto {
	coordinatePresenterDtos := make([]CoordinatePresenterDto, 0)

	for _, coord := range coordinates {
		coordinate := NewCoordinatePresenterDto(coord)
		coordinatePresenterDtos = append(coordinatePresenterDtos, coordinate)
	}

	return coordinatePresenterDtos
}

func ParseCoordinatePresenterDtos(coordPresenterDtos []CoordinatePresenterDto) ([]gamecommonmodel.Coordinate, error) {
	coordinates := make([]gamecommonmodel.Coordinate, 0)

	for _, coord := range coordPresenterDtos {
		coordinate, err := coord.ToValueObject()
		if err != nil {
			return nil, err
		}
		coordinates = append(coordinates, coordinate)
	}

	return coordinates, nil
}

func (dto CoordinatePresenterDto) ToValueObject() (gamecommonmodel.Coordinate, error) {
	return gamecommonmodel.NewCoordinate(dto.X, dto.Y)
}
