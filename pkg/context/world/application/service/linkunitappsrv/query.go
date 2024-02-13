package linkunitappsrv

import (
	"github.com/google/uuid"
)

type GetLinkUnitUrlQuery struct {
	Id      uuid.UUID
	WorldId uuid.UUID
}
