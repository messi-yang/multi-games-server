package commonmodel

type GameMap struct {
	gameMapUnitMatrix [][]GameMapUnit
}

func NewGameMap(gameMapUnitMatrix [][]GameMapUnit) GameMap {
	return GameMap{
		gameMapUnitMatrix: gameMapUnitMatrix,
	}
}

func (um GameMap) GetMapSize() MapSize {
	gameMapSize, _ := NewMapSize(len(um.gameMapUnitMatrix), len(um.gameMapUnitMatrix[0]))
	return gameMapSize
}

func (um GameMap) GetGameMapUnitMatrix() [][]GameMapUnit {
	return um.gameMapUnitMatrix
}

func (um GameMap) GetGameMapUnit(location Location) GameMapUnit {
	return (um.gameMapUnitMatrix)[location.GetX()][location.GetY()]
}

func (um GameMap) ReplaceGameMapUnitAt(location Location, gameMapUnit GameMapUnit) {
	(um.gameMapUnitMatrix)[location.GetX()][location.GetY()] = gameMapUnit
}
