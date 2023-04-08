package gamermodel

type Repository interface {
	GetAll() ([]GamerAgg, error)
}
