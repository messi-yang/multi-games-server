package commonmodel

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

func (um UnitBlock) GetUnit(location Location) Unit {
	return (um.unitMatrix)[location.GetX()][location.GetY()]
}

func (um UnitBlock) ReplaceUnitAt(location Location, unit Unit) {
	(um.unitMatrix)[location.GetX()][location.GetY()] = unit
}
