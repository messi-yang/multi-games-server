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

func (map_ MapVo) GetSize() commonmodel.SizeVo {
	if len(map_.unitMatrix) == 0 {
		size, _ := commonmodel.NewSizeVo(0, 0)
		return size
	}
	size, _ := commonmodel.NewSizeVo(len(map_.unitMatrix), len(map_.unitMatrix[0]))
	return size
}

func (map_ MapVo) GetUnitMatrix() [][]commonmodel.UnitVo {
	return map_.unitMatrix
}
