package authappsrv

import (
	"errors"

	"github.com/dum-dum-genius/zossi-server/pkg/util/uuidutil"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Service interface {
	Validate(accessToken string) (userIdDto uuid.UUID, err error)
}

type serve struct {
	authSecret string
}

func NewService(authSecret string) Service {
	return &serve{authSecret: authSecret}
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
