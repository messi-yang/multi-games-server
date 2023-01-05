package mapunitviewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type ViewModel struct {
	ItemId *string `json:"itemId"`
}

func New(mapUnit commonmodel.MapUnit) ViewModel {
	var itemId *string = lo.Ternary(mapUnit.GetItemId().IsEmpty(), nil, lo.ToPtr(mapUnit.GetItemId().ToString()))
	return ViewModel{
		ItemId: itemId,
	}
}

func (dto ViewModel) ToValueObject() (commonmodel.MapUnit, error) {
	var rawItemId string = lo.Ternary(dto.ItemId == nil, *dto.ItemId, uuid.Nil.String())
	itemId, err := itemmodel.NewItemId(rawItemId)
	if err != nil {
		return commonmodel.MapUnit{}, err
	}
	return commonmodel.NewMapUnit(itemId), nil
}
