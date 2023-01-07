package commonmodel

type UnitMap struct {
	unitMatrix [][]Unit
}

func NewUnitMap(unitMatrix [][]Unit) UnitMap {
	return UnitMap{
		unitMatrix: unitMatrix,
	}
}

func (um UnitMap) GetMapSize() MapSize {
	unitMapSize, _ := NewMapSize(len(um.unitMatrix), len(um.unitMatrix[0]))
	return unitMapSize
}

func (um UnitMap) GetUnitMatrix() [][]Unit {
	return um.unitMatrix
}

func (um UnitMap) GetUnit(location Location) Unit {
	return (um.unitMatrix)[location.GetX()][location.GetY()]
}

func (um UnitMap) ReplaceUnitAt(location Location, unit Unit) {
	(um.unitMatrix)[location.GetX()][location.GetY()] = unit
}
