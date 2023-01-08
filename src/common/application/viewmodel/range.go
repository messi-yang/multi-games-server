package viewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
)

type Range struct {
	From Location `json:"from"`
	To   Location `json:"to"`
}

func NewRange(rangeVo commonmodel.Range) Range {
	return Range{
		From: NewLocation(rangeVo.GetFrom()),
		To:   NewLocation(rangeVo.GetTo()),
	}
}

func (dto Range) ToValueObject() (commonmodel.Range, error) {
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
