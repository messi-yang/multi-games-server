package dto

import "github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"

type CoordinateDto struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func NewCoordinateDto(coordinate valueobject.Coordinate) CoordinateDto {
	return CoordinateDto{
		X: coordinate.GetX(),
		Y: coordinate.GetY(),
	}
}

func NewCoordinateDtos(coordinates []valueobject.Coordinate) []CoordinateDto {
	coordinateDtos := make([]CoordinateDto, 0)

	for _, coord := range coordinates {
		coordinate := NewCoordinateDto(coord)
		coordinateDtos = append(coordinateDtos, coordinate)
	}

	return coordinateDtos
}

func ParseCoordinateDtos(coordDtos []CoordinateDto) ([]valueobject.Coordinate, error) {
	coordinates := make([]valueobject.Coordinate, 0)

	for _, coord := range coordDtos {
		coordinate, err := coord.ToValueObject()
		if err != nil {
			return nil, err
		}
		coordinates = append(coordinates, coordinate)
	}

	return coordinates, nil
}

func (dto CoordinateDto) ToValueObject() (valueobject.Coordinate, error) {
	return valueobject.NewCoordinate(dto.X, dto.Y)
}
