package main

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/services/messageservice"
	"github.com/DumDumGeniuss/game-of-liberty-computer/application/workers/gameworker"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/service/gameroomservice"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/valueobject"
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/config"
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/memory"
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/router"
	"github.com/gin-gonic/gin"
)

func main() {
	gameRoomMemoryRepository := memory.NewGameRoomMemoryRepository()
	gameService := gameroomservice.NewGameRoomService(gameRoomMemoryRepository)

	size := config.GetConfig().GetGameMapSize()
	mapSize := valueobject.NewMapSize(size, size)
	gameRoom := gameService.CreateGameRoom(mapSize)
	config.GetConfig().SetGameId(gameRoom.GetGameId())

	messageService := messageservice.GetMessageService()
	gameWorker := gameworker.GetGameWorker()
	gameWorker.InjectGameRoomService(gameService)
	gameWorker.InjectMessageService(messageService)
	if err := gameWorker.StartGame(); err != nil {
		panic(err)
	}

	router.SetRouters(gin.Default())
}
