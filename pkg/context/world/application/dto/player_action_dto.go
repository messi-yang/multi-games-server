package dto

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/playermodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
)

type PlayerActionDto struct {
	Name            string             `json:"name"`
	PrecisePosition PrecisePositionDto `json:"precisePosition"`
	Direction       int8               `json:"direction"`
	Time            int64              `json:"time"`
}

func NewPlayerActionDto(playerAction playermodel.PlayerAction) PlayerActionDto {
	return PlayerActionDto{
		Name:            string(playerAction.GetName()),
		PrecisePosition: NewPrecisePositionDto(playerAction.GetPrecisePosition()),
		Time:            playerAction.GetTime().UnixMilli(),
	}
}

func ParsePlayerActionDto(dto PlayerActionDto) (playermodel.PlayerAction, error) {
	actionName, err := playermodel.ParsePlayerActionNameEnum(dto.Name)
	if err != nil {
		return playermodel.PlayerAction{}, err
	}
	return playermodel.NewPlayerAction(
		actionName,
		worldcommonmodel.NewPrecisePosition(dto.PrecisePosition.X, dto.PrecisePosition.Z),
		worldcommonmodel.NewDirection(dto.Direction),
		time.UnixMilli(dto.Time),
	), nil
}
