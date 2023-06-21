package sharedkernelmodel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmailAddress(t *testing.T) {
	t.Run("NewEmailAddress", func(t *testing.T) {
		_, err := NewEmailAddress("test@example.com")
		assert.NoError(t, err)
		_, err = NewEmailAddress("dumdumgenius@gmail.com")
		assert.NoError(t, err)

		_, err = NewEmailAddress("invalid-email")
		assert.Error(t, err)
	})

	t.Run("EmailAddress", func(t *testing.T) {
		t.Run("IsEqual", func(t *testing.T) {
			email1, _ := NewEmailAddress("test@example.com")
			email2, _ := NewEmailAddress("test@example.com")
			email3, _ := NewEmailAddress("test2@example.com")
			assert.True(t, email1.IsEqual(email2))
			assert.False(t, email1.IsEqual(email3))
		})
	})
}
