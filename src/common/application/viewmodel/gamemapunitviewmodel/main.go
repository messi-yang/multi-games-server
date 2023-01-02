package gamemapunitviewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
	"github.com/google/uuid"
)

type ViewModel struct {
	ItemId *string `json:"itemId"`
}

func New(gameMapUnit commonmodel.GameMapUnit) ViewModel {
	var itemId *string = nil
	if gameMapUnit.GetItemId().IsNotEmpty() {
		itemIdString := gameMapUnit.GetItemId().ToString()
		itemId = &itemIdString
	}
	return ViewModel{
		ItemId: itemId,
	}
}

func (dto ViewModel) ToValueObject() (commonmodel.GameMapUnit, error) {
	var rawItemId string = uuid.Nil.String()
	if dto.ItemId != nil {
		rawItemId = *dto.ItemId
	}
	itemId, err := itemmodel.NewItemId(rawItemId)
	if err != nil {
		return commonmodel.GameMapUnit{}, err
	}
	return commonmodel.NewGameMapUnit(itemId), nil
}
