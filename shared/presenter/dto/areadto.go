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
	toCoordinateDto := NewCoordinateDto(area.GetTo())

	return AreaDto{
		From: fromCoordinateDto,
		To:   toCoordinateDto,
	}
}

func (dto AreaDto) ToArea() (valueobject.Area, error) {
	fromCoordinate, err := dto.From.ToCoordinate()
	if err != nil {
		return valueobject.Area{}, err
	}

	toCoordinate, err := dto.To.ToCoordinate()
	if err != nil {
		return valueobject.Area{}, err
	}

	area, err := valueobject.NewArea(
		fromCoordinate,
		toCoordinate,
	)
	if err != nil {
		return valueobject.Area{}, err
	}

	return area, nil
}
