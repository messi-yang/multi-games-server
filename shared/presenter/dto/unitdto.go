package dto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"
	"github.com/google/uuid"
)

type UnitDto struct {
	Alive bool      `json:"alive"`
	Id    uuid.UUID `json:"id"`
}

func NewUnitDto(unit valueobject.Unit) UnitDto {
	return UnitDto{
		Alive: unit.GetAlive(),
		Id:    unit.GetId(),
	}
}

func (dto UnitDto) ToValueObject() valueobject.Unit {
	return valueobject.NewUnit(dto.Alive, dto.Id)
}
