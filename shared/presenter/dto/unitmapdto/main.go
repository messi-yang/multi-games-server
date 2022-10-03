package unitmapdto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/dto/unitdto"
)

type Dto [][]unitdto.Dto

func ToDto(unitMap *valueobject.UnitMap) Dto {
	unitMapDto := make(Dto, 0)

	for i := 0; i < unitMap.GetMapSize().GetWidth(); i += 1 {
		unitMapDto = append(unitMapDto, make([]unitdto.Dto, 0))
		for j := 0; j < unitMap.GetMapSize().GetHeight(); j += 1 {
			coord, _ := valueobject.NewCoordinate(i, j)
			unit := unitMap.GetUnit(coord)
			unitMapDto[i] = append(unitMapDto[i], unitdto.Dto{
				Alive: unit.GetAlive(),
				Age:   unit.GetAge(),
			})
		}
	}
	return unitMapDto
}
