package main

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/daos/gamedao"
	"github.com/DumDumGeniuss/game-of-liberty-computer/models/gamemodel"
	"github.com/DumDumGeniuss/game-of-liberty-computer/routers/gamesocketrouter"
	"github.com/DumDumGeniuss/game-of-liberty-computer/services/gameservice"
	"github.com/DumDumGeniuss/game-of-liberty-computer/workers/gameworker"
	"github.com/DumDumGeniuss/game-of-liberty-computer/workers/oldgameworker"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize our game and start it
	oldgameworker.Initialize()
	oldgameworker.Start()

	gameModel := gamemodel.GetGameModel()
	gameDAO := gamedao.GetGameDAO(gameModel)
	gameService, err := gameservice.CreateGameService(gameDAO)
	if err != nil {
		panic(err.Error())
	}

	gameworker, err := gameworker.CreateGameWorker(gameService)
	if err != nil {
		panic(err.Error())
	}
	gameworker.StartGame()

	// Setup routers
	router := gin.Default()
	gamesocketrouter.SetRouter(router)
	router.Run()
}
