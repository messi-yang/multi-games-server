package commandmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/gamemodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/playermodel"
)

type Command struct {
	CommandEntity
}

// Interface Implementation Check
var _ domain.Aggregate = (*Command)(nil)

func CreateCommand(
	id CommandId,
	gameId gamemodel.GameId,
	playerId playermodel.PlayerId,
	timestamp int64,
	name string,
	payload map[string]interface{},
) Command {
	return Command{
		NewCommandEntity(
			id,
			gameId,
			playerId,
			timestamp,
			name,
			payload,
		),
	}
}

func LoadCommand(
	id CommandId,
	gameId gamemodel.GameId,
	playerId playermodel.PlayerId,
	timestamp int64,
	name string,
	payload map[string]interface{},
) Command {
	return Command{
		NewCommandEntity(
			id,
			gameId,
			playerId,
			timestamp,
			name,
			payload,
		),
	}
}

func (command *Command) GetId() CommandId {
	return command.CommandEntity.GetId()
}
