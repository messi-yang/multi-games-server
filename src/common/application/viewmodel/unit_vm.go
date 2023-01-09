package viewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type UnitVm struct {
	ItemId *string `json:"itemId"`
}

func NewUnitVm(unit commonmodel.Unit) UnitVm {
	var itemId *string = lo.Ternary(unit.GetItemId().IsEmpty(), nil, lo.ToPtr(unit.GetItemId().ToString()))
	return UnitVm{
		ItemId: itemId,
	}
}

func (dto UnitVm) ToValueObject() (commonmodel.Unit, error) {
	var itemIdVm string = lo.Ternary(dto.ItemId == nil, *dto.ItemId, uuid.Nil.String())
	itemId, err := itemmodel.NewItemId(itemIdVm)
	if err != nil {
		return commonmodel.Unit{}, err
	}
	return commonmodel.NewUnit(itemId), nil
}
