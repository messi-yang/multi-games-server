package areadto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/dto/coordinatedto"
)

type Dto struct {
	From coordinatedto.Dto `json:"from"`
	To   coordinatedto.Dto `json:"to"`
}

func FromDto(areaDto Dto) (valueobject.Area, error) {
	fromCoordinate, err := coordinatedto.FromDto(areaDto.From)
	if err != nil {
		return valueobject.Area{}, err
	}

	toCoordinate, err := coordinatedto.FromDto(areaDto.To)
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
