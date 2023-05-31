package accessmodel

import "fmt"

type WorldRoleNameValue string

const (
	WorldRoleAdmin WorldRoleNameValue = "admin"
)

type WorldRoleName struct {
	name WorldRoleNameValue
}

func NewWorldRoleName(worldRoleNameStr string) (WorldRoleName, error) {
	switch worldRoleNameStr {
	case "admin":
		return WorldRoleName{
			name: WorldRoleAdmin,
		}, nil
	default:
		return WorldRoleName{}, fmt.Errorf("invalid WorldRoleName: %s", worldRoleNameStr)
	}
}

func (worldRoleName WorldRoleName) String() string {
	return string(worldRoleName.name)
}
