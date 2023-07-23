package worldaccessmodel

import (
	"testing"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/stretchr/testify/assert"
)

func TestWorldPermission(t *testing.T) {
	t.Run("NewWorldPermission", func(t *testing.T) {
		worldRole, _ := globalcommonmodel.NewWorldRole("owner")
		worldPermission := NewWorldPermission(worldRole)

		assert.Equal(t, worldPermission, WorldPermission{
			role: worldRole,
		})
	})

	t.Run("WorldPermission", func(t *testing.T) {
		t.Run("IsEqual", func(t *testing.T) {
			worldRoleOwner, _ := globalcommonmodel.NewWorldRole("owner")
			worldRoleAdmin, _ := globalcommonmodel.NewWorldRole("admin")
			worldPermission1 := NewWorldPermission(worldRoleOwner)
			worldPermission2 := NewWorldPermission(worldRoleOwner)
			worldPermission3 := NewWorldPermission(worldRoleAdmin)

			assert.True(t, worldPermission1.IsEqual(worldPermission2), "worldPermission1 should be equal to worldPermission2")
			assert.False(t, worldPermission1.IsEqual(worldPermission3), "worldPermission1 should not be equal to worldPermission3")
		})

		t.Run("CanGetWorldMembers", func(t *testing.T) {
			worldRoleOwner, _ := globalcommonmodel.NewWorldRole("owner")
			worldRoleAdmin, _ := globalcommonmodel.NewWorldRole("admin")
			worldRoleEditor, _ := globalcommonmodel.NewWorldRole("editor")
			worldRoleViewer, _ := globalcommonmodel.NewWorldRole("viewer")
			worldPermission1 := NewWorldPermission(worldRoleOwner)
			worldPermission2 := NewWorldPermission(worldRoleAdmin)
			worldPermission3 := NewWorldPermission(worldRoleEditor)
			worldPermission4 := NewWorldPermission(worldRoleViewer)

			assert.True(t, worldPermission1.CanGetWorldMembers(), "owner should be able to get world members")
			assert.True(t, worldPermission2.CanGetWorldMembers(), "admin should be able to get world members")
			assert.True(t, worldPermission3.CanGetWorldMembers(), "editor should be able to get world members")
			assert.True(t, worldPermission4.CanGetWorldMembers(), "viewer should be able to get world members")
		})

		t.Run("CanUpdateWorld", func(t *testing.T) {
			worldRoleOwner, _ := globalcommonmodel.NewWorldRole("owner")
			worldRoleAdmin, _ := globalcommonmodel.NewWorldRole("admin")
			worldRoleEditor, _ := globalcommonmodel.NewWorldRole("editor")
			worldRoleViewer, _ := globalcommonmodel.NewWorldRole("viewer")
			worldPermission1 := NewWorldPermission(worldRoleOwner)
			worldPermission2 := NewWorldPermission(worldRoleAdmin)
			worldPermission3 := NewWorldPermission(worldRoleEditor)
			worldPermission4 := NewWorldPermission(worldRoleViewer)

			assert.True(t, worldPermission1.CanUpdateWorld(), "owner should be able to update world")
			assert.True(t, worldPermission2.CanUpdateWorld(), "admin should be able to update world")
			assert.False(t, worldPermission3.CanUpdateWorld(), "editor should not be able to update world")
			assert.False(t, worldPermission4.CanUpdateWorld(), "viewer should not be able to update world")
		})

		t.Run("CanDeleteWorld", func(t *testing.T) {
			worldRoleOwner, _ := globalcommonmodel.NewWorldRole("owner")
			worldRoleAdmin, _ := globalcommonmodel.NewWorldRole("admin")
			worldRoleEditor, _ := globalcommonmodel.NewWorldRole("editor")
			worldRoleViewer, _ := globalcommonmodel.NewWorldRole("viewer")
			worldPermission1 := NewWorldPermission(worldRoleOwner)
			worldPermission2 := NewWorldPermission(worldRoleAdmin)
			worldPermission3 := NewWorldPermission(worldRoleEditor)
			worldPermission4 := NewWorldPermission(worldRoleViewer)

			assert.True(t, worldPermission1.CanDeleteWorld(), "owner should be able to delete world")
			assert.False(t, worldPermission2.CanDeleteWorld(), "admin should not be able to delete world")
			assert.False(t, worldPermission3.CanDeleteWorld(), "editor should not be able to delete world")
			assert.False(t, worldPermission4.CanDeleteWorld(), "viewer should not be able to delete world")
		})
	})
}
