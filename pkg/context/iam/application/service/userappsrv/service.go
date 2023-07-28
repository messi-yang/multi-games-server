package userappsrv

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/usermodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type Service interface {
	GetUserByEmailAddress(GetUserByEmailAddressQuery) (userDto *dto.UserDto, err error)
	GetUser(GetUserQuery) (userDto dto.UserDto, err error)
	GetUsersOfIds(GetUsersOfIdsQuery) (userDtoMap map[uuid.UUID]dto.UserDto, err error)
	UpdateUser(UpdateUserCommand) (err error)
}

type serve struct {
	userRepo usermodel.UserRepo
}

func NewService(userRepo usermodel.UserRepo) Service {
	return &serve{
		userRepo: userRepo,
	}
}

func (serve *serve) GetUserByEmailAddress(query GetUserByEmailAddressQuery) (*dto.UserDto, error) {
	emailAddress, err := globalcommonmodel.NewEmailAddress(query.EmailAddress)
	if err != nil {
		return nil, err
	}

	user, err := serve.userRepo.GetUserByEmailAddress(emailAddress)
	if err != nil {
		return nil, err
	}
	return lo.TernaryF(
		user == nil,
		func() *dto.UserDto {
			return nil
		},
		func() *dto.UserDto {
			return commonutil.ToPointer(dto.NewUserDto(*user))
		},
	), nil
}

func (serve *serve) GetUser(query GetUserQuery) (userDto dto.UserDto, err error) {
	userId := globalcommonmodel.NewUserId(query.UserId)
	user, err := serve.userRepo.Get(userId)
	if err != nil {
		return userDto, err
	}
	return dto.NewUserDto(user), nil
}

func (serve *serve) GetUsersOfIds(query GetUsersOfIdsQuery) (userDtoMap map[uuid.UUID]dto.UserDto, err error) {
	userIds := lo.Map(query.UserIds, func(userIdDto uuid.UUID, _ int) globalcommonmodel.UserId {
		return globalcommonmodel.NewUserId(userIdDto)
	})
	users, err := serve.userRepo.GetUsersOfIds(userIds)
	if err != nil {
		return userDtoMap, err
	}

	userDtos := lo.Map(users, func(user usermodel.User, _ int) dto.UserDto {
		return dto.NewUserDto(user)
	})

	return lo.KeyBy(userDtos, func(userDto dto.UserDto) uuid.UUID {
		return userDto.Id
	}), nil
}

func (serve *serve) UpdateUser(command UpdateUserCommand) (err error) {
	fmt.Println(command.Username)
	userId := globalcommonmodel.NewUserId(command.UserId)
	username, err := globalcommonmodel.NewUsername(command.Username)
	if err != nil {
		return err
	}

	user, err := serve.userRepo.Get(userId)
	if err != nil {
		return err
	}

	user.UpdateUsername(username)

	return serve.userRepo.Update(user)
}
