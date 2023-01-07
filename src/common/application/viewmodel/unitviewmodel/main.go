package unitviewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type ViewModel struct {
	ItemId *string `json:"itemId"`
}

func New(unit commonmodel.Unit) ViewModel {
	var itemId *string = lo.Ternary(unit.GetItemId().IsEmpty(), nil, lo.ToPtr(unit.GetItemId().ToString()))
	return ViewModel{
		ItemId: itemId,
	}
}

func (dto ViewModel) ToValueObject() (commonmodel.Unit, error) {
	var rawItemId string = lo.Ternary(dto.ItemId == nil, *dto.ItemId, uuid.Nil.String())
	itemId, err := itemmodel.NewItemId(rawItemId)
	if err != nil {
		return commonmodel.Unit{}, err
	}
	return commonmodel.NewUnit(itemId), nil
}
