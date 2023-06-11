package worldaccessmodel

import "fmt"

type WorldRoleValue string

const (
	WorldRoleOwner  WorldRoleValue = "owner"
	WorldRoleAdmin  WorldRoleValue = "admin"
	WorldRoleEditor WorldRoleValue = "editor"
	WorldRoleViewer WorldRoleValue = "viwer"
)

type WorldRole struct {
	name WorldRoleValue
}

func NewWorldRole(worldRoleValue string) (WorldRole, error) {
	switch worldRoleValue {
	case "owner":
		return WorldRole{
			name: WorldRoleOwner,
		}, nil
	case "admin":
		return WorldRole{
			name: WorldRoleAdmin,
		}, nil
	case "editor":
		return WorldRole{
			name: WorldRoleEditor,
		}, nil
	case "viewer":
		return WorldRole{
			name: WorldRoleViewer,
		}, nil
	default:
		return WorldRole{}, fmt.Errorf("invalid WorldRole: %s", worldRoleValue)
	}
}

func (worldRole WorldRole) String() string {
	return string(worldRole.name)
}
