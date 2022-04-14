package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Config struct {
	GAME_SIZE int
}

func SetupConfig() {
	godotenv.Load(".env")

	if os.Getenv("GAME_SIZE") == "" {
		panic("You must set the 'GAME_SIZE'")
	} else {
		Config.GAME_SIZE, _ = strconv.Atoi(os.Getenv("GAME_SIZE"))
	}
}
