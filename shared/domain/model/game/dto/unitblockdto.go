package dto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/game/valueobject"
)

type UnitBlockDto [][]UnitDto

func Dto(unitBlock valueobject.UnitBlock) UnitBlockDto {
	unitBlockDto := make(UnitBlockDto, 0)

	for i := 0; i < unitBlock.GetDimension().GetWidth(); i += 1 {
		unitBlockDto = append(unitBlockDto, make([]UnitDto, 0))
		for j := 0; j < unitBlock.GetDimension().GetHeight(); j += 1 {
			coord, _ := valueobject.NewCoordinate(i, j)
			unit := unitBlock.GetUnit(coord)
			unitBlockDto[i] = append(unitBlockDto[i], NewUnitDto(unit))
		}
	}
	return unitBlockDto
}

func (dto UnitBlockDto) ToValueObject() valueobject.UnitBlock {
	unitMatrix := make([][]valueobject.Unit, 0)

	for i := 0; i < len(dto); i += 1 {
		unitMatrix = append(unitMatrix, make([]valueobject.Unit, 0))
		for j := 0; j < len(dto[0]); j += 1 {
			unitMatrix[i] = append(unitMatrix[i], dto[i][j].ToValueObject())
		}
	}
	return valueobject.NewUnitBlock(unitMatrix)
}
