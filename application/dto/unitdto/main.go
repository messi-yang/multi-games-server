package unitdto

import "github.com/dum-dum-genius/game-of-liberty-computer/domain/game/valueobject"

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

func FromDTOMap(unitDTOMap [][]UnitDTO) [][]valueobject.Unit {
	unitMap := make([][]valueobject.Unit, 0)

	for i := 0; i < len(unitDTOMap); i += 1 {
		unitMap = append(unitMap, make([]valueobject.Unit, 0))
		for j := 0; j < len(unitDTOMap[i]); j += 1 {
			unitMap[i] = append(unitMap[i], FromDTO(unitDTOMap[i][j]))
		}
	}

	return unitMap
}

func ToDTOMap(units [][]valueobject.Unit) [][]UnitDTO {
	unitDTOMap := make([][]UnitDTO, 0)

	for i := 0; i < len(units); i += 1 {
		unitDTOMap = append(unitDTOMap, make([]UnitDTO, 0))
		for j := 0; j < len(units[i]); j += 1 {
			unitDTOMap[i] = append(unitDTOMap[i], UnitDTO{
				Alive: units[i][j].GetAlive(),
				Age:   units[i][j].GetAge(),
			})
		}
	}

	return unitDTOMap
}
