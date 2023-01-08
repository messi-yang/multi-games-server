package viewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
)

type MapRangeViewModel struct {
	From LocationViewModel `json:"from"`
	To   LocationViewModel `json:"to"`
}

func NewMapRangeViewModel(mapRange commonmodel.MapRange) MapRangeViewModel {
	return MapRangeViewModel{
		From: NewLocationViewModel(mapRange.GetFrom()),
		To:   NewLocationViewModel(mapRange.GetTo()),
	}
}

func (dto MapRangeViewModel) ToValueObject() (commonmodel.MapRange, error) {
	fromLocation, err := dto.From.ToValueObject()
	if err != nil {
		return commonmodel.MapRange{}, err
	}

	toLocation, err := dto.To.ToValueObject()
	if err != nil {
		return commonmodel.MapRange{}, err
	}

	mapRange, err := commonmodel.NewMapRange(
		fromLocation,
		toLocation,
	)
	if err != nil {
		return commonmodel.MapRange{}, err
	}

	return mapRange, nil
}
