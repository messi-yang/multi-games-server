package authappsrv

import (
	"errors"
	"fmt"
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/usermodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/uuidutil"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Service interface {
	Register(RegisterCommand) (userIdDto uuid.UUID, err error)
	Login(LoginCommand) (accessToken string, err error)
	Validate(accessToken string) (userIdDto uuid.UUID, err error)
}

type serve struct {
	userRepo   usermodel.UserRepo
	authSecret string
}

func NewService(userRepo usermodel.UserRepo, authSecret string) Service {
	return &serve{userRepo: userRepo, authSecret: authSecret}
}

func (serve *serve) Register(command RegisterCommand) (userIdDto uuid.UUID, err error) {
	emailAddress, err := globalcommonmodel.NewEmailAddress(command.EmailAddress)
	if err != nil {
		return userIdDto, err
	}

	user, err := serve.userRepo.GetUserByEmailAddress(emailAddress)
	if err != nil {
		return userIdDto, err
	}

	if user != nil {
		return userIdDto, fmt.Errorf("user with email address of %s already exists", command.EmailAddress)
	}

	randomUsername := globalcommonmodel.NewRandomUsername()
	randomeFriendlyName, err := usermodel.NewFriendlyName(randomUsername.String())
	if err != nil {
		return userIdDto, err
	}
	newUser := usermodel.NewUser(
		globalcommonmodel.NewUserId(uuid.New()),
		emailAddress,
		globalcommonmodel.NewRandomUsername(),
		randomeFriendlyName,
	)
	if err = serve.userRepo.Add(newUser); err != nil {
		return userIdDto, err
	}
	return newUser.GetId().Uuid(), nil
}

func (serve *serve) Login(command LoginCommand) (accessToken string, err error) {
	user, err := serve.userRepo.Get(globalcommonmodel.NewUserId(command.UserId))
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
		userIdDto := uuidutil.UnsafelyNewUuid(claims["jti"].(string))
		return userIdDto, nil
	} else {
		return userIdDto, errors.New("token is not valid")
	}
}
