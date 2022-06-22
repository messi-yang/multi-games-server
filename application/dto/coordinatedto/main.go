package coordinatedto

import "github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/valueobject"

type CoordinateDTO struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func FromDTOs(coordDTOs []CoordinateDTO) []valueobject.Coordinate {
	coordinates := make([]valueobject.Coordinate, 0)

	for _, coord := range coordDTOs {
		coordinate := valueobject.NewCoordinate(coord.X, coord.Y)
		coordinates = append(coordinates, coordinate)
	}

	return coordinates
}
