package linkunitappsrv

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/google/uuid"
)

type CreateLinkUnitCommand struct {
	Id        uuid.UUID
	WorldId   uuid.UUID
	ItemId    uuid.UUID
	Position  dto.PositionDto
	Direction int8
	Label     *string
	Url       string
}

type RemoveLinkUnitCommand struct {
	Id uuid.UUID
}
