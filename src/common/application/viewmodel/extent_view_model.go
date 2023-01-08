package viewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
)

type ExtentViewModel struct {
	From LocationViewModel `json:"from"`
	To   LocationViewModel `json:"to"`
}

func NewExtentViewModel(extent commonmodel.Extent) ExtentViewModel {
	return ExtentViewModel{
		From: NewLocationViewModel(extent.GetFrom()),
		To:   NewLocationViewModel(extent.GetTo()),
	}
}

func (dto ExtentViewModel) ToValueObject() (commonmodel.Extent, error) {
	fromLocation, err := dto.From.ToValueObject()
	if err != nil {
		return commonmodel.Extent{}, err
	}

	toLocation, err := dto.To.ToValueObject()
	if err != nil {
		return commonmodel.Extent{}, err
	}

	extent, err := commonmodel.NewExtent(
		fromLocation,
		toLocation,
	)
	if err != nil {
		return commonmodel.Extent{}, err
	}

	return extent, nil
}
