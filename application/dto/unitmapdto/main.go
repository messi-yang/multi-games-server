package unitmapdto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/unitdto"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/valueobject"
)

type UnitMapDTO [][]unitdto.UnitDTO

func ToDTO(unitMap valueobject.UnitMap) UnitMapDTO {
	unitMapDTO := make(UnitMapDTO, 0)

	for i := 0; i < len(unitMap); i += 1 {
		unitMapDTO = append(unitMapDTO, make([]unitdto.UnitDTO, 0))
		for j := 0; j < len(unitMap[i]); j += 1 {
			unitMapDTO[i] = append(unitMapDTO[i], unitdto.UnitDTO{
				Alive: unitMap[i][j].GetAlive(),
				Age:   unitMap[i][j].GetAge(),
			})
		}
	}
	return unitMapDTO
}
