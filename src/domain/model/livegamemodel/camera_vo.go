package livegamemodel

import "github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"

type CameraVo struct {
	center commonmodel.LocationVo
}

func NewCameraVo(center commonmodel.LocationVo) CameraVo {
	return CameraVo{
		center: center,
	}
}

func (camera CameraVo) GetCenter() commonmodel.LocationVo {
	return camera.center
}

func (camera CameraVo) GetViewBoundInMap(mapSize commonmodel.SizeVo) BoundVo {
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

	from := commonmodel.NewLocationVo(fromX, fromY)
	to := commonmodel.NewLocationVo(toX, toY)
	bound, _ := NewBoundVo(from, to)

	return bound
}
