package areaviewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/coordinateviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
)

type ViewModel struct {
	From coordinateviewmodel.ViewModel `json:"from"`
	To   coordinateviewmodel.ViewModel `json:"to"`
}

func New(area commonmodel.Area) ViewModel {
	fromViewModel := coordinateviewmodel.New(area.GetFrom())
	ToValueObjectViewModel := coordinateviewmodel.New(area.GetTo())

	return ViewModel{
		From: fromViewModel,
		To:   ToValueObjectViewModel,
	}
}

func (dto ViewModel) ToValueObject() (commonmodel.Area, error) {
	fromCoordinate, err := dto.From.ToValueObject()
	if err != nil {
		return commonmodel.Area{}, err
	}

	ToValueObject, err := dto.To.ToValueObject()
	if err != nil {
		return commonmodel.Area{}, err
	}

	area, err := commonmodel.NewArea(
		fromCoordinate,
		ToValueObject,
	)
	if err != nil {
		return commonmodel.Area{}, err
	}

	return area, nil
}
