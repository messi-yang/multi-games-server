package gamemapviewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/mapunitviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
)

type ViewModel [][]mapunitviewmodel.ViewModel

func New(gameMap commonmodel.GameMap) ViewModel {
	gameMapViewModel := make(ViewModel, 0)

	for i := 0; i < gameMap.GetMapSize().GetWidth(); i += 1 {
		gameMapViewModel = append(gameMapViewModel, make([]mapunitviewmodel.ViewModel, 0))
		for j := 0; j < gameMap.GetMapSize().GetHeight(); j += 1 {
			location, _ := commonmodel.NewLocation(i, j)
			mapUnit := gameMap.GetMapUnit(location)
			gameMapViewModel[i] = append(gameMapViewModel[i], mapunitviewmodel.New(mapUnit))
		}
	}
	return gameMapViewModel
}

func (dto ViewModel) ToValueObject() (commonmodel.GameMap, error) {
	mapUnitMatrix := make([][]commonmodel.MapUnit, 0)

	for i := 0; i < len(dto); i += 1 {
		mapUnitMatrix = append(mapUnitMatrix, make([]commonmodel.MapUnit, 0))
		for j := 0; j < len(dto[0]); j += 1 {
			mapUnit, err := dto[i][j].ToValueObject()
			if err != nil {
				return commonmodel.GameMap{}, err
			}
			mapUnitMatrix[i] = append(mapUnitMatrix[i], mapUnit)
		}
	}
	return commonmodel.NewGameMap(mapUnitMatrix), nil
}
