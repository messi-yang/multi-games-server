package areadto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/coordinatedto"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/valueobject"
)

type DTO struct {
	From coordinatedto.DTO `json:"from"`
	To   coordinatedto.DTO `json:"to"`
}

func FromDTO(areaDTO DTO) (valueobject.Area, error) {
	fromCoordinate, err := coordinatedto.FromDTO(areaDTO.From)
	if err != nil {
		return valueobject.Area{}, err
	}

	toCoordinate, err := coordinatedto.FromDTO(areaDTO.To)
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
