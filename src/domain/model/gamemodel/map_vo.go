package gamemodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
)

type MapVo struct {
	unitMatrix [][]commonmodel.UnitVo
}

func NewMapVo(unitMatrix [][]commonmodel.UnitVo) MapVo {
	return MapVo{
		unitMatrix: unitMatrix,
	}
}

func (_map MapVo) GetSize() commonmodel.SizeVo {
	size, _ := commonmodel.NewSizeVo(len(_map.unitMatrix), len(_map.unitMatrix[0]))
	return size
}

func (_map MapVo) GetUnitMatrix() [][]commonmodel.UnitVo {
	return _map.unitMatrix
}
