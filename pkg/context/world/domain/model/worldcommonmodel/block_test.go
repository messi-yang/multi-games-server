package worldcommonmodel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlock(t *testing.T) {
	t.Run("NewBlock", func(t *testing.T) {
		block := NewBlock(10, 10)
		assert.Equal(t, block, Block{x: 10, z: 10})
	})

	t.Run("GetDimension", func(t *testing.T) {
		block := NewBlock(10, 10)
		dimension, _ := NewDimension(50, 50)
		assert.True(t, block.GetDimension().IsEqual(dimension))
	})

	t.Run("GetBound", func(t *testing.T) {
		block := NewBlock(10, 10)
		bound, _ := NewBound(NewPosition(500, 500), NewPosition(549, 549))
		assert.True(t, block.GetBound().IsEqual(bound))
	})
}
