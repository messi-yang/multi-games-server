package usermodel

type Repository interface {
	GetAll() ([]UserAgg, error)
}
