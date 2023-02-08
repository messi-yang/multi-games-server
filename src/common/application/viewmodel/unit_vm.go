package viewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/unitmodel"
)

type UnitVm struct {
	ItemId   int16      `json:"itemId"`
	Location LocationVm `json:"location"`
}

func NewUnitVm(unit unitmodel.UnitAgg) UnitVm {
	return UnitVm{
		ItemId:   unit.GetItemId().ToInt16(),
		Location: NewLocationVm(unit.GetLocation()),
	}
}
