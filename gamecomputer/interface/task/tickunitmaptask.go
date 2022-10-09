package task

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/application/applicationservice"
)

type TickUnitMapTaskConfiguration struct {
	GameRoomApplicationService applicationservice.GameRoomApplicationService
}

func NewTickUnitMapTask(configuration TickUnitMapTaskConfiguration) {
	go func() {
		gameTicker := time.NewTicker(time.Millisecond * 1000)
		defer gameTicker.Stop()
		unitMapTickerStop := make(chan bool)

		for {
			select {
			case <-gameTicker.C:
				configuration.GameRoomApplicationService.TcikUnitMapInAllGames()
			case <-unitMapTickerStop:
				gameTicker.Stop()
				gameTicker = nil
			}
		}
	}()
}
