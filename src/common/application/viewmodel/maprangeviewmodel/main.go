package maprangeviewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/locationviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
)

type ViewModel struct {
	From locationviewmodel.ViewModel `json:"from"`
	To   locationviewmodel.ViewModel `json:"to"`
}

func New(mapRange commonmodel.MapRange) ViewModel {
	fromViewModel := locationviewmodel.New(mapRange.GetFrom())
	ToValueObjectViewModel := locationviewmodel.New(mapRange.GetTo())

	return ViewModel{
		From: fromViewModel,
		To:   ToValueObjectViewModel,
	}
}

func (dto ViewModel) ToValueObject() (commonmodel.MapRange, error) {
	fromLocation, err := dto.From.ToValueObject()
	if err != nil {
		return commonmodel.MapRange{}, err
	}

	ToValueObject, err := dto.To.ToValueObject()
	if err != nil {
		return commonmodel.MapRange{}, err
	}

	mapRange, err := commonmodel.NewMapRange(
		fromLocation,
		ToValueObject,
	)
	if err != nil {
		return commonmodel.MapRange{}, err
	}

	return mapRange, nil
}
