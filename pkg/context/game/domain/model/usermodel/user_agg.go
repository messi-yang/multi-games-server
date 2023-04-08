package usermodel

import "github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/identityaccess/domain/model/usermodel"

type UserAgg struct {
	id     UserIdVo
	userId usermodel.UserIdVo
}

func NewUserAgg(
	id UserIdVo,
	userId usermodel.UserIdVo,
) UserAgg {
	return UserAgg{id: id, userId: userId}
}

func (agg *UserAgg) GetId() UserIdVo {
	return agg.id
}

func (agg *UserAgg) GetUserId() usermodel.UserIdVo {
	return agg.userId
}
