package main

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/usecase/creategameroomusecase"
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/config"
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/job"
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/memory/gameroommemory"
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/router"
)

func main() {
	gameRoomMemory := gameroommemory.GetGameRoomMemory()
	createGameUseCase := creategameroomusecase.NewUseCase(gameRoomMemory)
	gameRoom := createGameUseCase.Execute()

	config.GetConfig().SetGameId(gameRoom.GetGameId())

	job.StartJobs()
	router.SetRouters()
}
