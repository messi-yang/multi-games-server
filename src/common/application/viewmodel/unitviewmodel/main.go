package unitviewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
	"github.com/google/uuid"
)

type ViewModel struct {
	ItemId *string `json:"itemId"`
}

func New(unit commonmodel.Unit) ViewModel {
	var itemId *string = nil
	if unit.GetItemId().IsNotEmpty() {
		itemIdString := unit.GetItemId().ToString()
		itemId = &itemIdString
	}
	return ViewModel{
		ItemId: itemId,
	}
}

func (dto ViewModel) ToValueObject() (commonmodel.Unit, error) {
	var rawItemId string = uuid.Nil.String()
	if dto.ItemId != nil {
		rawItemId = *dto.ItemId
	}
	itemId, err := itemmodel.NewItemId(rawItemId)
	if err != nil {
		return commonmodel.Unit{}, err
	}
	return commonmodel.NewUnit(itemId), nil
}
