package userappsrv

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/application/dto"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/iam/domain/model/identitymodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
)

type Service interface {
	FindUserByEmailAddress(FindUserByEmailAddressQuery) (userDto dto.UserDto, found bool, err error)
	GetUserQuery(GetUserQuery) (userDto dto.UserDto, err error)
}

type serve struct {
	userRepo identitymodel.UserRepo
}

func NewService(userRepo identitymodel.UserRepo) Service {
	return &serve{
		userRepo: userRepo,
	}
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

func (serve *serve) GetUserQuery(query GetUserQuery) (userDto dto.UserDto, err error) {
	userId := sharedkernelmodel.NewUserId(query.UserId)
	user, err := serve.userRepo.Get(userId)
	if err != nil {
		return userDto, err
	}
	return dto.NewUserDto(user), nil
}
