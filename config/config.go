package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config interface {
	GetGameSize() int
}

type configImpl struct {
	GAME_SIZE int
}

var config Config = nil

func GetConfig() Config {
	if config == nil {
		newConfig := &configImpl{}
		godotenv.Load(".env")

		if os.Getenv("GAME_SIZE") == "" {
			panic("You must set the 'GAME_SIZE'")
		} else {
			newConfig.GAME_SIZE, _ = strconv.Atoi(os.Getenv("GAME_SIZE"))
		}

		config = newConfig
		return newConfig
	} else {
		return config
	}
}

func (ci *configImpl) GetGameSize() int {
	return ci.GAME_SIZE
}
