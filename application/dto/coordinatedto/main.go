package coordinatedto

import "github.com/dum-dum-genius/game-of-liberty-computer/domain/game/valueobject"

type CoordinateDTO struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func FromDTO(coordinateDTO CoordinateDTO) (valueobject.Coordinate, error) {
	return valueobject.NewCoordinate(coordinateDTO.X, coordinateDTO.Y)
}

func FromDTOList(coordDTOs []CoordinateDTO) ([]valueobject.Coordinate, error) {
	coordinates := make([]valueobject.Coordinate, 0)

	for _, coord := range coordDTOs {
		coordinate, err := valueobject.NewCoordinate(coord.X, coord.Y)
		if err != nil {
			return nil, err
		}
		coordinates = append(coordinates, coordinate)
	}

	return coordinates, nil
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
