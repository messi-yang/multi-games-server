package dto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/game/valueobject"
	"github.com/google/uuid"
)

type UnitDto struct {
	Alive  bool      `json:"alive"`
	ItemId uuid.UUID `json:"id"`
}

func NewUnitDto(unit valueobject.Unit) UnitDto {
	return UnitDto{
		Alive:  unit.GetAlive(),
		ItemId: unit.GetItemId(),
	}
}

func (dto UnitDto) ToValueObject() valueobject.Unit {
	return valueobject.NewUnit(dto.Alive, dto.ItemId)
}
