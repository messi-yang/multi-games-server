package gameservice

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/valueobject"
	"github.com/DumDumGeniuss/ggol"
)

func convertGgolSizeToMapSize(ggolSize *ggol.Size) *valueobject.MapSize {
	return &valueobject.MapSize{
		Width:  ggolSize.Width,
		Height: ggolSize.Height,
	}
}

func convertMapSizeToGgolSize(gameSize *valueobject.MapSize) *ggol.Size {
	return &ggol.Size{
		Width:  gameSize.Width,
		Height: gameSize.Height,
	}
}

func convertGameCoordinateToGgolCoordinate(gameCoordinate *GameCoordinate) *ggol.Coordinate {
	return &ggol.Coordinate{
		X: gameCoordinate.X,
		Y: gameCoordinate.Y,
	}
}

func convertGameAreaToGgolArea(gameArea *GameArea) *ggol.Area {
	return &ggol.Area{
		From: ggol.Coordinate(gameArea.From),
		To:   ggol.Coordinate(gameArea.To),
	}
}
