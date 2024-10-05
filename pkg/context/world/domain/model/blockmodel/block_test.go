package blockmodel

import (
	"testing"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestBlock(t *testing.T) {
	t.Run("LoadBlock", func(t *testing.T) {
		blockId := NewBlockId(globalcommonmodel.NewWorldId(uuid.New()), 10, 10)
		block := LoadBlock(blockId)
		assert.Equal(t, block.GetId(), blockId)
		assert.Equal(t, block.GetX(), 10)
		assert.Equal(t, block.GetZ(), 10)
	})

	t.Run("GetDimension", func(t *testing.T) {
		blockId := NewBlockId(globalcommonmodel.NewWorldId(uuid.New()), 10, 10)
		block := LoadBlock(blockId)
		dimension, _ := worldcommonmodel.NewDimension(50, 50)
		assert.True(t, block.GetDimension().IsEqual(dimension))
	})

	t.Run("GetBound", func(t *testing.T) {
		blockId := NewBlockId(globalcommonmodel.NewWorldId(uuid.New()), 10, 10)
		block := LoadBlock(blockId)
		bound, _ := worldcommonmodel.NewBound(worldcommonmodel.NewPosition(500, 500), worldcommonmodel.NewPosition(549, 549))
		assert.True(t, block.GetBound().IsEqual(bound))
	})
}
