package main

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/application/service/gameroomservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/config"
	"github.com/dum-dum-genius/game-of-liberty-computer/http"
	"github.com/dum-dum-genius/game-of-liberty-computer/infrastructure/memory/gameroommemory"
	"github.com/dum-dum-genius/game-of-liberty-computer/job"
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

	job.StartJobs()
	http.SetWebsocketRouters()
}
