package areadto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/coordinatedto"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/valueobject"
)

type AreaDTO struct {
	From coordinatedto.CoordinateDTO `json:"from"`
	To   coordinatedto.CoordinateDTO `json:"to"`
}

func FromDTO(areaDTO AreaDTO) valueobject.Area {
	return valueobject.NewArea(
		coordinatedto.FromDTO(areaDTO.From),
		coordinatedto.FromDTO(areaDTO.To),
	)
}
