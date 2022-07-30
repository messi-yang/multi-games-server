package main

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/application/usecase/creategameroomusecase"
	"github.com/dum-dum-genius/game-of-liberty-computer/infrastructure/config"
	"github.com/dum-dum-genius/game-of-liberty-computer/infrastructure/job"
	"github.com/dum-dum-genius/game-of-liberty-computer/infrastructure/memory/gameroommemory"
	"github.com/dum-dum-genius/game-of-liberty-computer/infrastructure/router"
)

func main() {
	gameRoomMemory := gameroommemory.GetGameRoomMemory()
	gameRoom := creategameroomusecase.New(gameRoomMemory).Execute()

	config.GetConfig().SetGameId(gameRoom.GetGameId())

	job.StartJobs()
	router.SetRouters()
}
