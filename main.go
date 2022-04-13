package main

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/models"
	"github.com/DumDumGeniuss/game-of-liberty-computer/routers/gameblockrouter"
	"github.com/DumDumGeniuss/game-of-liberty-computer/workers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	// messages.InitializeMessages()
	models.InitializeModels()

	workers.InitializeWorks()

	router := gin.Default()
	gameblockrouter.SetRouters(router.Group("/game-block"))
	router.Run()
}
