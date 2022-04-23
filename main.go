package main

import (
	"fmt"

	"github.com/DumDumGeniuss/game-of-liberty-computer/config"
	"github.com/DumDumGeniuss/game-of-liberty-computer/routers/gamesocketrouter"
	"github.com/DumDumGeniuss/game-of-liberty-computer/stores/gamestore"
	"github.com/DumDumGeniuss/game-of-liberty-computer/workers/gameworker"
	"github.com/gin-gonic/gin"
)

func main() {
	// Setup config
	config.SetupConfig()

	// Initialize our game and start it
	gameworker.Initialize()
	gameworker.Start()

	gamestore.Store.StartGame()
	fmt.Println(gamestore.Store.GetGameUnitsInArea(&gamestore.GameArea{
		From: gamestore.GameCoordinate{
			X: 0,
			Y: 0,
		},
		To: gamestore.GameCoordinate{
			X: 2,
			Y: 2,
		},
	}))
	fmt.Println(gamestore.Store.GetGameFieldSize())

	// Setup routers
	router := gin.Default()
	gamesocketrouter.SetRouter(router)
	router.Run()
}
