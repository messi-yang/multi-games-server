package worldappservice

import "github.com/google/uuid"

type CreateWorldRequestDto = struct {
	UserId uuid.UUID `json:"userId"`
}
