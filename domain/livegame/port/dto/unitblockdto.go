package dto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/gamecommonmodel"
)

type UnitBlockDto [][]UnitDto

func NewUnitBlockDto(unitBlock gamecommonmodel.UnitBlock) UnitBlockDto {
	unitBlockDto := make(UnitBlockDto, 0)

	for i := 0; i < unitBlock.GetDimension().GetWidth(); i += 1 {
		unitBlockDto = append(unitBlockDto, make([]UnitDto, 0))
		for j := 0; j < unitBlock.GetDimension().GetHeight(); j += 1 {
			coord, _ := gamecommonmodel.NewCoordinate(i, j)
			unit := unitBlock.GetUnit(coord)
			unitBlockDto[i] = append(unitBlockDto[i], NewUnitDto(unit))
		}
	}
	return unitBlockDto
}

func (dto UnitBlockDto) ToValueObject() gamecommonmodel.UnitBlock {
	unitMatrix := make([][]gamecommonmodel.Unit, 0)

	for i := 0; i < len(dto); i += 1 {
		unitMatrix = append(unitMatrix, make([]gamecommonmodel.Unit, 0))
		for j := 0; j < len(dto[0]); j += 1 {
			unitMatrix[i] = append(unitMatrix[i], dto[i][j].ToValueObject())
		}
	}
	return gamecommonmodel.NewUnitBlock(unitMatrix)
}
