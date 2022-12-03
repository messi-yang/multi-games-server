package jsondto

import (
	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
)

type UnitBlockJsonDto [][]UnitJsonDto

func NewUnitBlockJsonDto(unitBlock gamecommonmodel.UnitBlock) UnitBlockJsonDto {
	unitBlockJsonDto := make(UnitBlockJsonDto, 0)

	for i := 0; i < unitBlock.GetDimension().GetWidth(); i += 1 {
		unitBlockJsonDto = append(unitBlockJsonDto, make([]UnitJsonDto, 0))
		for j := 0; j < unitBlock.GetDimension().GetHeight(); j += 1 {
			coord, _ := gamecommonmodel.NewCoordinate(i, j)
			unit := unitBlock.GetUnit(coord)
			unitBlockJsonDto[i] = append(unitBlockJsonDto[i], NewUnitJsonDto(unit))
		}
	}
	return unitBlockJsonDto
}

func (dto UnitBlockJsonDto) ToValueObject() gamecommonmodel.UnitBlock {
	unitMatrix := make([][]gamecommonmodel.Unit, 0)

	for i := 0; i < len(dto); i += 1 {
		unitMatrix = append(unitMatrix, make([]gamecommonmodel.Unit, 0))
		for j := 0; j < len(dto[0]); j += 1 {
			unitMatrix[i] = append(unitMatrix[i], dto[i][j].ToValueObject())
		}
	}
	return gamecommonmodel.NewUnitBlock(unitMatrix)
}
