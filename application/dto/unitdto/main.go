package unitdto

import "github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/valueobject"

type UnitDTO struct {
	Alive bool `json:"alive"`
	Age   int  `json:"age"`
}

func FromDTO(unitDTO UnitDTO) valueobject.Unit {
	return valueobject.NewUnit(unitDTO.Alive, unitDTO.Age)
}

func ToDTO(unit valueobject.Unit) UnitDTO {
	return UnitDTO{
		Alive: unit.GetAlive(),
		Age:   unit.GetAge(),
	}
}

func ToDTOList(units []valueobject.Unit) []UnitDTO {
	unitDTOs := make([]UnitDTO, 0)

	for i := 0; i < len(units); i += 1 {
		unitDTOs = append(unitDTOs, ToDTO(units[i]))
	}

	return unitDTOs
}

func FromDTOMatrix(unitDTOMatrix [][]UnitDTO) [][]valueobject.Unit {
	unitMatrix := make([][]valueobject.Unit, 0)

	for i := 0; i < len(unitDTOMatrix); i += 1 {
		unitMatrix = append(unitMatrix, make([]valueobject.Unit, 0))
		for j := 0; j < len(unitDTOMatrix[i]); j += 1 {
			unitMatrix[i] = append(unitMatrix[i], FromDTO(unitDTOMatrix[i][j]))
		}
	}

	return unitMatrix
}

func ToDTOMatrix(units [][]valueobject.Unit) [][]UnitDTO {
	unitDTOMatrix := make([][]UnitDTO, 0)

	for i := 0; i < len(units); i += 1 {
		unitDTOMatrix = append(unitDTOMatrix, make([]UnitDTO, 0))
		for j := 0; j < len(units[i]); j += 1 {
			unitDTOMatrix[i] = append(unitDTOMatrix[i], UnitDTO{
				Alive: units[i][j].GetAlive(),
				Age:   units[i][j].GetAge(),
			})
		}
	}

	return unitDTOMatrix
}
