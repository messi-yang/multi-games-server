package config

import (
	"os"
	"strconv"

	"github.com/google/uuid"
)

type Config interface {
	GetGameMapSize() int
	GetGameId() uuid.UUID
	SetGameId(uuid.UUID)
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

func (ci *configImpl) GetGameMapSize() int {
	return ci.GAME_MAP_SIZE
}

func (ci *configImpl) GetGameId() uuid.UUID {
	return ci.GAME_ID
}

func (ci *configImpl) SetGameId(id uuid.UUID) {
	ci.GAME_ID = id
}
