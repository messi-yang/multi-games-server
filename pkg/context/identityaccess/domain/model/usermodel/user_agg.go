package usermodel

import "github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"

type UserAgg struct {
	id           sharedkernelmodel.UserIdVo
	emailAddress string
	username     string
}

func NewUserAgg(
	id sharedkernelmodel.UserIdVo,
	emailAddress string,
	username string,
) UserAgg {
	return UserAgg{id: id, emailAddress: emailAddress, username: username}
}

func (agg *UserAgg) GetId() sharedkernelmodel.UserIdVo {
	return agg.id
}

func (agg *UserAgg) GetEmailAddress() string {
	return agg.emailAddress
}

func (agg *UserAgg) GetUsername() string {
	return agg.username
}
