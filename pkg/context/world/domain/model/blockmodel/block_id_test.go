package blockmodel

import (
	"testing"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestBlockId(t *testing.T) {
	t.Run("NewBlockId", func(t *testing.T) {
		worldId := globalcommonmodel.NewWorldId(uuid.New())
		blockId := NewBlockId(worldId, 10, 10)
		assert.Equal(t, blockId.GetWorldId(), worldId)
		assert.Equal(t, blockId.GetX(), 10)
		assert.Equal(t, blockId.GetZ(), 10)
	})
}
