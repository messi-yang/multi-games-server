package worldappservice

import (
	"github.com/google/uuid"
)

type CreateWorldCommand struct {
	UserId uuid.UUID
}
