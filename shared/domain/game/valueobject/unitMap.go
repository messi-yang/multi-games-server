package valueobject

type UnitMap struct {
	unitMatrix [][]Unit
}

func NewUnitMap(unitMatrix [][]Unit) UnitMap {
	return UnitMap{
		unitMatrix: unitMatrix,
	}
}

func (um UnitMap) ToValueObjectMatrix() [][]Unit {
	return um.unitMatrix
}

func (um UnitMap) GetMapSize() MapSize {
	gameMapSize, _ := NewMapSize(len(um.unitMatrix), len(um.unitMatrix[0]))
	return gameMapSize
}

func (um UnitMap) GetUnit(coord Coordinate) Unit {
	return (um.unitMatrix)[coord.GetX()][coord.GetY()]
}

func (um UnitMap) SetUnit(coord Coordinate, unit Unit) {
	(um.unitMatrix)[coord.GetX()][coord.GetY()] = unit
}
