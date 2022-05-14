package messageservicetopic

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/domain/valueobject"
)

const GameWorkerTickedMessageTopic = "GAME_WORKER_TICKED"

const GameUnitsUpdatedMessageTopic = "GAME_UNITS_UPDATED"

type GameUnitsUpdatedMessageTopicPayloadUnit struct {
	Coordinate valueobject.Coordinate
	Unit       valueobject.GameUnit
}
type GameUnitsUpdatedMessageTopicPayload []GameUnitsUpdatedMessageTopicPayloadUnit

const GamePlayerJoinedMessageTopic = "GAME_PLAYER_JOINED"

const GamePlayerLeftMessageTopic = "GAME_PLAYER_LEFT"
