package identitydomainsrv

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/identityaccess/domain/model/usermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/google/uuid"
)

type Service interface {
	Register(emailAddress string, username string) (user usermodel.UserAgg, err error)
}

type serve struct {
	userRepository usermodel.Repository
}

func NewService(userRepository usermodel.Repository) Service {
	return &serve{userRepository: userRepository}
}

func (serve *serve) Register(emailAddress string, username string) (user usermodel.UserAgg, err error) {
	user = usermodel.NewUserAgg(sharedkernelmodel.NewUserIdVo(uuid.New()), emailAddress, username)
	err = serve.userRepository.Add(user)
	if err != nil {
		return user, err
	}
	return user, nil
}
