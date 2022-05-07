package entity

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/valueobject"
	"github.com/google/uuid"
)

type Game struct {
	Id    uuid.UUID
	Units [][]valueobject.GameUnit
	Size  valueobject.GameSize
}
