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

func NewItemRepository() (repository itemmodel.Repository, err error) {
	gormDb, err := NewSession()
	if err != nil {
		return
	}
	repository = &itemRepository{gormDb: gormDb}
	return
}

func (repo *itemRepository) GetAll() (items []itemmodel.ItemAgg, err error) {
	var itemModels []psqlmodel.ItemModel
	result := repo.gormDb.Find(&itemModels)
	if result.Error != nil {
		err = result.Error
		return
	}

	items = lo.Map(itemModels, func(model psqlmodel.ItemModel, _ int) itemmodel.ItemAgg {
		return model.ToAggregate()
	})
	return
}

func (repo *itemRepository) Get(itemId itemmodel.ItemIdVo) (item itemmodel.ItemAgg, err error) {
	itemModel := psqlmodel.ItemModel{Id: itemId.Uuid()}
	result := repo.gormDb.First(&itemModel)
	if result.Error != nil {
		err = result.Error
		return
	}

	item = itemModel.ToAggregate()
	return
}
