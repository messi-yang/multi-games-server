package commandmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
)

type Command struct {
	CommandEntity
}

// Interface Implementation Check
var _ domain.Aggregate = (*Command)(nil)

func CreateCommand(
	id CommandId,
	timestamp int64,
	name string,
	payload any,
) Command {
	return Command{
		NewCommandEntity(
			id,
			timestamp,
			name,
			payload,
		),
	}
}

func (command *Command) GetId() CommandId {
	return command.CommandEntity.GetId()
}
