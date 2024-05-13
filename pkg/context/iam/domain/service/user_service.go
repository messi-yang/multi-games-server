package service

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/usermodel"
	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(emailAddress globalcommonmodel.EmailAddress) (userId globalcommonmodel.UserId, err error)
}

type userServe struct {
	userRepo usermodel.UserRepo
}

func NewUserService(userRepo usermodel.UserRepo) UserService {
	return &userServe{userRepo: userRepo}
}

func (userServe *userServe) CreateUser(emailAddress globalcommonmodel.EmailAddress) (userId globalcommonmodel.UserId, err error) {
	user, err := userServe.userRepo.GetUserByEmailAddress(emailAddress)
	if err != nil {
		return userId, err
	}

	if user != nil {
		return userId, fmt.Errorf("user with email address of %s already exists", emailAddress)
	}

	randomUsername := globalcommonmodel.NewRandomUsername()
	randomeFriendlyName, err := usermodel.NewFriendlyName(randomUsername.String())
	if err != nil {
		return userId, err
	}
	newUser := usermodel.NewUser(
		globalcommonmodel.NewUserId(uuid.New()),
		emailAddress,
		globalcommonmodel.NewRandomUsername(),
		randomeFriendlyName,
	)
	if err = userServe.userRepo.Add(newUser); err != nil {
		return userId, err
	}
	return newUser.GetId(), nil
}
