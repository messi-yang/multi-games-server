package unitblockviewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/interface/viewmodel/unitviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
)

type UnitBlockViewModel [][]unitviewmodel.UnitViewModel

func New(unitBlock commonmodel.UnitBlock) UnitBlockViewModel {
	unitBlockViewModel := make(UnitBlockViewModel, 0)

	for i := 0; i < unitBlock.GetDimension().GetWidth(); i += 1 {
		unitBlockViewModel = append(unitBlockViewModel, make([]unitviewmodel.UnitViewModel, 0))
		for j := 0; j < unitBlock.GetDimension().GetHeight(); j += 1 {
			coord, _ := commonmodel.NewCoordinate(i, j)
			unit := unitBlock.GetUnit(coord)
			unitBlockViewModel[i] = append(unitBlockViewModel[i], unitviewmodel.New(unit))
		}
	}
	return unitBlockViewModel
}

func (dto UnitBlockViewModel) ToValueObject() (commonmodel.UnitBlock, error) {
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
