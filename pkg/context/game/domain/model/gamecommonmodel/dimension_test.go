package gamecommonmodel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDimension(t *testing.T) {
	t.Run("NewDimension", func(t *testing.T) {
		t.Run("Should throw error when width or depth is less than 1", func(t *testing.T) {
			_, err := NewDimension(0, 1)
			assert.Error(t, err)

			_, err = NewDimension(1, 0)
			assert.Error(t, err)
		})
	})

	t.Run("IsEqual", func(t *testing.T) {
		t.Run("Should correctly check the equalty of two dimensions", func(t *testing.T) {
			dimension1, _ := NewDimension(2, 2)
			dimension2, _ := NewDimension(2, 2)
			dimension3, _ := NewDimension(2, 3)

			assert.True(t, dimension1.IsEqual(dimension2))
			assert.False(t, dimension1.IsEqual(dimension3))
		})
	})
}
