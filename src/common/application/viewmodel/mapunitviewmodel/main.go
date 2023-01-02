package mapunitviewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
	"github.com/google/uuid"
)

type ViewModel struct {
	ItemId *string `json:"itemId"`
}

func New(mapUnit commonmodel.MapUnit) ViewModel {
	var itemId *string = nil
	if mapUnit.GetItemId().IsNotEmpty() {
		itemIdString := mapUnit.GetItemId().ToString()
		itemId = &itemIdString
	}
	return ViewModel{
		ItemId: itemId,
	}
}

func (dto ViewModel) ToValueObject() (commonmodel.MapUnit, error) {
	var rawItemId string = uuid.Nil.String()
	if dto.ItemId != nil {
		rawItemId = *dto.ItemId
	}
	itemId, err := itemmodel.NewItemId(rawItemId)
	if err != nil {
		return commonmodel.MapUnit{}, err
	}
	return commonmodel.NewMapUnit(itemId), nil
}
