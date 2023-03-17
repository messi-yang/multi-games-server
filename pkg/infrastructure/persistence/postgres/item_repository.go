package postgres

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/infrastructure/persistence/postgres/psqlmodel"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type itemRepository struct {
	gormDb *gorm.DB
}

func NewItemRepository() (itemmodel.Repository, error) {
	gormDb, err := NewSession()
	if err != nil {
		return nil, err
	}
	return &itemRepository{gormDb: gormDb}, nil
}

func (repo *itemRepository) GetAll() ([]itemmodel.ItemAgg, error) {
	var itemModels []psqlmodel.ItemModel
	result := repo.gormDb.Find(&itemModels)
	if result.Error != nil {
		return nil, result.Error
	}

	worlds := lo.Map(itemModels, func(model psqlmodel.ItemModel, _ int) itemmodel.ItemAgg {
		return model.ToAggregate()
	})

	return worlds, nil
}

func (repo *itemRepository) Get(itemId itemmodel.ItemIdVo) (itemmodel.ItemAgg, error) {
	itemModel := psqlmodel.ItemModel{Id: itemId.Uuid()}
	result := repo.gormDb.First(&itemModel)
	if result.Error != nil {
		return itemmodel.ItemAgg{}, result.Error
	}

	return itemModel.ToAggregate(), nil
}
