package dto

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/gamemodel"
	"github.com/google/uuid"
)

type GameDto struct {
	Id        uuid.UUID               `json:"id"`
	RoomId    uuid.UUID               `json:"roomId"`
	Name      string                  `json:"name"`
	Started   bool                    `json:"started"`
	State     *map[string]interface{} `json:"state"`
	CreatedAt time.Time               `json:"createdAt"`
	UpdatedAt time.Time               `json:"updatedAt"`
}

func NewGameDto(game gamemodel.Game) GameDto {
	return GameDto{
		Id:        game.GetId().Uuid(),
		RoomId:    game.GetRoomId().Uuid(),
		Name:      game.GetName(),
		Started:   game.GetStarted(),
		State:     game.GetState(),
		CreatedAt: game.GetCreatedAt(),
		UpdatedAt: game.GetUpdatedAt(),
	}
}
