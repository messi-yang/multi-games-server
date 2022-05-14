package gamesocketcontroller

import (
	"sync"

	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/valueobject"
)

type session struct {
	gameAreaToWatch *valueobject.Area
	socketLocker    sync.RWMutex
}
