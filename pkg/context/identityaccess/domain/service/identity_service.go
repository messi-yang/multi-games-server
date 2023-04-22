package service

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/identityaccess/domain/model/usermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/google/uuid"
)

type IdentityService interface {
	Register(emailAddress string, username string) (user usermodel.UserAgg, err error)
}

type identityServe struct {
	userRepository usermodel.Repository
}

func NewIdentityService(userRepository usermodel.Repository) IdentityService {
	return &identityServe{userRepository: userRepository}
}

func (serve *identityServe) Register(emailAddress string, username string) (user usermodel.UserAgg, err error) {
	user = usermodel.NewUserAgg(sharedkernelmodel.NewUserIdVo(uuid.New()), emailAddress, username)
	err = serve.userRepository.Add(user)
	if err != nil {
		return user, err
	}
	return user, nil
}
