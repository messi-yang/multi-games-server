package globalcommonmodel

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
)

type WorldRoleValue string

const (
	WorldRoleOwner  WorldRoleValue = "owner"
	WorldRoleAdmin  WorldRoleValue = "admin"
	WorldRoleEditor WorldRoleValue = "editor"
	WorldRoleViewer WorldRoleValue = "viewer"
)

type WorldRole struct {
	value WorldRoleValue
}

// Interface Implementation Check
var _ domain.ValueObject[WorldRole] = (*WorldRole)(nil)

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

func (worldRole WorldRole) IsEqual(otherWorldRole WorldRole) bool {
	return worldRole.value == otherWorldRole.value
}

func (worldRole WorldRole) String() string {
	return string(worldRole.value)
}

func (worldRole WorldRole) IsOwner() bool {
	return worldRole.value == WorldRoleOwner
}

func (worldRole WorldRole) IsAdmin() bool {
	return worldRole.value == WorldRoleAdmin
}

func (worldRole WorldRole) IsEditor() bool {
	return worldRole.value == WorldRoleEditor
}

func (worldRole WorldRole) IsViewer() bool {
	return worldRole.value == WorldRoleViewer
}
