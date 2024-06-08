package dto

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/playermodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
)

type PlayerActionDto struct {
	Name      string `json:"name"`
	Direction int8   `json:"direction"`
}

func NewPlayerActionDto(playerAction playermodel.PlayerAction) PlayerActionDto {
	return PlayerActionDto{
		Name:      string(playerAction.GetName()),
		Direction: playerAction.GetDirection().Int8(),
	}
}

func ParsePlayerActionDto(dto PlayerActionDto) (playermodel.PlayerAction, error) {
	actionName, err := playermodel.ParsePlayerActionNameEnum(dto.Name)
	if err != nil {
		return playermodel.PlayerAction{}, err
	}
	return playermodel.NewPlayerAction(
		actionName,
		worldcommonmodel.NewDirection(dto.Direction),
	), nil
}
