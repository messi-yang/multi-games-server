package worldapi

import "github.com/google/uuid"

type createWorldRequestDto = struct {
	UserId uuid.UUID `json:"userId"`
}
