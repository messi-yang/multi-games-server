package livegamemodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/library/tool"
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
	size, _ := commonmodel.NewSizeVo(len(map_.unitMatrix), len(map_.unitMatrix[0]))
	return size
}

func (map_ MapVo) GetUnitMatrix() [][]commonmodel.UnitVo {
	return map_.unitMatrix
}

func (map_ MapVo) GetUnit(location commonmodel.LocationVo) commonmodel.UnitVo {
	return (map_.unitMatrix)[location.GetX()][location.GetY()]
}

func (map_ MapVo) UpdateUnit(location commonmodel.LocationVo, unit commonmodel.UnitVo) {
	(map_.unitMatrix)[location.GetX()][location.GetY()] = unit
}

func (map_ MapVo) GetViewInBound(bound BoundVo) ViewVo {
	offsetX := bound.GetFrom().GetX()
	offsetY := bound.GetFrom().GetY()
	boundWidth := bound.GetWidth()
	boundHeight := bound.GetHeight()
	unitMatrix, _ := tool.RangeMatrix(boundWidth, boundHeight, func(x int, y int) (commonmodel.UnitVo, error) {
		location := commonmodel.NewLocationVo(x+offsetX, y+offsetY)
		return map_.GetUnit(location), nil
	})
	return NewViewVo(NewMapVo(unitMatrix), bound)
}
