package main

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/application/usecase/creategameusecase"
	"github.com/dum-dum-genius/game-of-liberty-computer/infrastructure/config"
	"github.com/dum-dum-genius/game-of-liberty-computer/infrastructure/job"
	"github.com/dum-dum-genius/game-of-liberty-computer/infrastructure/memory/gameroommemory"
	"github.com/dum-dum-genius/game-of-liberty-computer/infrastructure/router"
)

func main() {
	gameRoomMemory := gameroommemory.GetGameRoomMemory()
	gameId, err := creategameusecase.New(gameRoomMemory).Execute()
	if err != nil {
		panic(err.Error())
	}

	config.GetConfig().SetGameId(gameId)

	job.StartJobs()
	router.SetRouters()
}
