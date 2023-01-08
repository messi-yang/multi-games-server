package viewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
)

type UnitMapViewModel [][]UnitViewModel

func NewUnitMapViewModel(unitMap commonmodel.UnitMap) UnitMapViewModel {
	unitMapViewModel := make(UnitMapViewModel, 0)

	for i := 0; i < unitMap.GetMapSize().GetWidth(); i += 1 {
		unitMapViewModel = append(unitMapViewModel, make([]UnitViewModel, 0))
		for j := 0; j < unitMap.GetMapSize().GetHeight(); j += 1 {
			location, _ := commonmodel.NewLocation(i, j)
			unit := unitMap.GetUnit(location)
			unitMapViewModel[i] = append(unitMapViewModel[i], NewUnitViewModel(unit))
		}
	}
	return unitMapViewModel
}

func (dto UnitMapViewModel) ToValueObject() (commonmodel.UnitMap, error) {
	unitMatrix := make([][]commonmodel.Unit, 0)

	for i := 0; i < len(dto); i += 1 {
		unitMatrix = append(unitMatrix, make([]commonmodel.Unit, 0))
		for j := 0; j < len(dto[0]); j += 1 {
			unit, err := dto[i][j].ToValueObject()
			if err != nil {
				return commonmodel.UnitMap{}, err
			}
			unitMatrix[i] = append(unitMatrix[i], unit)
		}
	}
	return commonmodel.NewUnitMap(unitMatrix), nil
}
