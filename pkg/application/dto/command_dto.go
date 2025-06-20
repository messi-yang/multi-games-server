package dto

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/commandmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/gamemodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/playermodel"
	"github.com/google/uuid"
)

type CommandDto struct {
	Id        uuid.UUID              `json:"id"`
	GameId    uuid.UUID              `json:"gameId"`
	PlayerId  uuid.UUID              `json:"playerId"`
	Timestamp int64                  `json:"timestamp"`
	Name      string                 `json:"name"`
	Payload   map[string]interface{} `json:"payload"`
}

func NewCommandDto(command commandmodel.Command) CommandDto {
	return CommandDto{
		Id:        command.GetId().Uuid(),
		GameId:    command.GetGameId().Uuid(),
		PlayerId:  command.GetPlayerId().Uuid(),
		Timestamp: command.GetTimestamp(),
		Name:      command.GetName(),
		Payload:   command.GetPayload(),
	}
}

func ParseCommandDto(commandDto CommandDto) (command commandmodel.Command, err error) {
	return commandmodel.LoadCommand(
		commandmodel.NewCommandId(commandDto.Id),
		gamemodel.NewGameId(commandDto.GameId),
		playermodel.NewPlayerId(commandDto.PlayerId),
		commandDto.Timestamp,
		commandDto.Name,
		commandDto.Payload,
	), nil
}
