package jsondto

import (
	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/itemmodel"
)

type UnitJsonDto struct {
	Alive  bool   `json:"alive"`
	ItemId string `json:"itemId"`
}

func NewUnitJsonDto(unit gamecommonmodel.Unit) UnitJsonDto {
	return UnitJsonDto{
		Alive:  unit.GetAlive(),
		ItemId: unit.GetItemId().ToString(),
	}
}

func (dto UnitJsonDto) ToValueObject() (gamecommonmodel.Unit, error) {
	itemId, err := itemmodel.NewItemId(dto.ItemId)
	if err != nil {
		return gamecommonmodel.Unit{}, err
	}
	return gamecommonmodel.NewUnit(itemId), nil
}
