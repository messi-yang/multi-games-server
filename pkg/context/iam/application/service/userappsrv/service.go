package userappsrv

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/usermodel"
)

type Service interface {
	GetUser(GetUserQuery) (userDto dto.UserDto, err error)
}

type serve struct {
	userRepo usermodel.UserRepo
}

func NewService(userRepo usermodel.UserRepo) Service {
	return &serve{
		userRepo: userRepo,
	}
}

func (serve *serve) GetUser(query GetUserQuery) (userDto dto.UserDto, err error) {
	userId := globalcommonmodel.NewUserId(query.UserId)
	user, err := serve.userRepo.Get(userId)
	if err != nil {
		return userDto, err
	}
	return dto.NewUserDto(user), nil
}
