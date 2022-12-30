package unitviewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
)

type ViewModel struct {
	Alive  bool   `json:"alive"`
	ItemId string `json:"itemId"`
}

func New(unit commonmodel.Unit) ViewModel {
	return ViewModel{
		Alive:  unit.GetAlive(),
		ItemId: unit.GetItemId().ToString(),
	}
}

func (dto ViewModel) ToValueObject() (commonmodel.Unit, error) {
	itemId, err := itemmodel.NewItemId(dto.ItemId)
	if err != nil {
		return commonmodel.Unit{}, err
	}
	return commonmodel.NewUnit(itemId), nil
}
