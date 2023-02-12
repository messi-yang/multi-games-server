package dto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/unitmodel"
)

type UnitDto struct {
	ItemId   int16       `json:"itemId"`
	Location LocationDto `json:"location"`
}

func NewUnitDto(unit unitmodel.UnitAgg) UnitDto {
	return UnitDto{
		ItemId:   unit.GetItemId().ToInt16(),
		Location: NewLocationDto(unit.GetLocation()),
	}
}
