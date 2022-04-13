package game

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/DumDumGeniuss/game-of-liberty-computer/models/gameblockmodel"
)

var gameUnitsGenerationTicker *time.Ticker

func InitializeGameWorker() {
	scale, _ := strconv.Atoi(os.Getenv("GAME_MAP_SCALE"))
	blockSize, _ := strconv.Atoi(os.Getenv("GAME_BLOCK_SIZE"))
	game := newGame(scale, blockSize)

	gameUnitsGenerationTicker = time.NewTicker(time.Millisecond * 1000)
	go func() {
		for range gameUnitsGenerationTicker.C {
			game.GenerateNextUnits()

			for rowIdx := 0; rowIdx < scale; rowIdx += 1 {
				for colIdx := 0; colIdx < scale; colIdx += 1 {
					area := getBlockArea(rowIdx, colIdx, blockSize)
					units, _ := game.GetUnitsInArea(&area)
					go gameblockmodel.CreateOrUpdateGameBlocks(rowIdx, colIdx, units)
					// bytes, _ := json.Marshal(*units)

					// fmt.Print(bytes)
					// messages_game.WriteBlockUpdateMessage(x, y, bytes)
				}
			}

			fmt.Println("Done inserting")
		}
	}()
	// go messages_game.WatchBlockUpdateMessages(0, 0)
	// go messages_game.WatchBlockUpdateMessages(1, 1)
	// go messages_game.WatchBlockUpdateMessages(2, 2)
	// go messages_game.WatchBlockUpdateMessages(3, 3)
	// go messages_game.WatchBlockUpdateMessages(9, 9)

	gameblockmodel.GetGameBlock(0, 0)
	// fmt.Print(*gameblock)
}
