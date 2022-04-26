package main

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/daos/gamedao"
	"github.com/DumDumGeniuss/game-of-liberty-computer/models/gamemodel"
	"github.com/DumDumGeniuss/game-of-liberty-computer/routers/gamesocketrouter"
	"github.com/DumDumGeniuss/game-of-liberty-computer/services/gameservice"
	"github.com/DumDumGeniuss/game-of-liberty-computer/services/messageservice"
	"github.com/DumDumGeniuss/game-of-liberty-computer/workers/gameworker"
	"github.com/gin-gonic/gin"
)

func main() {
	gameModel := gamemodel.GetGameModel()

	gameDAO := gamedao.GetGameDAO()
	gameDAO.InjectGameModel(gameModel)

	gameService := gameservice.GetGameService()
	gameService.InjectGameDAO(gameDAO)
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
