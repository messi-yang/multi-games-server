package main

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/gameclientcommunicator"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/application/applicationservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/infrastructure/memoryrepository"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/config"
)

func main() {
	gameRoomRepositoryMemory := memoryrepository.GetGameRoomRepositoryMemory()
	size := config.GetConfig().GetGameMapSize()
	gameRoomApplicationService := applicationservice.NewGameRoomApplicationService(
		applicationservice.GameRoomApplicationServiceConfiguration{GameRoomRepository: gameRoomRepositoryMemory},
	)
	newGameRoomId, err := gameRoomApplicationService.CreateRoom(size, size)
	if err != nil {
		panic(err.Error())
	}

	config.GetConfig().SetGameId(newGameRoomId)

	gamecomputer.Start()
	gameclientcommunicator.Start()
}
