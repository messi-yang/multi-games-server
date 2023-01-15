package livegamemodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
)

type Map struct {
	unitMatrix [][]commonmodel.Unit
}

func NewMap(unitMatrix [][]commonmodel.Unit) Map {
	return Map{
		unitMatrix: unitMatrix,
	}
}

func (map_ Map) GetSize() commonmodel.Size {
	size, _ := commonmodel.NewSize(len(map_.unitMatrix), len(map_.unitMatrix[0]))
	return size
}

func (map_ Map) GetUnitMatrix() [][]commonmodel.Unit {
	return map_.unitMatrix
}

func (map_ Map) GetUnit(location commonmodel.Location) commonmodel.Unit {
	return (map_.unitMatrix)[location.GetX()][location.GetY()]
}

func (map_ Map) ReplaceUnitAt(location commonmodel.Location, unit commonmodel.Unit) {
	(map_.unitMatrix)[location.GetX()][location.GetY()] = unit
}
