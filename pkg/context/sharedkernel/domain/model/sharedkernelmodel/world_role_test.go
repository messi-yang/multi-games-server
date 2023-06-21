package sharedkernelmodel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorldRole(t *testing.T) {
	t.Run("NewWorldRole", func(t *testing.T) {
		worldRole, err := NewWorldRole("admin")
		assert.NoError(t, err)
		assert.Equal(t, worldRole.String(), "admin")

		_, err = NewWorldRole("invalid")
		assert.Error(t, err)
	})

	t.Run("WorldRole", func(t *testing.T) {
		t.Run("IsEqual", func(t *testing.T) {
			worldRole1, _ := NewWorldRole("admin")
			worldRole2, _ := NewWorldRole("admin")
			worldRole3, _ := NewWorldRole("owner")
			assert.True(t, worldRole1.IsEqual(worldRole2))
			assert.False(t, worldRole1.IsEqual(worldRole3))
		})
		t.Run("String", func(t *testing.T) {
			worldRole, _ := NewWorldRole("admin")
			assert.Equal(t, worldRole.String(), "admin")
		})

		t.Run("IsOwner", func(t *testing.T) {
			worldRole, _ := NewWorldRole("owner")
			assert.True(t, worldRole.IsOwner())
		})

		t.Run("IsAdmin", func(t *testing.T) {
			worldRole, _ := NewWorldRole("admin")
			assert.True(t, worldRole.IsAdmin())
		})

		t.Run("IsEditor", func(t *testing.T) {
			worldRole, _ := NewWorldRole("editor")
			assert.True(t, worldRole.IsEditor())
		})

		t.Run("IsViewer", func(t *testing.T) {
			worldRole, _ := NewWorldRole("viewer")
			assert.True(t, worldRole.IsViewer())
		})
	})
}
