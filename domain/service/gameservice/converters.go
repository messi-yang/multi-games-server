package gameservice

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/valueobject"
	"github.com/DumDumGeniuss/ggol"
)

func convertMapSizeToGgolSize(gameSize *valueobject.MapSize) *ggol.Size {
	return &ggol.Size{
		Width:  gameSize.GetWidth(),
		Height: gameSize.GetHeight(),
	}
}

func convertGameCoordinateToGgolCoordinate(gameCoordinate *valueobject.Coordinate) *ggol.Coordinate {
	return &ggol.Coordinate{
		X: gameCoordinate.GetX(),
		Y: gameCoordinate.GetY(),
	}
}

func convertGameAreaToGgolArea(gameArea *valueobject.Area) *ggol.Area {
	return &ggol.Area{
		From: ggol.Coordinate{X: gameArea.GetFrom().GetX(), Y: gameArea.GetFrom().GetY()},
		To:   ggol.Coordinate{X: gameArea.GetTo().GetX(), Y: gameArea.GetTo().GetY()},
	}
}
