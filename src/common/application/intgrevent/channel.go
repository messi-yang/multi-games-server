package intgrevent

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/livegamemodel"
)

func CreateLiveGameAdminChannel() string {
	return "LIVE_GAME_ADMIN"
}

func CreateLiveGameClientChannel(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId) string {
	return fmt.Sprintf("LIVE_GAME_%s_CLIENT_PLAYER_%s", liveGameId.ToString(), playerId.ToString())
}
