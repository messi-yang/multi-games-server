package intevent

import (
	"fmt"
)

func CreateLiveGameAdminChannel() string {
	return "LIVE_GAME_ADMIN"
}

func CreateLiveGameClientChannel(liveGameId string, playerId string) string {
	return fmt.Sprintf("LIVE_GAME_%s_CLIENT_PLAYER_%s", liveGameId, playerId)
}
