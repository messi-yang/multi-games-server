package dto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/gamecommonmodel"
)

type AreaDto struct {
	From CoordinateDto `json:"from"`
	To   CoordinateDto `json:"to"`
}

func NewAreaDto(area gamecommonmodel.Area) AreaDto {
	fromCoordinateDto := NewCoordinateDto(area.GetFrom())
	ToValueObjectDto := NewCoordinateDto(area.GetTo())

	return AreaDto{
		From: fromCoordinateDto,
		To:   ToValueObjectDto,
	}
}

func (dto AreaDto) ToValueObject() (gamecommonmodel.Area, error) {
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
