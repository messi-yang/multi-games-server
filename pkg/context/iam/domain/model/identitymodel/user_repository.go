package identitymodel

import "github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"

type UserRepo interface {
	Add(User) error
	Get(sharedkernelmodel.UserId) (user User, err error)
	GetUserByEmailAddress(emailAddress sharedkernelmodel.EmailAddress) (user *User, err error)
}
