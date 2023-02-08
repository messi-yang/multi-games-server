package viewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/unitmodel"
)

type UnitVm struct {
	ItemId   string     `json:"itemId"`
	Location LocationVm `json:"location"`
}

func NewUnitVm(unit unitmodel.UnitAgg) UnitVm {
	return UnitVm{
		ItemId:   unit.GetItemId().ToString(),
		Location: NewLocationVm(unit.GetLocation()),
	}
}
