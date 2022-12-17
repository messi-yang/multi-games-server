package jsondto

import (
	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
)

type UnitJsonDto struct {
	Alive  bool          `json:"alive"`
	ItemId ItemIdJsonDto `json:"itemId"`
}

func NewUnitJsonDto(unit gamecommonmodel.Unit) UnitJsonDto {
	return UnitJsonDto{
		Alive:  unit.GetAlive(),
		ItemId: NewItemIdJsonDto(unit.GetItemId()),
	}
}

func (dto UnitJsonDto) ToValueObject() (gamecommonmodel.Unit, error) {
	itemId, err := dto.ItemId.ToValueObject()
	if err != nil {
		return gamecommonmodel.Unit{}, err
	}
	return gamecommonmodel.NewUnit(dto.Alive, itemId), nil
}
