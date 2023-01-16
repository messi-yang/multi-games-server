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

func (map_ Map) GetViewWithCamera(camera Camera) View {
	bound := camera.GetViewBoundInMap(map_.GetSize())
	offsetX := bound.GetFrom().GetX()
	offsetY := bound.GetFrom().GetY()
	boundWidth := bound.GetWidth()
	boundHeight := bound.GetHeight()
	unitMatrix, _ := tool.RangeMatrix(boundWidth, boundHeight, func(x int, y int) (commonmodel.Unit, error) {
		location, _ := commonmodel.NewLocation(x+offsetX, y+offsetY)
		return map_.GetUnit(location), nil
	})
	return NewView(NewMap(unitMatrix), bound)
}
