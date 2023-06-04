package accessmodel

import "fmt"

type WorldRoleNameValue string

const (
	WorldRoleAdmin WorldRoleNameValue = "admin"
)

type WorldRoleName struct {
	name WorldRoleNameValue
}

func NewWorldRoleName(worldRoleNameValue string) (WorldRoleName, error) {
	switch worldRoleNameValue {
	case "admin":
		return WorldRoleName{
			name: WorldRoleAdmin,
		}, nil
	default:
		return WorldRoleName{}, fmt.Errorf("invalid WorldRoleName: %s", worldRoleNameValue)
	}
}

func (worldRoleName WorldRoleName) String() string {
	return string(worldRoleName.name)
}
