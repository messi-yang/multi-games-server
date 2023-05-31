package identitymodel

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/google/uuid"
)

type IdentityService interface {
	Register(emailAddress string, username string) (user User, err error)
}

type identityServe struct {
	userRepo UserRepo
}

func NewIdentityService(
	userRepo UserRepo,
) IdentityService {
	return &identityServe{
		userRepo: userRepo,
	}
}

func (identityServe *identityServe) Register(emailAddress string, username string) (user User, err error) {
	user = NewUser(sharedkernelmodel.NewUserId(uuid.New()), emailAddress, username)
	if err = identityServe.userRepo.Add(user); err != nil {
		return user, err
	}
	return user, nil
}
