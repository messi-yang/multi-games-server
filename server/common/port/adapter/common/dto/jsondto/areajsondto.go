package jsondto

import (
	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
)

type AreaJsonDto struct {
	From CoordinateJsonDto `json:"from"`
	To   CoordinateJsonDto `json:"to"`
}

func NewAreaJsonDto(area gamecommonmodel.Area) AreaJsonDto {
	fromCoordinateJsonDto := NewCoordinateJsonDto(area.GetFrom())
	ToValueObjectJsonDto := NewCoordinateJsonDto(area.GetTo())

	return AreaJsonDto{
		From: fromCoordinateJsonDto,
		To:   ToValueObjectJsonDto,
	}
}

func (dto AreaJsonDto) ToValueObject() (gamecommonmodel.Area, error) {
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
