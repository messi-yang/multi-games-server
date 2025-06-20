package commandmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/gamemodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/playermodel"
)

type CommandEntity struct {
	id        CommandId
	gameId    gamemodel.GameId
	playerId  playermodel.PlayerId
	timestamp int64
	name      string
	payload   map[string]interface{}
}

// Interface Implementation Check
var _ domain.Entity[CommandId] = (*CommandEntity)(nil)

func NewCommandEntity(
	id CommandId,
	gameId gamemodel.GameId,
	playerId playermodel.PlayerId,
	timestamp int64,
	name string,
	payload map[string]interface{},
) CommandEntity {
	return CommandEntity{
		id:        id,
		gameId:    gameId,
		playerId:  playerId,
		timestamp: timestamp,
		name:      name,
		payload:   payload,
	}
}

func LoadCommandEntity(
	id CommandId,
	gameId gamemodel.GameId,
	playerId playermodel.PlayerId,
	timestamp int64,
	name string,
	payload map[string]interface{},
) CommandEntity {
	return CommandEntity{
		id:        id,
		gameId:    gameId,
		playerId:  playerId,
		timestamp: timestamp,
		name:      name,
		payload:   payload,
	}
}

func (command *CommandEntity) GetId() CommandId {
	return command.id
}

func (command *CommandEntity) GetGameId() gamemodel.GameId {
	return command.gameId
}

func (command *CommandEntity) GetPlayerId() playermodel.PlayerId {
	return command.playerId
}

func (command *CommandEntity) GetTimestamp() int64 {
	return command.timestamp
}

func (command *CommandEntity) GetName() string {
	return command.name
}

func (command *CommandEntity) GetPayload() map[string]interface{} {
	return command.payload
}
