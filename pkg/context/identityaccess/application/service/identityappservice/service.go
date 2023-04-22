package identityappservice

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/identityaccess/domain/model/usermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/identityaccess/domain/service"
	"github.com/golang-jwt/jwt/v5"
)

type Service interface {
	LoginOrRegister(LoginOrRegisterCommand) (accessToken string, err error)
}

type serve struct {
	userRepository  usermodel.Repository
	identityService service.IdentityService
	authSecret      string
}

func NewService(userRepository usermodel.Repository, identityService service.IdentityService, authSecret string) Service {
	return &serve{userRepository: userRepository, identityService: identityService, authSecret: authSecret}
}

func (serve *serve) LoginOrRegister(command LoginOrRegisterCommand) (accessToken string, err error) {
	user, userFound, err := serve.userRepository.GetByEmailAddress(command.EmailAddress)
	if err != nil {
		return accessToken, err
	}

	if !userFound {
		user, err = serve.identityService.Register(command.EmailAddress, "New User")
		if err != nil {
			return accessToken, err
		}
	}

	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		ID:        user.GetId().Uuid().String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err = token.SignedString([]byte(serve.authSecret))
	if err != nil {
		return accessToken, err
	}
	return accessToken, nil
}
