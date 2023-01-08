package viewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
)

type RangeViewModel struct {
	From LocationViewModel `json:"from"`
	To   LocationViewModel `json:"to"`
}

func NewRangeViewModel(rangeVo commonmodel.Range) RangeViewModel {
	return RangeViewModel{
		From: NewLocationViewModel(rangeVo.GetFrom()),
		To:   NewLocationViewModel(rangeVo.GetTo()),
	}
}

func (dto RangeViewModel) ToValueObject() (commonmodel.Range, error) {
	fromLocation, err := dto.From.ToValueObject()
	if err != nil {
		return commonmodel.Range{}, err
	}

	toLocation, err := dto.To.ToValueObject()
	if err != nil {
		return commonmodel.Range{}, err
	}

	rangeVo, err := commonmodel.NewRange(
		fromLocation,
		toLocation,
	)
	if err != nil {
		return commonmodel.Range{}, err
	}

	return rangeVo, nil
}
