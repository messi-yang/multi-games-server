package dto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/gamecommonmodel"
)

type UnitDto struct {
	Alive    bool                     `json:"alive"`
	ItemType gamecommonmodel.ItemType `json:"itemType"`
}

func NewUnitDto(unit gamecommonmodel.Unit) UnitDto {
	return UnitDto{
		Alive:    unit.GetAlive(),
		ItemType: unit.GetItemType(),
	}
}

func (dto UnitDto) ToValueObject() gamecommonmodel.Unit {
	return gamecommonmodel.NewUnit(dto.Alive, dto.ItemType)
}
