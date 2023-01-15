package livegamemodel

import "github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"

type Camera struct {
	center commonmodel.Location
}

func NewCamera(center commonmodel.Location) Camera {
	return Camera{
		center: center,
	}
}

func (camera Camera) GetCenter() commonmodel.Location {
	return camera.center
}

func (camera Camera) GetViwBoundInMap(mapSize commonmodel.Size) commonmodel.Bound {
	fromX := camera.GetCenter().GetX() - 25
	toX := camera.GetCenter().GetX() + 25
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
	bound, _ := commonmodel.NewBound(from, to)

	return bound
}
