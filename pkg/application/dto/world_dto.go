package dto

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
	"github.com/google/uuid"
)

type WorldDto struct {
	Id        uuid.UUID `json:"id"`
	UserId    uuid.UUID `json:"userId"`
	Name      string    `json:"name"`
	Bound     BoundDto  `json:"bound"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewWorldDto(world worldmodel.World) WorldDto {
	return WorldDto{
		Id:        world.GetId().Uuid(),
		UserId:    world.GetUserId().Uuid(),
		Name:      world.GetName(),
		Bound:     NewBoundDto(world.GetBound()),
		CreatedAt: world.GetCreatedAt(),
		UpdatedAt: world.GetUpdatedAt(),
	}
}
