package identitydomainsrv

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/domain/model/usermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/google/uuid"
)

type Service interface {
	Register(emailAddress string, username string) (user usermodel.UserAgg, err error)
}

type serve struct {
	userRepo usermodel.Repo
}

func NewService(userRepo usermodel.Repo) Service {
	return &serve{userRepo: userRepo}
}

func (serve *serve) Register(emailAddress string, username string) (user usermodel.UserAgg, err error) {
	user = usermodel.NewUserAgg(sharedkernelmodel.NewUserIdVo(uuid.New()), emailAddress, username)
	err = serve.userRepo.Add(user)
	if err != nil {
		return user, err
	}
	return user, nil
}
