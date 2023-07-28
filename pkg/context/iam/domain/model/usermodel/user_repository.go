package usermodel

import "github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"

type UserRepo interface {
	Add(User) error
	Update(User) error
	Get(globalcommonmodel.UserId) (user User, err error)
	GetUserByEmailAddress(emailAddress globalcommonmodel.EmailAddress) (user *User, err error)
	GetUsersOfIds([]globalcommonmodel.UserId) (users []User, err error)
}
