package gamemapviewmodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/gamemapunitviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
)

type ViewModel [][]gamemapunitviewmodel.ViewModel

func New(gameMap commonmodel.GameMap) ViewModel {
	gameMapViewModel := make(ViewModel, 0)

	for i := 0; i < gameMap.GetMapSize().GetWidth(); i += 1 {
		gameMapViewModel = append(gameMapViewModel, make([]gamemapunitviewmodel.ViewModel, 0))
		for j := 0; j < gameMap.GetMapSize().GetHeight(); j += 1 {
			location, _ := commonmodel.NewLocation(i, j)
			gameMapUnit := gameMap.GetGameMapUnit(location)
			gameMapViewModel[i] = append(gameMapViewModel[i], gamemapunitviewmodel.New(gameMapUnit))
		}
	}
	return gameMapViewModel
}

func (dto ViewModel) ToValueObject() (commonmodel.GameMap, error) {
	gameMapUnitMatrix := make([][]commonmodel.GameMapUnit, 0)

	for i := 0; i < len(dto); i += 1 {
		gameMapUnitMatrix = append(gameMapUnitMatrix, make([]commonmodel.GameMapUnit, 0))
		for j := 0; j < len(dto[0]); j += 1 {
			gameMapUnit, err := dto[i][j].ToValueObject()
			if err != nil {
				return commonmodel.GameMap{}, err
			}
			gameMapUnitMatrix[i] = append(gameMapUnitMatrix[i], gameMapUnit)
		}
	}
	return commonmodel.NewGameMap(gameMapUnitMatrix), nil
}
