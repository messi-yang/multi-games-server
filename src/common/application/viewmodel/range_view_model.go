package viewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
)

type RangeViewModel struct {
	From LocationViewModel `json:"from"`
	To   LocationViewModel `json:"to"`
}

func NewRangeViewModel(rangeVo commonmodel.RangeVo) RangeViewModel {
	return RangeViewModel{
		From: NewLocationViewModel(rangeVo.GetFrom()),
		To:   NewLocationViewModel(rangeVo.GetTo()),
	}
}

func (dto RangeViewModel) ToValueObject() (commonmodel.RangeVo, error) {
	fromLocation, err := dto.From.ToValueObject()
	if err != nil {
		return commonmodel.RangeVo{}, err
	}

	toLocation, err := dto.To.ToValueObject()
	if err != nil {
		return commonmodel.RangeVo{}, err
	}

	rangeVo, err := commonmodel.NewRangeVo(
		fromLocation,
		toLocation,
	)
	if err != nil {
		return commonmodel.RangeVo{}, err
	}

	return rangeVo, nil
}
