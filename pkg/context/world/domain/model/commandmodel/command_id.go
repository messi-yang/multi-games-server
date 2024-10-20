package commandmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/google/uuid"
)

type CommandId struct {
	id uuid.UUID
}

// Interface Implementation Check
var _ domain.ValueObject[CommandId] = (*CommandId)(nil)

func NewCommandId(id uuid.UUID) CommandId {
	return CommandId{
		id: id,
	}
}

func (commandId CommandId) IsEqual(otherCommandId CommandId) bool {
	return commandId.id == otherCommandId.id
}

func (commandId CommandId) Uuid() uuid.UUID {
	return commandId.id
}
