package intevent

import (
	"fmt"
)

func CreateGameChannel(gameId string) string {
	return fmt.Sprintf("GAME_%s", gameId)
}
