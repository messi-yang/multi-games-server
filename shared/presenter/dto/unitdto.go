package dto

import "github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"

type UnitDto struct {
	Alive bool `json:"alive"`
	Age   int  `json:"age"`
}

func NewUnitDto(unit valueobject.Unit) UnitDto {
	return UnitDto{
		Alive: unit.GetAlive(),
		Age:   unit.GetAge(),
	}
}

func (dto UnitDto) ToValueObject() valueobject.Unit {
	return valueobject.NewUnit(dto.Alive, dto.Age)
}
