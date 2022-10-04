package dto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"
)

type AreaDto struct {
	From CoordinateDto `json:"from"`
	To   CoordinateDto `json:"to"`
}

func NewAreaDto(area valueobject.Area) AreaDto {
	fromCoordinateDto := NewCoordinateDto(area.GetFrom())
	ToValueObjectDto := NewCoordinateDto(area.GetTo())

	return AreaDto{
		From: fromCoordinateDto,
		To:   ToValueObjectDto,
	}
}

func (dto AreaDto) ToValueObject() (valueobject.Area, error) {
	fromCoordinate, err := dto.From.ToValueObject()
	if err != nil {
		return valueobject.Area{}, err
	}

	ToValueObject, err := dto.To.ToValueObject()
	if err != nil {
		return valueobject.Area{}, err
	}

	area, err := valueobject.NewArea(
		fromCoordinate,
		ToValueObject,
	)
	if err != nil {
		return valueobject.Area{}, err
	}

	return area, nil
}
