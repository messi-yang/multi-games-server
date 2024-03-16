package embedunitappsrv

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/google/uuid"
)

type CreateEmbedUnitCommand struct {
	Id        uuid.UUID
	WorldId   uuid.UUID
	ItemId    uuid.UUID
	Position  dto.PositionDto
	Direction int8
	Label     *string
	EmbedCode string
}

type RemoveEmbedUnitCommand struct {
	Id uuid.UUID
}
