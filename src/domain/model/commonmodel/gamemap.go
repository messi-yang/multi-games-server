package commonmodel

type GameMap struct {
	mapUnitMatrix [][]MapUnit
}

func NewGameMap(mapUnitMatrix [][]MapUnit) GameMap {
	return GameMap{
		mapUnitMatrix: mapUnitMatrix,
	}
}

func (um GameMap) GetMapSize() MapSize {
	gameMapSize, _ := NewMapSize(len(um.mapUnitMatrix), len(um.mapUnitMatrix[0]))
	return gameMapSize
}

func (um GameMap) GetMapUnitMatrix() [][]MapUnit {
	return um.mapUnitMatrix
}

func (um GameMap) GetMapUnit(location Location) MapUnit {
	return (um.mapUnitMatrix)[location.GetX()][location.GetY()]
}

func (um GameMap) ReplaceMapUnitAt(location Location, mapUnit MapUnit) {
	(um.mapUnitMatrix)[location.GetX()][location.GetY()] = mapUnit
}
