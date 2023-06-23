package userappsrv

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/domain/model/identitymodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
	"github.com/samber/lo"
)

type Service interface {
	GetUserByEmailAddress(GetUserByEmailAddressQuery) (userDto *dto.UserDto, err error)
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

func (serve *serve) GetUserByEmailAddress(query GetUserByEmailAddressQuery) (*dto.UserDto, error) {
	emailAddress, err := sharedkernelmodel.NewEmailAddress(query.EmailAddress)
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

func (serve *serve) GetUserQuery(query GetUserQuery) (userDto dto.UserDto, err error) {
	userId := sharedkernelmodel.NewUserId(query.UserId)
	user, err := serve.userRepo.Get(userId)
	if err != nil {
		return userDto, err
	}
	return dto.NewUserDto(user), nil
}
