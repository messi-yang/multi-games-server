package unitmapviewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/mapunitviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
)

type ViewModel [][]mapunitviewmodel.ViewModel

func New(unitMap commonmodel.UnitMap) ViewModel {
	unitMapViewModel := make(ViewModel, 0)

	for i := 0; i < unitMap.GetMapSize().GetWidth(); i += 1 {
		unitMapViewModel = append(unitMapViewModel, make([]mapunitviewmodel.ViewModel, 0))
		for j := 0; j < unitMap.GetMapSize().GetHeight(); j += 1 {
			location, _ := commonmodel.NewLocation(i, j)
			mapUnit := unitMap.GetMapUnit(location)
			unitMapViewModel[i] = append(unitMapViewModel[i], mapunitviewmodel.New(mapUnit))
		}
	}
	return unitMapViewModel
}

func (dto ViewModel) ToValueObject() (commonmodel.UnitMap, error) {
	mapUnitMatrix := make([][]commonmodel.MapUnit, 0)

	for i := 0; i < len(dto); i += 1 {
		mapUnitMatrix = append(mapUnitMatrix, make([]commonmodel.MapUnit, 0))
		for j := 0; j < len(dto[0]); j += 1 {
			mapUnit, err := dto[i][j].ToValueObject()
			if err != nil {
				return commonmodel.UnitMap{}, err
			}
			mapUnitMatrix[i] = append(mapUnitMatrix[i], mapUnit)
		}
	}
	return commonmodel.NewUnitMap(mapUnitMatrix), nil
}
