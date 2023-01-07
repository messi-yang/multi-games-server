package commonmodel

type UnitMap struct {
	mapUnitMatrix [][]MapUnit
}

func NewUnitMap(mapUnitMatrix [][]MapUnit) UnitMap {
	return UnitMap{
		mapUnitMatrix: mapUnitMatrix,
	}
}

func (um UnitMap) GetMapSize() MapSize {
	unitMapSize, _ := NewMapSize(len(um.mapUnitMatrix), len(um.mapUnitMatrix[0]))
	return unitMapSize
}

func (um UnitMap) GetMapUnitMatrix() [][]MapUnit {
	return um.mapUnitMatrix
}

func (um UnitMap) GetMapUnit(location Location) MapUnit {
	return (um.mapUnitMatrix)[location.GetX()][location.GetY()]
}

func (um UnitMap) ReplaceMapUnitAt(location Location, mapUnit MapUnit) {
	(um.mapUnitMatrix)[location.GetX()][location.GetY()] = mapUnit
}
