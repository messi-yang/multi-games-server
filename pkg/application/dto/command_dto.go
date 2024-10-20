package dto

import (
	"github.com/google/uuid"
)

type CommandDto struct {
	Id        uuid.UUID `json:"id"`
	Timestamp int64     `json:"timestamp"`
	Name      string    `json:"name"`
	Payload   any       `json:"payload"`
}
