package main

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/config"
	"github.com/DumDumGeniuss/game-of-liberty-computer/managers/gamemanager"
	"github.com/DumDumGeniuss/game-of-liberty-computer/routers/gameblockrouter"
	"github.com/DumDumGeniuss/game-of-liberty-computer/workers/gameworker"
	"github.com/gin-gonic/gin"
)

func main() {
	// Setup config
	config.SetupConfig()

	// Initialize our game
	gamemanager.InitializeGameManager()

	// Kick off game in background process
	gameworker.StartGameWorker()

	// Setup routers
	router := gin.Default()
	gameblockrouter.SetRouter(router.Group("/game-block"))
	router.Run()
}
