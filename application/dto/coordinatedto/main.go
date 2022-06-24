package coordinatedto

import "github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/valueobject"

type CoordinateDTO struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func FromDTO(coordinateDTO CoordinateDTO) valueobject.Coordinate {
	return valueobject.NewCoordinate(coordinateDTO.X, coordinateDTO.Y)
}

func FromDTOList(coordDTOs []CoordinateDTO) []valueobject.Coordinate {
	coordinates := make([]valueobject.Coordinate, 0)

	for _, coord := range coordDTOs {
		coordinate := valueobject.NewCoordinate(coord.X, coord.Y)
		coordinates = append(coordinates, coordinate)
	}

	return coordinates
}

func ToDTOList(coordinates []valueobject.Coordinate) []CoordinateDTO {
	coordinateDTOs := make([]CoordinateDTO, 0)

	for _, coord := range coordinates {
		coordinate := CoordinateDTO{
			X: coord.GetX(),
			Y: coord.GetY(),
		}
		coordinateDTOs = append(coordinateDTOs, coordinate)
	}

	return coordinateDTOs
}
