package commandmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
)

// Unit here is only for reading purpose, for writing units,
// please check the unit model of the type you are looking for.
type Command struct {
	CommandEntity
}

// Interface Implementation Check
var _ domain.Aggregate = (*Command)(nil)

func CreateCommand(
	id CommandId,
	timestamp int64,
	name CommandName,
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
