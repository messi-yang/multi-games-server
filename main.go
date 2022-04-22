package main

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/config"
	"github.com/DumDumGeniuss/game-of-liberty-computer/routers/gamesocketrouter"
	"github.com/DumDumGeniuss/game-of-liberty-computer/workers/gameworker"
	"github.com/gin-gonic/gin"
)

func main() {
	// Setup config
	config.SetupConfig()

	// Initialize our game and start it
	gameworker.Initialize()
	gameworker.Start()

	// Setup routers
	router := gin.Default()
	gamesocketrouter.SetRouter(router)
	router.Run()
}
