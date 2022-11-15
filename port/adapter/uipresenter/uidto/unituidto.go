package uidto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/gamecommonmodel"
)

type UnitUiDto struct {
	Alive    bool                     `json:"alive"`
	ItemType gamecommonmodel.ItemType `json:"itemType"`
}

func NewUnitUiDto(unit gamecommonmodel.Unit) UnitUiDto {
	return UnitUiDto{
		Alive:    unit.GetAlive(),
		ItemType: unit.GetItemType(),
	}
}

func (dto UnitUiDto) ToValueObject() gamecommonmodel.Unit {
	return gamecommonmodel.NewUnit(dto.Alive, dto.ItemType)
}
