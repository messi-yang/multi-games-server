package unitdto

import "github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"

type Dto struct {
	Alive bool `json:"alive"`
	Age   int  `json:"age"`
}

func FromDto(unitDto Dto) valueobject.Unit {
	return valueobject.NewUnit(unitDto.Alive, unitDto.Age)
}

func ToDto(unit valueobject.Unit) Dto {
	return Dto{
		Alive: unit.GetAlive(),
		Age:   unit.GetAge(),
	}
}

func ToDtoList(units []valueobject.Unit) []Dto {
	unitDtos := make([]Dto, 0)

	for i := 0; i < len(units); i += 1 {
		unitDtos = append(unitDtos, ToDto(units[i]))
	}

	return unitDtos
}
