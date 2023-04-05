package postgres

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/usermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/postgres/psqlmodel"
	"gorm.io/gorm"
)

type userRepository struct {
	dbClient *gorm.DB
}

func NewUserRepository() (repository usermodel.Repository, err error) {
	dbClient, err := NewDbClient()
	if err != nil {
		return repository, err
	}
	return &userRepository{dbClient: dbClient}, nil
}

func (repo *userRepository) Add(user usermodel.UserAgg) error {
	userModel := psqlmodel.NewUserModel(user)
	res := repo.dbClient.Create(&userModel)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
