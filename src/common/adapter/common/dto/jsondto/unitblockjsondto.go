package jsondto

import "github.com/dum-dum-genius/game-of-liberty-computer/src/domain/domainmodel/commonmodel"

type UnitBlockJsonDto [][]UnitJsonDto

func NewUnitBlockJsonDto(unitBlock commonmodel.UnitBlock) UnitBlockJsonDto {
	unitBlockJsonDto := make(UnitBlockJsonDto, 0)

	for i := 0; i < unitBlock.GetDimension().GetWidth(); i += 1 {
		unitBlockJsonDto = append(unitBlockJsonDto, make([]UnitJsonDto, 0))
		for j := 0; j < unitBlock.GetDimension().GetHeight(); j += 1 {
			coord, _ := commonmodel.NewCoordinate(i, j)
			unit := unitBlock.GetUnit(coord)
			unitBlockJsonDto[i] = append(unitBlockJsonDto[i], NewUnitJsonDto(unit))
		}
	}
	return unitBlockJsonDto
}

func (dto UnitBlockJsonDto) ToValueObject() (commonmodel.UnitBlock, error) {
	unitMatrix := make([][]commonmodel.Unit, 0)

	for i := 0; i < len(dto); i += 1 {
		unitMatrix = append(unitMatrix, make([]commonmodel.Unit, 0))
		for j := 0; j < len(dto[0]); j += 1 {
			unit, err := dto[i][j].ToValueObject()
			if err != nil {
				return commonmodel.UnitBlock{}, err
			}
			unitMatrix[i] = append(unitMatrix[i], unit)
		}
	}
	return commonmodel.NewUnitBlock(unitMatrix), nil
}
