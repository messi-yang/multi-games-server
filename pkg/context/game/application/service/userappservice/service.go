package userappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/jsondto"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/usermodel"
	"github.com/samber/lo"
)

type Service interface {
	GetUsers(GetUsersQuery) ([]jsondto.UserAggDto, error)
}

type serve struct {
	userRepository usermodel.Repository
}

func NewService(userRepository usermodel.Repository) Service {
	return &serve{
		userRepository: userRepository,
	}
}

func (serve *serve) GetUsers(query GetUsersQuery) (itemDtos []jsondto.UserAggDto, err error) {
	users, err := serve.userRepository.GetAll()
	if err != nil {
		return itemDtos, err
	}

	return lo.Map(users, func(user usermodel.UserAgg, _ int) jsondto.UserAggDto {
		return jsondto.NewUserAggDto(user)
	}), nil
}
