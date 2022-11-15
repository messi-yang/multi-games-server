package uidto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/gamecommonmodel"
)

type AreaUiDto struct {
	From CoordinateUiDto `json:"from"`
	To   CoordinateUiDto `json:"to"`
}

func NewAreaUiDto(area gamecommonmodel.Area) AreaUiDto {
	fromCoordinateUiDto := NewCoordinateUiDto(area.GetFrom())
	ToValueObjectUiDto := NewCoordinateUiDto(area.GetTo())

	return AreaUiDto{
		From: fromCoordinateUiDto,
		To:   ToValueObjectUiDto,
	}
}

func (dto AreaUiDto) ToValueObject() (gamecommonmodel.Area, error) {
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
