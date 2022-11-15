package uidto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/gamecommonmodel"
)

type UnitBlockUiDto [][]UnitUiDto

func NewUnitBlockUiDto(unitBlock gamecommonmodel.UnitBlock) UnitBlockUiDto {
	unitBlockUiDto := make(UnitBlockUiDto, 0)

	for i := 0; i < unitBlock.GetDimension().GetWidth(); i += 1 {
		unitBlockUiDto = append(unitBlockUiDto, make([]UnitUiDto, 0))
		for j := 0; j < unitBlock.GetDimension().GetHeight(); j += 1 {
			coord, _ := gamecommonmodel.NewCoordinate(i, j)
			unit := unitBlock.GetUnit(coord)
			unitBlockUiDto[i] = append(unitBlockUiDto[i], NewUnitUiDto(unit))
		}
	}
	return unitBlockUiDto
}

func (dto UnitBlockUiDto) ToValueObject() gamecommonmodel.UnitBlock {
	unitMatrix := make([][]gamecommonmodel.Unit, 0)

	for i := 0; i < len(dto); i += 1 {
		unitMatrix = append(unitMatrix, make([]gamecommonmodel.Unit, 0))
		for j := 0; j < len(dto[0]); j += 1 {
			unitMatrix[i] = append(unitMatrix[i], dto[i][j].ToValueObject())
		}
	}
	return gamecommonmodel.NewUnitBlock(unitMatrix)
}
