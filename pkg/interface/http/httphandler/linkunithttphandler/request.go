package linkunithttphandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/google/uuid"
)

type getLinkUnitUrlRequestBody struct {
	WorldId  uuid.UUID       `json:"worldId"`
	Position dto.PositionDto `json:"position"`
}
