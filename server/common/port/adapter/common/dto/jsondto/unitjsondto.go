package jsondto

import (
	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
)

type UnitJsonDto struct {
	Alive    bool                     `json:"alive"`
	ItemType gamecommonmodel.ItemType `json:"itemType"`
}

func NewUnitJsonDto(unit gamecommonmodel.Unit) UnitJsonDto {
	return UnitJsonDto{
		Alive:    unit.GetAlive(),
		ItemType: unit.GetItemType(),
	}
}

func (dto UnitJsonDto) ToValueObject() gamecommonmodel.Unit {
	return gamecommonmodel.NewUnit(dto.Alive, dto.ItemType)
}
