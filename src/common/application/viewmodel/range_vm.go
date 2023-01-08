package viewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
)

type RangeVm struct {
	From LocationVm `json:"from"`
	To   LocationVm `json:"to"`
}

func NewRangeVm(rangeVo commonmodel.Range) RangeVm {
	return RangeVm{
		From: NewLocationVm(rangeVo.GetFrom()),
		To:   NewLocationVm(rangeVo.GetTo()),
	}
}

func (dto RangeVm) ToValueObject() (commonmodel.Range, error) {
	fromLocationVm, err := dto.From.ToValueObject()
	if err != nil {
		return commonmodel.Range{}, err
	}

	toLocationVm, err := dto.To.ToValueObject()
	if err != nil {
		return commonmodel.Range{}, err
	}

	rangeVo, err := commonmodel.NewRange(
		fromLocationVm,
		toLocationVm,
	)
	if err != nil {
		return commonmodel.Range{}, err
	}

	return rangeVo, nil
}
