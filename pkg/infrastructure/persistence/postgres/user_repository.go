package postgres

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/usermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/infrastructure/persistence/postgres/psqlmodel"
	"gorm.io/gorm"
)

type userRepository struct {
	gormDb *gorm.DB
}

func NewUserRepository() (usermodel.Repository, error) {
	gormDb, err := NewSession()
	if err != nil {
		return nil, err
	}
	return &userRepository{gormDb: gormDb}, nil
}

func (repo *userRepository) Add(user usermodel.UserAgg) error {
	userModel := psqlmodel.NewUserModel(user)
	res := repo.gormDb.Create(&userModel)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
