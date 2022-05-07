package gameservice

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/valueobject"
	"github.com/DumDumGeniuss/ggol"
)

func convertGgolSizeToGameSize(ggolSize *ggol.Size) *valueobject.GameSize {
	return &valueobject.GameSize{
		Width:  ggolSize.Width,
		Height: ggolSize.Height,
	}
}

func convertGameSizeToGgolSize(gameSize *valueobject.GameSize) *ggol.Size {
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
