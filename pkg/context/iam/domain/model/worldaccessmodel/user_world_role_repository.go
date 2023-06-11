package worldaccessmodel

import "github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"

type UserWorldRoleRepo interface {
	Add(UserWorldRole) error
	Get(UserWorldRoleId) (UserWorldRole, error)
	FindWorldRoleOfUser(sharedkernelmodel.WorldId, sharedkernelmodel.UserId) (userWorldRole UserWorldRole, found bool, err error)
	GetUserWorldRolesInWorld(sharedkernelmodel.WorldId) ([]UserWorldRole, error)
}
