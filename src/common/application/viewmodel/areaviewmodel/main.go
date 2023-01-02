package areaviewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/locationviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
)

type ViewModel struct {
	From locationviewmodel.ViewModel `json:"from"`
	To   locationviewmodel.ViewModel `json:"to"`
}

func New(area commonmodel.Area) ViewModel {
	fromViewModel := locationviewmodel.New(area.GetFrom())
	ToValueObjectViewModel := locationviewmodel.New(area.GetTo())

	return ViewModel{
		From: fromViewModel,
		To:   ToValueObjectViewModel,
	}
}

func (dto ViewModel) ToValueObject() (commonmodel.Area, error) {
	fromLocation, err := dto.From.ToValueObject()
	if err != nil {
		return commonmodel.Area{}, err
	}

	ToValueObject, err := dto.To.ToValueObject()
	if err != nil {
		return commonmodel.Area{}, err
	}

	area, err := commonmodel.NewArea(
		fromLocation,
		ToValueObject,
	)
	if err != nil {
		return commonmodel.Area{}, err
	}

	return area, nil
}
