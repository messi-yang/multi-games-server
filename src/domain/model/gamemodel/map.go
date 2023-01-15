package gamemodel

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

func (_map Map) GetSize() commonmodel.Size {
	size, _ := commonmodel.NewSize(len(_map.unitMatrix), len(_map.unitMatrix[0]))
	return size
}

func (_map Map) GetUnitMatrix() [][]commonmodel.Unit {
	return _map.unitMatrix
}
