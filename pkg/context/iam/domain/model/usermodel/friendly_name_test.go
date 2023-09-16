package usermodel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFriendlyName(t *testing.T) {
	t.Run("NewFriendlyName", func(t *testing.T) {
		friendlyName, _ := NewFriendlyName("john")
		assert.Equal(t, friendlyName, FriendlyName{
			value: "john",
		})

		_, err := NewFriendlyName("")
		assert.Error(t, err)

		_, err = NewFriendlyName("+++++++++++++++++++++")
		assert.Error(t, err)
	})

	t.Run("FriendlyName", func(t *testing.T) {
		t.Run("IsEqual", func(t *testing.T) {

			friendlyNameA, _ := NewFriendlyName("john")
			friendlyNameB, _ := NewFriendlyName("john")
			friendlyNameC, _ := NewFriendlyName("tom")
			assert.True(t, friendlyNameA.IsEqual(friendlyNameB))
			assert.False(t, friendlyNameA.IsEqual(friendlyNameC))
		})
	})
}
