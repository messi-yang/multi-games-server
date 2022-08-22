package unitmapdto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/unitdto"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/valueobject"
)

type UnitMapDTO [][]unitdto.UnitDTO

func ToDTO(unitMap valueobject.UnitMap) UnitMapDTO {
	unitMapDTO := make(UnitMapDTO, 0)

	for i := 0; i < unitMap.GetMapSize().GetWidth(); i += 1 {
		unitMapDTO = append(unitMapDTO, make([]unitdto.UnitDTO, 0))
		for j := 0; j < unitMap.GetMapSize().GetHeight(); j += 1 {
			coord := valueobject.NewCoordinate(i, j)
			unit := unitMap.GetUnit(coord)
			unitMapDTO[i] = append(unitMapDTO[i], unitdto.UnitDTO{
				Alive: unit.GetAlive(),
				Age:   unit.GetAge(),
			})
		}
	}
	return unitMapDTO
}
