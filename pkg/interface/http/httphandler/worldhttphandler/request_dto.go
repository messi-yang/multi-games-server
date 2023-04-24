package worldhttphandler

import "github.com/google/uuid"

type createWorldRequestDto = struct {
	GamerId uuid.UUID `json:"gamerId"`
}
