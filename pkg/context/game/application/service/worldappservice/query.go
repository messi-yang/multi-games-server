package worldappservice

import (
	"github.com/google/uuid"
)

type GetWorldQuery struct {
	WorldId uuid.UUID
}

type QueryWorldsQuery struct{}
