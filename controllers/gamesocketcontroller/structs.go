package gamesocketcontroller

import (
	"sync"

	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/service/gameservice"
)

type session struct {
	gameAreaToWatch *gameservice.GameArea
	socketLocker    sync.RWMutex
}
