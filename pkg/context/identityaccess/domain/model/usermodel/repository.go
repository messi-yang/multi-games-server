package usermodel

import "github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"

type Repository interface {
	Add(UserAgg) error
	Get(sharedkernelmodel.UserIdVo) (user UserAgg, err error)
	FindUserByEmailAddress(emailAddress string) (user UserAgg, userFound bool, err error)
}
