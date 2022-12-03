package dto

import (
	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
)

type AreaPresenterDto struct {
	From CoordinatePresenterDto `json:"from"`
	To   CoordinatePresenterDto `json:"to"`
}

func NewAreaPresenterDto(area gamecommonmodel.Area) AreaPresenterDto {
	fromCoordinatePresenterDto := NewCoordinatePresenterDto(area.GetFrom())
	ToValueObjectPresenterDto := NewCoordinatePresenterDto(area.GetTo())

	return AreaPresenterDto{
		From: fromCoordinatePresenterDto,
		To:   ToValueObjectPresenterDto,
	}
}

func (dto AreaPresenterDto) ToValueObject() (gamecommonmodel.Area, error) {
	fromCoordinate, err := dto.From.ToValueObject()
	if err != nil {
		return gamecommonmodel.Area{}, err
	}

	ToValueObject, err := dto.To.ToValueObject()
	if err != nil {
		return gamecommonmodel.Area{}, err
	}

	area, err := gamecommonmodel.NewArea(
		fromCoordinate,
		ToValueObject,
	)
	if err != nil {
		return gamecommonmodel.Area{}, err
	}

	return area, nil
}
