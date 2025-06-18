package commandmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
)

type CommandEntity struct {
	id        CommandId
	timestamp int64
	name      string
	payload   any
}

// Interface Implementation Check
var _ domain.Entity[CommandId] = (*CommandEntity)(nil)

func NewCommandEntity(
	id CommandId,
	timestamp int64,
	name string,
	payload any,
) CommandEntity {
	return CommandEntity{
		id:        id,
		timestamp: timestamp,
		name:      name,
		payload:   payload,
	}
}

func LoadCommandEntity(
	id CommandId,
	timestamp int64,
	name string,
	payload any,
) CommandEntity {
	return CommandEntity{
		id:        id,
		timestamp: timestamp,
		name:      name,
		payload:   payload,
	}
}

func (command *CommandEntity) GetId() CommandId {
	return command.id
}

func (command *CommandEntity) GetTimestamp() int64 {
	return command.timestamp
}

func (command *CommandEntity) GetCommandName() string {
	return command.name
}

func (command *CommandEntity) GetPayload() any {
	return command.name
}
