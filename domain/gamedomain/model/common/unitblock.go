package common

type UnitBlock struct {
	unitMatrix [][]Unit
}

func NewUnitBlock(unitMatrix [][]Unit) UnitBlock {
	return UnitBlock{
		unitMatrix: unitMatrix,
	}
}

func (um UnitBlock) GetDimension() Dimension {
	gameDimension, _ := NewDimension(len(um.unitMatrix), len(um.unitMatrix[0]))
	return gameDimension
}

func (um UnitBlock) GetUnitMatrix() [][]Unit {
	return um.unitMatrix
}

func (um UnitBlock) GetUnit(coord Coordinate) Unit {
	return (um.unitMatrix)[coord.GetX()][coord.GetY()]
}

func (um UnitBlock) SetUnit(coord Coordinate, unit Unit) {
	(um.unitMatrix)[coord.GetX()][coord.GetY()] = unit
}
