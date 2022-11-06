package dto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/game/valueobject"
)

type UnitMapDto [][]UnitDto

func Dto(unitMap valueobject.UnitMap) UnitMapDto {
	unitMapDto := make(UnitMapDto, 0)

	for i := 0; i < unitMap.GetDimension().GetWidth(); i += 1 {
		unitMapDto = append(unitMapDto, make([]UnitDto, 0))
		for j := 0; j < unitMap.GetDimension().GetHeight(); j += 1 {
			coord, _ := valueobject.NewCoordinate(i, j)
			unit := unitMap.GetUnit(coord)
			unitMapDto[i] = append(unitMapDto[i], NewUnitDto(unit))
		}
	}
	return unitMapDto
}

func (dto UnitMapDto) ToValueObject() valueobject.UnitMap {
	unitMatrix := make([][]valueobject.Unit, 0)

	for i := 0; i < len(dto); i += 1 {
		unitMatrix = append(unitMatrix, make([]valueobject.Unit, 0))
		for j := 0; j < len(dto[0]); j += 1 {
			unitMatrix[i] = append(unitMatrix[i], dto[i][j].ToValueObject())
		}
	}
	return valueobject.NewUnitMap(unitMatrix)
}
