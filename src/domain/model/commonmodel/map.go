package commonmodel

type Map struct {
	unitMatrix [][]Unit
}

func NewMap(unitMatrix [][]Unit) Map {
	return Map{
		unitMatrix: unitMatrix,
	}
}

func (um Map) GetSize() Size {
	size, _ := NewSize(len(um.unitMatrix), len(um.unitMatrix[0]))
	return size
}

func (um Map) GetUnitMatrix() [][]Unit {
	return um.unitMatrix
}

func (um Map) GetUnit(location Location) Unit {
	return (um.unitMatrix)[location.GetX()][location.GetY()]
}

func (um Map) ReplaceUnitAt(location Location, unit Unit) {
	(um.unitMatrix)[location.GetX()][location.GetY()] = unit
}
