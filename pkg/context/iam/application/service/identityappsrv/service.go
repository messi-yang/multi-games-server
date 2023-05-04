package identityappsrv

import (
	"errors"
	"fmt"
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/application/dto"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/domain/model/usermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/domain/service/identitydomainsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/util/uuidutil"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Service interface {
	FindUserByEmailAddress(FindUserByEmailAddressQuery) (userDto dto.UserDto, found bool, err error)
	Register(RegisterCommand) (userIdDto uuid.UUID, err error)
	Login(LoginCommand) (accessToken string, err error)
	Validate(accessToken string) (userIdDto uuid.UUID, err error)
}

type serve struct {
	userRepo        usermodel.Repo
	identityService identitydomainsrv.Service
	authSecret      string
}

func NewService(userRepo usermodel.Repo, identityService identitydomainsrv.Service, authSecret string) Service {
	return &serve{userRepo: userRepo, identityService: identityService, authSecret: authSecret}
}

func (serve *serve) FindUserByEmailAddress(query FindUserByEmailAddressQuery) (userDto dto.UserDto, found bool, err error) {
	user, found, err := serve.userRepo.FindUserByEmailAddress(query.EmailAddress)
	if err != nil {
		return userDto, found, err
	}
	if !found {
		return userDto, false, err
	}
	return dto.NewUserDto(user), true, nil
}

func (serve *serve) Register(command RegisterCommand) (userIdDto uuid.UUID, err error) {
	_, userFound, err := serve.userRepo.FindUserByEmailAddress(command.EmailAddress)
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
	user, err := serve.userRepo.Get(sharedkernelmodel.NewUserId(command.UserId))
	if err != nil {
		return accessToken, err
	}

	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
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
		// TODO - Check expiration date
		userIdDto := uuidutil.UnsafelyNewUuid(claims["jti"].(string))
		return userIdDto, nil
	} else {
		return userIdDto, errors.New("token is not valid")
	}
}
