package usermodel

import "github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"

type Repository interface {
	Add(UserAgg) error
	Get(sharedkernelmodel.UserIdVo) (user UserAgg, err error)
	GetByEmailAddress(emailAddress string) (user UserAgg, exists bool, err error)
}
