package identityappsrv

import (
	"errors"
	"fmt"
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/identityaccess/application/jsondto"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/identityaccess/domain/model/usermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/identityaccess/domain/service/identitydomainsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/util/uuidutil"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Service interface {
	FindUserByEmailAddress(FindUserByEmailAddressQuery) (userAggDto jsondto.UserAggDto, found bool, err error)
	Register(RegisterCommand) (userIdDto uuid.UUID, err error)
	Login(LoginCommand) (accessToken string, err error)
	Validate(accessToken string) (userIdDto uuid.UUID, err error)
}

type serve struct {
	userRepository  usermodel.Repository
	identityService identitydomainsrv.Service
	authSecret      string
}

func NewService(userRepository usermodel.Repository, identityService identitydomainsrv.Service, authSecret string) Service {
	return &serve{userRepository: userRepository, identityService: identityService, authSecret: authSecret}
}

func (serve *serve) FindUserByEmailAddress(query FindUserByEmailAddressQuery) (userAggDto jsondto.UserAggDto, found bool, err error) {
	user, found, err := serve.userRepository.FindUserByEmailAddress(query.EmailAddress)
	if err != nil {
		return userAggDto, found, err
	}
	if !found {
		return userAggDto, false, err
	}
	return jsondto.NewUserAggDto(user), true, nil
}

func (serve *serve) Register(command RegisterCommand) (userIdDto uuid.UUID, err error) {
	_, userFound, err := serve.userRepository.FindUserByEmailAddress(command.EmailAddress)
	if err != nil {
		return userIdDto, err
	}

	if userFound {
		return userIdDto, fmt.Errorf("user with email address of %s already exists", command.EmailAddress)
	}

	user, err := serve.identityService.Register(command.EmailAddress, "New User")
	if err != nil {
		return userIdDto, err
	}

	return user.GetId().Uuid(), nil
}

func (serve *serve) Login(command LoginCommand) (accessToken string, err error) {
	user, err := serve.userRepository.Get(sharedkernelmodel.NewUserIdVo(command.UserId))
	if err != nil {
		return accessToken, err
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

func (serve *serve) Validate(accessToken string) (userIdDto uuid.UUID, err error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(serve.authSecret), nil
	})
	if err != nil {
		return userIdDto, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userIdDto := uuidutil.UnsafelyNewUuid(claims["jti"].(string))
		return userIdDto, nil
	} else {
		return userIdDto, errors.New("token is not valid")
	}
}
