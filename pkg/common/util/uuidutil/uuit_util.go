package uuidutil

import "github.com/google/uuid"

// Use this only when you're confident that the uuid string is valid
func UnsafelyNewUuid(uuidStr string) uuid.UUID {
	id, _ := uuid.Parse(uuidStr)
	return id
}
