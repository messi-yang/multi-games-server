package sharedkernelmodel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsername(t *testing.T) {
	t.Run("NewUsername", func(t *testing.T) {
		username, err := NewUsername("hello_world")
		assert.NoError(t, err)
		assert.Equal(t, "hello_world", username.String())

		t.Run("Validation", func(t *testing.T) {
			t.Run("Can only have lower cases and underscores", func(t *testing.T) {
				_, err = NewUsername("invalid-name")
				assert.Error(t, err)

				_, err = NewUsername("Invalid-Name")
				assert.Error(t, err)

				_, err = NewUsername("invalid$name")
				assert.Error(t, err)

				_, err = NewUsername("invalid name")
				assert.Error(t, err)
			})

			t.Run("Cannot have two underscores in a row", func(t *testing.T) {
				_, err = NewUsername("invalid__name")
				assert.Error(t, err)
			})

			t.Run("Cannot have underscores at the start and the end", func(t *testing.T) {
				_, err = NewUsername("_invalid_name")
				assert.Error(t, err)

				_, err = NewUsername("invalid_name_")
				assert.Error(t, err)
			})

			t.Run("Length has to be between 8 and 20", func(t *testing.T) {
				_, err = NewUsername("aaaaaaa")
				assert.Error(t, err)
				_, err = NewUsername("aaaaaaaa")
				assert.NoError(t, err)

				_, err = NewUsername("aaaaaaaaaaaaaaaaaaaa")
				assert.NoError(t, err)
				_, err = NewUsername("aaaaaaaaaaaaaaaaaaaaa")
				assert.Error(t, err)
			})
		})
	})

	t.Run("NewUsername", func(t *testing.T) {
		for i := 0; i < 100; i += 1 {
			newRandomUsername := NewRandomUsername()
			_, err := NewUsername(newRandomUsername.String())
			assert.NoError(t, err)
		}
	})

	t.Run("Username", func(t *testing.T) {
		t.Run("IsEqual", func(t *testing.T) {
			username1, _ := NewUsername("hello_world")
			username2, _ := NewUsername("hello_world")
			username3, _ := NewUsername("hello_world_another")
			assert.True(t, username1.IsEqual(username2))
			assert.False(t, username1.IsEqual(username3))
		})
	})
}
