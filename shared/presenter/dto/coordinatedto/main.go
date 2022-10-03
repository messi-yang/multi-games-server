package coordinatedto

import "github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"

type Dto struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func ToDto(coordinate valueobject.Coordinate) Dto {
	return Dto{
		X: coordinate.GetX(),
		Y: coordinate.GetY(),
	}
}

func FromDto(coordinateDto Dto) (valueobject.Coordinate, error) {
	return valueobject.NewCoordinate(coordinateDto.X, coordinateDto.Y)
}

func FromDtoList(coordDtos []Dto) ([]valueobject.Coordinate, error) {
	coordinates := make([]valueobject.Coordinate, 0)

	for _, coord := range coordDtos {
		coordinate, err := valueobject.NewCoordinate(coord.X, coord.Y)
		if err != nil {
			return nil, err
		}
		coordinates = append(coordinates, coordinate)
	}

	return coordinates, nil
}

func ToDtoList(coordinates []valueobject.Coordinate) []Dto {
	coordinateDtos := make([]Dto, 0)

	for _, coord := range coordinates {
		coordinate := Dto{
			X: coord.GetX(),
			Y: coord.GetY(),
		}
		coordinateDtos = append(coordinateDtos, coordinate)
	}

	return coordinateDtos
}
