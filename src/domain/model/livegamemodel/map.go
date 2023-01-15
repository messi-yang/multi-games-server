package livegamemodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/tool"
)

type Map struct {
	unitMatrix [][]commonmodel.Unit
}

func NewMap(unitMatrix [][]commonmodel.Unit) Map {
	return Map{
		unitMatrix: unitMatrix,
	}
}

func (_map Map) GetSize() commonmodel.Size {
	size, _ := commonmodel.NewSize(len(_map.unitMatrix), len(_map.unitMatrix[0]))
	return size
}

func (_map Map) GetUnitMatrix() [][]commonmodel.Unit {
	return _map.unitMatrix
}

func (_map Map) GetMapInBound(bound Bound) Map {
	offsetX := bound.GetFrom().GetX()
	offsetY := bound.GetFrom().GetY()
	boundWidth := bound.GetWidth()
	boundHeight := bound.GetHeight()
	unitMatrix, _ := tool.RangeMatrix(boundWidth, boundHeight, func(x int, y int) (commonmodel.Unit, error) {
		location, _ := commonmodel.NewLocation(x+offsetX, y+offsetY)
		return _map.GetUnit(location), nil
	})
	return NewMap(unitMatrix)
}

func (_map Map) GetUnit(location commonmodel.Location) commonmodel.Unit {
	return (_map.unitMatrix)[location.GetX()][location.GetY()]
}

func (_map Map) ReplaceUnitAt(location commonmodel.Location, unit commonmodel.Unit) {
	(_map.unitMatrix)[location.GetX()][location.GetY()] = unit
}
