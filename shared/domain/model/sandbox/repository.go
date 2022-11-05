package sandbox

import (
	"github.com/google/uuid"
)

type SandboxRepository interface {
	Add(Sandbox) error
	Get(id uuid.UUID) (sandbox Sandbox, err error)
	GetFirstGameId() (id uuid.UUID, err error)
	Update(id uuid.UUID, sandbox Sandbox) error
}
