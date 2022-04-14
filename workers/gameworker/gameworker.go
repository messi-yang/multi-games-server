package gameworker

import (
	"time"

	"github.com/DumDumGeniuss/game-of-liberty-computer/managers/gamemanager"
)

func StartGameWorker() {
	gameProgressTicker := time.NewTicker(time.Millisecond * 20)
	go func() {
		for range gameProgressTicker.C {
			gamemanager.TickGame()
		}
	}()
}
