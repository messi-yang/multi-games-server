package entity

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/valueobject"
	"github.com/google/uuid"
)

type Game struct {
	Id         uuid.UUID
	UnitMatrix [][]valueobject.GameUnit
	MapSize    valueobject.MapSize
}

func NewGame() Game {
	id, _ := uuid.NewUUID()
	return Game{
		Id:         id,
		UnitMatrix: make([][]valueobject.GameUnit, 0),
		MapSize:    valueobject.NewMapSize(0, 0),
	}
}
