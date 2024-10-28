package globalcommonmodel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUrl(t *testing.T) {
	t.Run("NewUrl", func(t *testing.T) {
		t.Run("Should throw error when url is not valid", func(t *testing.T) {
			_, err := NewUrl("dsafads")
			assert.Error(t, err)
		})

		t.Run("Should throw error when url is not https", func(t *testing.T) {
			_, err := NewUrl("http://google.com")
			assert.Error(t, err)
		})

		t.Run("Should not throw error when url is valid", func(t *testing.T) {
			_, err := NewUrl("https://google.com")
			assert.NoError(t, err)
		})
	})

	t.Run("IsEqual", func(t *testing.T) {
		t.Run("Should compare urls correctly", func(t *testing.T) {
			url1, _ := NewUrl("https://google.com")
			url2, _ := NewUrl("https://google.com")
			url3, _ := NewUrl("https://google.com/hello")
			assert.True(t, url1.IsEqual(url2))
			assert.False(t, url1.IsEqual(url3))
		})
	})
}
