package dto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/game/domain/valueobject"
)

type UnitDto struct {
	Alive    bool                 `json:"alive"`
	ItemType valueobject.ItemType `json:"itemType"`
}

func NewUnitDto(unit valueobject.Unit) UnitDto {
	return UnitDto{
		Alive:    unit.GetAlive(),
		ItemType: unit.GetItemType(),
	}
}

func (dto UnitDto) ToValueObject() valueobject.Unit {
	return valueobject.NewUnit(dto.Alive, dto.ItemType)
}
