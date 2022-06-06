package main

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/service/gameroomservice"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/valueobject"
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/config"
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/job"
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/memory/gameroommemory"
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/router"
)

func main() {
	gameRoomMemoryRepository := gameroommemory.GetGameRoomMemoryRepository()
	gameService := gameroomservice.NewGameRoomService(gameRoomMemoryRepository)

	size := config.GetConfig().GetGameMapSize()
	mapSize := valueobject.NewMapSize(size, size)
	gameRoom := gameService.CreateGameRoom(mapSize)
	config.GetConfig().SetGameId(gameRoom.GetGameId())

	job.StartJobs()
	router.SetRouters()
}
