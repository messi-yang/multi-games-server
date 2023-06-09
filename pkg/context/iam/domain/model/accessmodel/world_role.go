package accessmodel

import "fmt"

type WorldRoleValue string

const (
	WorldRoleAdmin WorldRoleValue = "admin"
)

type WorldRole struct {
	name WorldRoleValue
}

func NewWorldRole(worldRoleValue string) (WorldRole, error) {
	switch worldRoleValue {
	case "admin":
		return WorldRole{
			name: WorldRoleAdmin,
		}, nil
	default:
		return WorldRole{}, fmt.Errorf("invalid WorldRole: %s", worldRoleValue)
	}
}

func (worldRole WorldRole) String() string {
	return string(worldRole.name)
}
