package roomaccessmodel

import (
	"testing"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/stretchr/testify/assert"
)

func TestRoomPermission(t *testing.T) {
	t.Run("NewRoomPermission", func(t *testing.T) {
		roomRole, _ := globalcommonmodel.NewRoomRole("owner")
		roomPermission := NewRoomPermission(roomRole)

		assert.Equal(t, roomPermission, RoomPermission{
			role: roomRole,
		})
	})

	t.Run("RoomPermission", func(t *testing.T) {
		t.Run("IsEqual", func(t *testing.T) {
			roomRoleOwner, _ := globalcommonmodel.NewRoomRole("owner")
			roomRoleAdmin, _ := globalcommonmodel.NewRoomRole("admin")
			roomPermission1 := NewRoomPermission(roomRoleOwner)
			roomPermission2 := NewRoomPermission(roomRoleOwner)
			roomPermission3 := NewRoomPermission(roomRoleAdmin)

			assert.True(t, roomPermission1.IsEqual(roomPermission2), "roomPermission1 should be equal to roomPermission2")
			assert.False(t, roomPermission1.IsEqual(roomPermission3), "roomPermission1 should not be equal to roomPermission3")
		})

		t.Run("CanGetRoomMembers", func(t *testing.T) {
			roomRoleOwner, _ := globalcommonmodel.NewRoomRole("owner")
			roomRoleAdmin, _ := globalcommonmodel.NewRoomRole("admin")
			roomRoleEditor, _ := globalcommonmodel.NewRoomRole("editor")
			roomRoleViewer, _ := globalcommonmodel.NewRoomRole("viewer")
			roomPermission1 := NewRoomPermission(roomRoleOwner)
			roomPermission2 := NewRoomPermission(roomRoleAdmin)
			roomPermission3 := NewRoomPermission(roomRoleEditor)
			roomPermission4 := NewRoomPermission(roomRoleViewer)

			assert.True(t, roomPermission1.CanGetRoomMembers(), "owner should be able to get room members")
			assert.True(t, roomPermission2.CanGetRoomMembers(), "admin should be able to get room members")
			assert.True(t, roomPermission3.CanGetRoomMembers(), "editor should be able to get room members")
			assert.True(t, roomPermission4.CanGetRoomMembers(), "viewer should be able to get room members")
		})

		t.Run("CanUpdateRoom", func(t *testing.T) {
			roomRoleOwner, _ := globalcommonmodel.NewRoomRole("owner")
			roomRoleAdmin, _ := globalcommonmodel.NewRoomRole("admin")
			roomRoleEditor, _ := globalcommonmodel.NewRoomRole("editor")
			roomRoleViewer, _ := globalcommonmodel.NewRoomRole("viewer")
			roomPermission1 := NewRoomPermission(roomRoleOwner)
			roomPermission2 := NewRoomPermission(roomRoleAdmin)
			roomPermission3 := NewRoomPermission(roomRoleEditor)
			roomPermission4 := NewRoomPermission(roomRoleViewer)

			assert.True(t, roomPermission1.CanUpdateRoom(), "owner should be able to update room")
			assert.True(t, roomPermission2.CanUpdateRoom(), "admin should be able to update room")
			assert.False(t, roomPermission3.CanUpdateRoom(), "editor should not be able to update room")
			assert.False(t, roomPermission4.CanUpdateRoom(), "viewer should not be able to update room")
		})

		t.Run("CanDeleteRoom", func(t *testing.T) {
			roomRoleOwner, _ := globalcommonmodel.NewRoomRole("owner")
			roomRoleAdmin, _ := globalcommonmodel.NewRoomRole("admin")
			roomRoleEditor, _ := globalcommonmodel.NewRoomRole("editor")
			roomRoleViewer, _ := globalcommonmodel.NewRoomRole("viewer")
			roomPermission1 := NewRoomPermission(roomRoleOwner)
			roomPermission2 := NewRoomPermission(roomRoleAdmin)
			roomPermission3 := NewRoomPermission(roomRoleEditor)
			roomPermission4 := NewRoomPermission(roomRoleViewer)

			assert.True(t, roomPermission1.CanDeleteRoom(), "owner should be able to delete room")
			assert.False(t, roomPermission2.CanDeleteRoom(), "admin should not be able to delete room")
			assert.False(t, roomPermission3.CanDeleteRoom(), "editor should not be able to delete room")
			assert.False(t, roomPermission4.CanDeleteRoom(), "viewer should not be able to delete room")
		})
	})
}
