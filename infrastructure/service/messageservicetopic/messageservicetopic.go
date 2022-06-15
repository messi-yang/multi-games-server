package messageservicetopic

import (
	"github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/dto"
)

const GameUnitsUpdatedMessageTopic = "GAME_UNITS_UPDATED"

type GameUnitsUpdatedMessageTopicPayloadUnit struct {
	Coordinate dto.CoordinateDTO
	Unit       dto.GameUnitDTO
}
type GameUnitsUpdatedMessageTopicPayload []GameUnitsUpdatedMessageTopicPayloadUnit
