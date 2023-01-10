package viewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
)

type RangeVm struct {
	From LocationVm `json:"from"`
	To   LocationVm `json:"to"`
}

func NewRangeVm(range_ commonmodel.Range) RangeVm {
	return RangeVm{
		From: NewLocationVm(range_.GetFrom()),
		To:   NewLocationVm(range_.GetTo()),
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

	range_, err := commonmodel.NewRange(
		fromLocationVm,
		toLocationVm,
	)
	if err != nil {
		return commonmodel.Range{}, err
	}

	return range_, nil
}
