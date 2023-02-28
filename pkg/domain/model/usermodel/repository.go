package usermodel

type Repo interface {
	Add(user UserAgg) error
}
