package worldcommonmodel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnitType(t *testing.T) {
	t.Run("NewUnitType", func(t *testing.T) {
		unitType, err := NewUnitType("static")
		assert.NoError(t, err)
		assert.Equal(t, unitTypeStatic, unitType.value)

		unitType, err = NewUnitType("portal")
		assert.NoError(t, err)
		assert.Equal(t, unitTypePortal, unitType.value)

		unitType, err = NewUnitType("invalid")
		assert.Error(t, err)
	})

	t.Run("UnitType", func(t *testing.T) {
		t.Run("IsEqual", func(t *testing.T) {
			unitType1, _ := NewUnitType("static")
			unitType2, _ := NewUnitType("static")
			unitType3, _ := NewUnitType("portal")
			assert.True(t, unitType1.IsEqual(unitType2))
			assert.False(t, unitType1.IsEqual(unitType3))
		})

		t.Run("IsStatic", func(t *testing.T) {
			unitType1, _ := NewUnitType("static")
			unitType2, _ := NewUnitType("portal")
			assert.True(t, unitType1.IsStatic())
			assert.False(t, unitType2.IsStatic())
		})

		t.Run("IsPortal", func(t *testing.T) {
			unitType1, _ := NewUnitType("portal")
			unitType2, _ := NewUnitType("static")
			assert.True(t, unitType1.IsPortal())
			assert.False(t, unitType2.IsPortal())
		})
	})
}
