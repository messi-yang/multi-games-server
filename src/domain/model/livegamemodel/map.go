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

func (map_ Map) GetViewOfCamera(camera Camera) View {
	bound := map_.GetViwBoundOfCamera(camera)
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

func (map_ Map) GetViwBoundOfCamera(camera Camera) Bound {
	fromX := camera.GetCenter().GetX() - 25
	toX := camera.GetCenter().GetX() + 25
	mapSize := map_.GetSize()
	mapWidth := mapSize.GetWidth()
	if fromX < 0 {
		toX -= fromX
		fromX = 0
	} else if toX > mapWidth-1 {
		fromX -= toX - mapWidth - 1
		toX = mapWidth - 1
	}

	fromY := camera.GetCenter().GetY() - 25
	toY := camera.GetCenter().GetY() + 25
	mapHeight := mapSize.GetHeight()
	if fromY < 0 {
		toY -= fromY
		fromY = 0
	} else if toY > mapHeight-1 {
		fromY -= toY - mapHeight - 1
		toY = mapHeight - 1
	}

	from, _ := commonmodel.NewLocation(fromX, fromY)
	to, _ := commonmodel.NewLocation(toX, toY)
	bound, _ := NewBound(from, to)

	return bound
}

func (map_ Map) GetUnit(location commonmodel.Location) commonmodel.Unit {
	return (map_.unitMatrix)[location.GetX()][location.GetY()]
}

func (map_ Map) ReplaceUnitAt(location commonmodel.Location, unit commonmodel.Unit) {
	(map_.unitMatrix)[location.GetX()][location.GetY()] = unit
}
