package unitdto

import "github.com/dum-dum-genius/game-of-liberty-computer/domain/game/valueobject"

type DTO struct {
	Alive bool `json:"alive"`
	Age   int  `json:"age"`
}

func FromDTO(unitDTO DTO) valueobject.Unit {
	return valueobject.NewUnit(unitDTO.Alive, unitDTO.Age)
}

func ToDTO(unit valueobject.Unit) DTO {
	return DTO{
		Alive: unit.GetAlive(),
		Age:   unit.GetAge(),
	}
}

func ToDTOList(units []valueobject.Unit) []DTO {
	unitDTOs := make([]DTO, 0)

	for i := 0; i < len(units); i += 1 {
		unitDTOs = append(unitDTOs, ToDTO(units[i]))
	}

	return unitDTOs
}
