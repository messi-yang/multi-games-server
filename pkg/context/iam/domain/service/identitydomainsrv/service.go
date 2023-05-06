package identitydomainsrv

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/domain/model/usermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/google/uuid"
)

type Service interface {
	Register(emailAddress string, username string) (user usermodel.User, err error)
}

type serve struct {
	userRepo              usermodel.Repo
	domainEventDispatcher domain.DomainEventDispatcher
}

func NewService(
	userRepo usermodel.Repo,
	domainEventDispatcher domain.DomainEventDispatcher,
) Service {
	return &serve{
		userRepo:              userRepo,
		domainEventDispatcher: domainEventDispatcher,
	}
}

func (serve *serve) Register(emailAddress string, username string) (user usermodel.User, err error) {
	user = usermodel.NewUser(sharedkernelmodel.NewUserId(uuid.New()), emailAddress, username)
	if err = serve.userRepo.Add(user); err != nil {
		return user, err
	}
	if err = serve.domainEventDispatcher.Dispatch(&user); err != nil {
		return user, err
	}
	return user, nil
}
