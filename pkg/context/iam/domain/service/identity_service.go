package service

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/identitymodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/google/uuid"
)

type IdentityService interface {
	Register(emailAddress string, username string) (user identitymodel.User, err error)
}

type identityServe struct {
	userRepo identitymodel.UserRepo
}

func NewIdentityService(
	userRepo identitymodel.UserRepo,
) IdentityService {
	return &identityServe{
		userRepo: userRepo,
	}
}

func (identityServe *identityServe) Register(emailAddress string, username string) (user identitymodel.User, err error) {
	user = identitymodel.NewUser(sharedkernelmodel.NewUserId(uuid.New()), emailAddress, username)
	if err = identityServe.userRepo.Add(user); err != nil {
		return user, err
	}
	return user, nil
}
