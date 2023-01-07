package unitmapviewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/unitviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
)

type ViewModel [][]unitviewmodel.ViewModel

func New(unitMap commonmodel.UnitMap) ViewModel {
	unitMapViewModel := make(ViewModel, 0)

	for i := 0; i < unitMap.GetMapSize().GetWidth(); i += 1 {
		unitMapViewModel = append(unitMapViewModel, make([]unitviewmodel.ViewModel, 0))
		for j := 0; j < unitMap.GetMapSize().GetHeight(); j += 1 {
			location, _ := commonmodel.NewLocation(i, j)
			unit := unitMap.GetUnit(location)
			unitMapViewModel[i] = append(unitMapViewModel[i], unitviewmodel.New(unit))
		}
	}
	return unitMapViewModel
}

func (dto ViewModel) ToValueObject() (commonmodel.UnitMap, error) {
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
