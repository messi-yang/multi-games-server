package commonmodel

import "github.com/dum-dum-genius/game-of-liberty-computer/src/library/tool"

type Map struct {
	unitMatrix [][]Unit
}

func NewMap(unitMatrix [][]Unit) Map {
	return Map{
		unitMatrix: unitMatrix,
	}
}

func (_map Map) GetSize() Size {
	size, _ := NewSize(len(_map.unitMatrix), len(_map.unitMatrix[0]))
	return size
}

func (_map Map) GetUnitMatrix() [][]Unit {
	return _map.unitMatrix
}

func (_map Map) GetMapInBound(bound Bound) Map {
	offsetX := bound.GetFrom().GetX()
	offsetY := bound.GetFrom().GetY()
	boundWidth := bound.GetWidth()
	boundHeight := bound.GetHeight()
	unitMatrix, _ := tool.RangeMatrix(boundWidth, boundHeight, func(x int, y int) (Unit, error) {
		location, _ := NewLocation(x+offsetX, y+offsetY)
		return _map.GetUnit(location), nil
	})
	return NewMap(unitMatrix)
}

func (_map Map) GetUnit(location Location) Unit {
	return (_map.unitMatrix)[location.GetX()][location.GetY()]
}

func (_map Map) ReplaceUnitAt(location Location, unit Unit) {
	(_map.unitMatrix)[location.GetX()][location.GetY()] = unit
}
