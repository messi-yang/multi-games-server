package worldmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
)

type WorldPermission struct {
	role sharedkernelmodel.WorldRole
}

// Interface Implementation Check
var _ domain.ValueObject[WorldPermission] = (*WorldPermission)(nil)

func NewWorldPermission(role sharedkernelmodel.WorldRole) WorldPermission {
	return WorldPermission{
		role: role,
	}
}

func (worldPermission WorldPermission) IsEqual(otherWorldPermission WorldPermission) bool {
	return worldPermission.role.IsEqual(otherWorldPermission.role)
}

func (worldPermission WorldPermission) CanUpdateWorld() bool {
	return worldPermission.role.IsOwner() || worldPermission.role.IsAdmin()
}
