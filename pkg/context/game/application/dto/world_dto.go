package dto

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/worldmodel"
	"github.com/google/uuid"
)

type WorldDto struct {
	Id        uuid.UUID `json:"id"`
	UserId    uuid.UUID `json:"userId"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewWorldDto(world worldmodel.World) WorldDto {
	return WorldDto{
		Id:        world.GetId().Uuid(),
		UserId:    world.GetUserId().Uuid(),
		Name:      world.GetName(),
		CreatedAt: world.GetCreatedAt(),
		UpdatedAt: world.GetUpdatedAt(),
	}
}
