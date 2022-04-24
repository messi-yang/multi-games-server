package main

import (
	"fmt"

	"github.com/DumDumGeniuss/game-of-liberty-computer/config"
	"github.com/DumDumGeniuss/game-of-liberty-computer/routers/gamesocketrouter"
	"github.com/DumDumGeniuss/game-of-liberty-computer/stores/gamestore"
	"github.com/DumDumGeniuss/game-of-liberty-computer/workers/gameworker"
	"github.com/DumDumGeniuss/game-of-liberty-computer/workers/oldgameworker"
	"github.com/gin-gonic/gin"
)

func main() {
	// Setup config
	config.SetupConfig()

	// Initialize our game and start it
	oldgameworker.Initialize()
	oldgameworker.Start()

	gamestore.Store.InitializeGame()
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
	gameworker.Worker.SetGameStore(gamestore.Store)
	err := gameworker.Worker.Start()
	if err != nil {
		panic(err.Error())
	}

	// Setup routers
	router := gin.Default()
	gamesocketrouter.SetRouter(router)
	router.Run()
}
