package jsondto

import "github.com/dum-dum-genius/game-of-liberty-computer/domain/model/commonmodel"

type AreaJsonDto struct {
	From CoordinateJsonDto `json:"from"`
	To   CoordinateJsonDto `json:"to"`
}

func NewAreaJsonDto(area commonmodel.Area) AreaJsonDto {
	fromCoordinateJsonDto := NewCoordinateJsonDto(area.GetFrom())
	ToValueObjectJsonDto := NewCoordinateJsonDto(area.GetTo())

	return AreaJsonDto{
		From: fromCoordinateJsonDto,
		To:   ToValueObjectJsonDto,
	}
}

func (dto AreaJsonDto) ToValueObject() (commonmodel.Area, error) {
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
