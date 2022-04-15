package main

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/config"
	"github.com/DumDumGeniuss/game-of-liberty-computer/entities/gameentity"
	"github.com/DumDumGeniuss/game-of-liberty-computer/routers/gamerouter"
	"github.com/DumDumGeniuss/game-of-liberty-computer/workers/gameworker"
	"github.com/gin-gonic/gin"
)

func main() {
	// Setup config
	config.SetupConfig()

	// Initialize our game
	gameentity.InitializeGame()

	// Kick off game in background process
	gameworker.StartGameWorker()

	// Setup routers
	router := gin.Default()
	gamerouter.SetRouter(router.Group("/game-block"))
	router.Run()
}
