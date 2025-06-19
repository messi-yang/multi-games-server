package dto

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/commandmodel"
	"github.com/google/uuid"
)

type CommandDto struct {
	Id        uuid.UUID `json:"id"`
	Timestamp int64     `json:"timestamp"`
	Name      string    `json:"name"`
	Payload   any       `json:"payload"`
}

func NewCommandDto(command commandmodel.Command) CommandDto {
	return CommandDto{
		Id:        command.GetId().Uuid(),
		Timestamp: command.GetTimestamp(),
		Name:      command.GetName(),
		Payload:   command.GetPayload(),
	}
}
