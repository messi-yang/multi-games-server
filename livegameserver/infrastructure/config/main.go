package config

import (
	"os"
	"strconv"

	"github.com/google/uuid"
)

type Config interface {
	GetLiveGameDimension() int
}

type configImpl struct {
	GAME_MAP_SIZE int
	GAME_ID       uuid.UUID
}

var config Config = nil

func GetConfig() Config {
	if config == nil {
		newConfig := &configImpl{}

		if os.Getenv("GAME_MAP_SIZE") == "" {
			panic("You must set the 'GAME_MAP_SIZE'")
		} else {
			newConfig.GAME_MAP_SIZE, _ = strconv.Atoi(os.Getenv("GAME_MAP_SIZE"))
		}

		config = newConfig
		return newConfig
	} else {
		return config
	}
}

func (ci *configImpl) GetLiveGameDimension() int {
	return ci.GAME_MAP_SIZE
}
