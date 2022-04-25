package main

import (
	"fmt"

	"github.com/DumDumGeniuss/game-of-liberty-computer/daos/gamedao"
	"github.com/DumDumGeniuss/game-of-liberty-computer/models/gamemodel"
	"github.com/DumDumGeniuss/game-of-liberty-computer/routers/gamesocketrouter"
	"github.com/DumDumGeniuss/game-of-liberty-computer/stores/gamestore"
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
	gameStore := gamestore.GetGameStore(gameDAO)

	fmt.Println(gameStore.GetGameUnitsInArea(&gamestore.GameArea{
		From: gamestore.GameCoordinate{
			X: 0,
			Y: 0,
		},
		To: gamestore.GameCoordinate{
			X: 2,
			Y: 2,
		},
	}))
	fmt.Println(gameStore.GetGameSize())
	gameworker.Worker.SetGameStore(gameStore)
	err := gameworker.Worker.Start()
	if err != nil {
		panic(err.Error())
	}

	// Setup routers
	router := gin.Default()
	gamesocketrouter.SetRouter(router)
	router.Run()
}
