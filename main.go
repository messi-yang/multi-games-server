package main

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/service/gameroomservice"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/game/valueobject"
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/config"
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/job/gameroomjob"
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/memory/gameroommemory"
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/router"
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/service/messageservice"
	"github.com/gin-gonic/gin"
)

func main() {
	gameRoomMemoryRepository := gameroommemory.NewGameRoomMemoryRepository()
	gameService := gameroomservice.NewGameRoomService(gameRoomMemoryRepository)

	size := config.GetConfig().GetGameMapSize()
	mapSize := valueobject.NewMapSize(size, size)
	gameRoom := gameService.CreateGameRoom(mapSize)
	config.GetConfig().SetGameId(gameRoom.GetGameId())

	messageService := messageservice.GetMessageService()
	gameRoomJob := gameroomjob.NewGameRoomJob(gameService, messageService)
	if err := gameRoomJob.StartGame(gameRoom.GetGameId()); err != nil {
		panic(err)
	}

	router.SetRouters(gin.Default())
}
