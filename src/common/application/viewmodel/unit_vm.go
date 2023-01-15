package viewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
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
