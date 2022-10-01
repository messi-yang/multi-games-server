package main

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/application/service/gameroomservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/config"
	"github.com/dum-dum-genius/game-of-liberty-computer/gameclient"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer"
	"github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/infrastructure/memory/gameroommemory"
)

func main() {
	gameRoomMemory := gameroommemory.GetRepository()
	size := config.GetConfig().GetGameMapSize()
	gameRoomService := gameroomservice.NewService(
		gameroomservice.Configuration{GameRoomRepository: gameRoomMemory},
	)
	newGameRoomId, err := gameRoomService.CreateRoom(size, size)
	if err != nil {
		panic(err.Error())
	}

	config.GetConfig().SetGameId(newGameRoomId)

	gamecomputer.StartJobs()
	gameclient.SetWebsocketRouters()
}
