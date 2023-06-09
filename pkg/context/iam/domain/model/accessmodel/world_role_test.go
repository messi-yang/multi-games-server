package accessmodel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWorldRole(t *testing.T) {
	worldRole, err := NewWorldRole("admin")
	expectedWorldRole := WorldRole{
		name: "admin",
	}

	assert.Nil(t, err, "error should be nil when creating a valid world role name")
	assert.Equal(t, expectedWorldRole, worldRole, "created world role name should match the expected value")
}

func TestNewWorldRole_InvalidName(t *testing.T) {
	_, err := NewWorldRole("invalid")

	assert.Error(t, err, "error should be returned for an invalid world role name")
}

func TestWorldRole_String(t *testing.T) {
	worldRole, _ := NewWorldRole("admin")

	expectedString := "admin"
	assert.Equal(t, expectedString, worldRole.String(), "string representation of the world role name should match the expected value")
}
