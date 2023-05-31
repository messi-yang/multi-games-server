package identitymodel

import "github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"

type UserRepo interface {
	Add(User) error
	Get(sharedkernelmodel.UserId) (user User, err error)
	FindUserByEmailAddress(emailAddress string) (user User, userFound bool, err error)
}
