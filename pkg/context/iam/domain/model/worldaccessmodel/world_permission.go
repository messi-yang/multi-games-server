package worldaccessmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
)

type WorldPermission struct {
	role globalcommonmodel.WorldRole
}

// Interface Implementation Check
var _ domain.ValueObject[WorldPermission] = (*WorldPermission)(nil)

func NewWorldPermission(role globalcommonmodel.WorldRole) WorldPermission {
	return WorldPermission{
		role: role,
	}
}

func (worldPermission WorldPermission) IsEqual(otherWorldPermission WorldPermission) bool {
	return worldPermission.role.IsEqual(otherWorldPermission.role)
}

func (worldPermission WorldPermission) CanGetWorldMembers() bool {
	return true
}

func (worldPermission WorldPermission) CanUpdateWorld() bool {
	return worldPermission.role.IsOwner() || worldPermission.role.IsAdmin()
}

func (worldPermission WorldPermission) CanDeleteWorld() bool {
	return worldPermission.role.IsOwner()
}
