package commonmodel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSize(t *testing.T) {
	size, err := NewSize(10, 5)
	assert.NoError(t, err, "NewSize should not return an error")
	assert.Equal(t, Size{10, 5}, size, "size should have the provided width and height")

	_, err = NewSize(0, 5)
	assert.Error(t, err, "NewSize should return an error")

	_, err = NewSize(10, 0)
	assert.Error(t, err, "NewSize should return an error")
}

func TestSize_IsEqual(t *testing.T) {
	size1 := Size{10, 5}
	size2 := Size{10, 5}
	size3 := Size{8, 3}

	assert.True(t, size1.IsEqual(size2), "size1 should be equal to size2")
	assert.False(t, size1.IsEqual(size3), "size1 should not be equal to size3")
}

func TestSize_GetWidth(t *testing.T) {
	size := Size{10, 5}
	assert.Equal(t, 10, size.GetWidth(), "GetWidth() should return the width of the size")
}

func TestSize_GetHeight(t *testing.T) {
	size := Size{10, 5}
	assert.Equal(t, 5, size.GetHeight(), "GetHeight() should return the height of the size")
}
