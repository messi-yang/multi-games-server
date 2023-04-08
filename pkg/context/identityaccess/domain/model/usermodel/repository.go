package usermodel

type Repository interface {
	Add(UserAgg) error
}
