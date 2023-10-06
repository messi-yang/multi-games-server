package worldjourneyhandler

import (
	"fmt"

	"github.com/google/uuid"
)

func newWorldServerMessageChannel(worldIdDto uuid.UUID) string {
	return fmt.Sprintf("WORLD_%s_CHANNEL", worldIdDto)
}
