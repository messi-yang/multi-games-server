package areaviewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/interface/viewmodel/coordinateviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
)

type AreaViewModel struct {
	From coordinateviewmodel.CoordinateViewModel `json:"from"`
	To   coordinateviewmodel.CoordinateViewModel `json:"to"`
}

func New(area commonmodel.Area) AreaViewModel {
	fromCoordinateViewModel := coordinateviewmodel.New(area.GetFrom())
	ToValueObjectViewModel := coordinateviewmodel.New(area.GetTo())

	return AreaViewModel{
		From: fromCoordinateViewModel,
		To:   ToValueObjectViewModel,
	}
}

func (dto AreaViewModel) ToValueObject() (commonmodel.Area, error) {
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
