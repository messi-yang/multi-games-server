package main

import (
	game_of_liberty "github.com/DumDumGeniuss/game-of-liberty-computer/worker/game-of-liberty"
	"github.com/gin-gonic/gin"
)

func main() {
	game_of_liberty.StartGame()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, game_of_liberty.GetGameField())
	})
	r.Run()
}
