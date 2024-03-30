package fenceunitappsrv

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/google/uuid"
)

type CreateFenceUnitCommand struct {
	Id        uuid.UUID
	WorldId   uuid.UUID
	ItemId    uuid.UUID
	Position  dto.PositionDto
	Direction int8
}

type RemoveFenceUnitCommand struct {
	Id uuid.UUID
}
