package accessmodel

import "github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"

type UserWorldRoleRepo interface {
	Add(UserWorldRole) error
	Get(UserWorldRoleId) (UserWorldRole, error)
	GetUserWorldRolesInWorld(sharedkernelmodel.WorldId) ([]UserWorldRole, error)
}
