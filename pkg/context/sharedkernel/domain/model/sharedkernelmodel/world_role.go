package sharedkernelmodel

import "fmt"

type WorldRoleValue string

const (
	WorldRoleOwner  WorldRoleValue = "owner"
	WorldRoleAdmin  WorldRoleValue = "admin"
	WorldRoleEditor WorldRoleValue = "editor"
	WorldRoleViewer WorldRoleValue = "viwer"
)

type WorldRole struct {
	value WorldRoleValue
}

func NewWorldRole(worldRoleValue string) (WorldRole, error) {
	switch worldRoleValue {
	case "owner":
		return WorldRole{
			value: WorldRoleOwner,
		}, nil
	case "admin":
		return WorldRole{
			value: WorldRoleAdmin,
		}, nil
	case "editor":
		return WorldRole{
			value: WorldRoleEditor,
		}, nil
	case "viewer":
		return WorldRole{
			value: WorldRoleViewer,
		}, nil
	default:
		return WorldRole{}, fmt.Errorf("invalid WorldRole: %s", worldRoleValue)
	}
}

func (worldRole WorldRole) CanUpdateWorldInfo() bool {
	return worldRole.value == "owner" || worldRole.value == "admin"
}

func (worldRole WorldRole) String() string {
	return string(worldRole.value)
}
