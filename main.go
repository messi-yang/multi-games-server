package main

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/routers"
	"github.com/DumDumGeniuss/game-of-liberty-computer/workers"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	workers.InitializeWorks()
	routers.InitializeRouters()
}
