package accessmodel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWorldRoleName(t *testing.T) {
	roleName, err := NewWorldRoleName("admin")
	expectedRoleName := WorldRoleName{
		name: "admin",
	}

	assert.Nil(t, err, "error should be nil when creating a valid world role name")
	assert.Equal(t, expectedRoleName, roleName, "created world role name should match the expected value")
}

func TestNewWorldRoleName_InvalidName(t *testing.T) {
	_, err := NewWorldRoleName("invalid")

	assert.Error(t, err, "error should be returned for an invalid world role name")
}

func TestWorldRoleName_String(t *testing.T) {
	roleName, _ := NewWorldRoleName("admin")

	expectedString := "admin"
	assert.Equal(t, expectedString, roleName.String(), "string representation of the world role name should match the expected value")
}
