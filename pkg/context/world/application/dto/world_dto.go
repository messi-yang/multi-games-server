package dto

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/usermodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
	"github.com/google/uuid"
)

type WorldDto struct {
	Id        uuid.UUID   `json:"id"`
	User      dto.UserDto `json:"user"`
	Name      string      `json:"name"`
	Bound     BoundDto    `json:"bound"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
}

func NewWorldDto(world worldmodel.World, user usermodel.User) WorldDto {
	return WorldDto{
		Id:        world.GetId().Uuid(),
		User:      dto.NewUserDto(user),
		Name:      world.GetName(),
		Bound:     NewBoundDto(world.GetBound()),
		CreatedAt: world.GetCreatedAt(),
		UpdatedAt: world.GetUpdatedAt(),
	}
}
