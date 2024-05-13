package service

import (
	"errors"
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/uuidutil"
	"github.com/golang-jwt/jwt/v5"
)

type AuthService interface {
	GenerateAccessToken(userId globalcommonmodel.UserId) (accessToken string, err error)
	ValidateAccessToken(accessToken string) (userId globalcommonmodel.UserId, err error)
}

type authServe struct {
	authSecret string
}

func NewAuthService(authSecret string) AuthService {
	return &authServe{authSecret: authSecret}
}

func (authServe *authServe) GenerateAccessToken(userId globalcommonmodel.UserId) (accessToken string, err error) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		ID:        userId.Uuid().String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err = token.SignedString([]byte(authServe.authSecret))
	if err != nil {
		return accessToken, err
	}
	return accessToken, nil
}

func (authServe *authServe) ValidateAccessToken(accessToken string) (userId globalcommonmodel.UserId, err error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(authServe.authSecret), nil
	})
	if err != nil {
		return userId, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userIdDto := uuidutil.UnsafelyNewUuid(claims["jti"].(string))
		return globalcommonmodel.NewUserId(userIdDto), nil
	} else {
		return userId, errors.New("token is not valid")
	}
}
