package game

import (
	"encoding/json"
	"os"
	"strconv"
	"time"

	messages_game "github.com/DumDumGeniuss/game-of-liberty-computer/messages/game"
)

var gameUnitsGenerationTicker *time.Ticker

func InitializeGameWorker() {
	scale, _ := strconv.Atoi(os.Getenv("GAME_MAP_SCALE"))
	blockSize := 30
	game := newGame(scale, blockSize)

	gameUnitsGenerationTicker = time.NewTicker(time.Millisecond * 1000)
	go func() {
		for range gameUnitsGenerationTicker.C {
			game.GenerateNextUnits()

			for rowIdx := 0; rowIdx < scale; rowIdx += 1 {
				for colIdx := 0; colIdx < scale; colIdx += 1 {
					area := getBlockArea(rowIdx, colIdx, blockSize)
					units, _ := game.GetUnitsInArea(&area)
					bytes, _ := json.Marshal(*units)

					go messages_game.WriteBlockUpdateMessage(rowIdx, colIdx, bytes)
				}
			}
		}
	}()
}
