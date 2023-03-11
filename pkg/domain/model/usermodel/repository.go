package usermodel

type Repository interface {
	Add(user UserAgg) error
}
