package postgres

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/usermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/infrastructure/persistence/postgres/psqlmodel"
	"gorm.io/gorm"
)

type userRepository struct {
	gormDb *gorm.DB
}

func NewUserRepository() (repository usermodel.Repository, err error) {
	gormDb, err := NewSession()
	if err != nil {
		return
	}
	repository = &userRepository{gormDb: gormDb}
	return
}

func (repo *userRepository) Add(user usermodel.UserAgg) (err error) {
	userModel := psqlmodel.NewUserModel(user)
	res := repo.gormDb.Create(&userModel)
	if res.Error != nil {
		err = res.Error
		return
	}
	return
}
