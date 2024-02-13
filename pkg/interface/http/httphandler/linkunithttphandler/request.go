package linkunithttphandler

import (
	"github.com/google/uuid"
)

type getLinkUnitUrlRequestBody struct {
	Id uuid.UUID `json:"id"`
}
