package globalcommonmodel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColor(t *testing.T) {
	t.Run("NewColorFromHexString", func(t *testing.T) {
		t.Run("Should throw error when hex string is not valid", func(t *testing.T) {
			_, err := NewColorFromHexString("#dmdmdm")
			assert.Error(t, err)
		})
		t.Run("Should not throw error when hex string is valid", func(t *testing.T) {
			_, err := NewColorFromHexString("#fafafa")
			assert.Nil(t, err)
		})
	})

	t.Run("NewColor", func(t *testing.T) {
		t.Run("Should throw error when any of the r, g, b is not valid", func(t *testing.T) {
			_, err := NewColor(0, 100, 300)
			assert.Error(t, err)
		})

		t.Run("Should now error when all of the r, g, b are valid", func(t *testing.T) {
			_, err := NewColor(0, 100, 200)
			assert.NoError(t, err)
		})
	})

	t.Run("HexString", func(t *testing.T) {
		t.Run("Color string should be a valid hex string with length of 7", func(t *testing.T) {
			color, _ := NewColor(63, 31, 15)
			assert.Equal(t, "#3f1f0f", color.HexString())
		})
	})

	t.Run("IsEqual", func(t *testing.T) {
		t.Run("Should compare colors correctly", func(t *testing.T) {
			color1, _ := NewColor(63, 31, 15)
			color2, _ := NewColor(63, 31, 15)
			color3, _ := NewColor(163, 131, 115)
			assert.True(t, color1.IsEqual(color2))
			assert.False(t, color1.IsEqual(color3))
		})
	})
}
