package main

import (
	"math/rand"

	"github.com/DumDumGeniuss/game-of-liberty-computer/config"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/aggregate"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/entity"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/service/gameservice"
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/valueobject"
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/memory"
	"github.com/DumDumGeniuss/game-of-liberty-computer/routers/gamesocketrouter"
	"github.com/DumDumGeniuss/game-of-liberty-computer/services/messageservice"
	"github.com/DumDumGeniuss/game-of-liberty-computer/workers/gameworker"
	"github.com/gin-gonic/gin"
)

func main() {
	size := config.GetConfig().GetMapSize()

	gameRoomMemoryRepository := memory.NewGameRoomMemoryRepository()

	mapSize := valueobject.NewMapSize(size, size)
	gameUnitMatrix := make([][]valueobject.GameUnit, size)
	for i := 0; i < size; i += 1 {
		gameUnitMatrix[i] = make([]valueobject.GameUnit, size)
		for j := 0; j < size; j += 1 {
			gameUnitMatrix[i][j] = valueobject.NewGameUnit(rand.Intn(2) == 0, 0)
		}
	}
	game := entity.NewGame()
	game.SetMapSize(mapSize)
	game.SetUnitMatrix(gameUnitMatrix)

	gameRoom := aggregate.NewGameRoom(game)
	gameRoomMemoryRepository.Add(gameRoom)

	gameService := gameservice.NewGameService()
	gameService.SetGameId(game.GetId())
	gameService.InjectGameRoomMemoryRepository(gameRoomMemoryRepository)
	if err := gameService.InitializeGame(); err != nil {
		panic(err)
	}

	messageService := messageservice.GetMessageService()
	gameWorker := gameworker.GetGameWorker()
	gameWorker.InjectGameService(gameService)
	gameWorker.InjectMessageService(messageService)
	if err := gameWorker.StartGame(); err != nil {
		panic(err)
	}

	// Setup routers
	router := gin.Default()
	gamesocketrouter.SetRouter(router)
	router.Run()
}
