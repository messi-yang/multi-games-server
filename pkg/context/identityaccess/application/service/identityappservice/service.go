package identityappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/identityaccess/application/jsondto"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/identityaccess/domain/model/usermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/identityaccess/domain/service"
)

type Service interface {
	LoginOrRegister(LoginOrRegisterCommand) (userDto jsondto.UserAggDto, err error)
}

type serve struct {
	userRepository  usermodel.Repository
	identityService service.IdentityService
}

func NewService(userRepository usermodel.Repository, identityService service.IdentityService) Service {
	return &serve{userRepository: userRepository, identityService: identityService}
}

func (serve *serve) LoginOrRegister(command LoginOrRegisterCommand) (userDto jsondto.UserAggDto, err error) {
	user, userFound, err := serve.userRepository.GetByEmailAddress(command.EmailAddress)
	if err != nil {
		return userDto, err
	}

	if userFound {
		return jsondto.NewUserAggDto(user), nil
	} else {
		user, err := serve.identityService.Register(command.EmailAddress, "New User")
		if err != nil {
			return userDto, err
		}
		return jsondto.NewUserAggDto(user), nil
	}
}
