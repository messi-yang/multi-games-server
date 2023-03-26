package usermodel

type UserAgg struct {
	id           UserIdVo
	emailAddress string
	username     string
}

func NewUserAgg(
	id UserIdVo,
	emailAddress string,
	username string,
) UserAgg {
	return UserAgg{id: id, emailAddress: emailAddress, username: username}
}

func (agg *UserAgg) GetId() UserIdVo {
	return agg.id
}

func (agg *UserAgg) GetEmailAddress() string {
	return agg.emailAddress
}

func (agg *UserAgg) GetUsername() string {
	return agg.username
}
