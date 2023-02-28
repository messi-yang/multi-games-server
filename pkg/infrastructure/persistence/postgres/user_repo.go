package postgres

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/usermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/infrastructure/persistence/postgres/psqlmodel"
	"gorm.io/gorm"
)

type userRepo struct {
	gormDb *gorm.DB
}

func NewUserRepo() (usermodel.Repo, error) {
	gormDb, err := NewSession()
	if err != nil {
		return nil, err
	}
	return &userRepo{gormDb: gormDb}, nil
}

func (repo *userRepo) Add(user usermodel.UserAgg) error {
	userModel := psqlmodel.NewUserModel(user)
	res := repo.gormDb.Create(&userModel)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
