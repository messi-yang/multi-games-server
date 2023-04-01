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
		return repository, err
	}
	return &itemRepository{gormDb: gormDb}, nil
}

func (repo *itemRepository) GetAll() (items []itemmodel.ItemAgg, err error) {
	var itemModels []psqlmodel.ItemModel
	result := repo.gormDb.Find(&itemModels)
	if result.Error != nil {
		err = result.Error
		return items, err
	}

	items = lo.Map(itemModels, func(model psqlmodel.ItemModel, _ int) itemmodel.ItemAgg {
		return model.ToAggregate()
	})
	return items, nil
}

func (repo *itemRepository) Get(itemId itemmodel.ItemIdVo) (item itemmodel.ItemAgg, err error) {
	itemModel := psqlmodel.ItemModel{Id: itemId.Uuid()}
	result := repo.gormDb.First(&itemModel)
	if result.Error != nil {
		return item, result.Error
	}

	return itemModel.ToAggregate(), nil
}

func (repo *itemRepository) GetFirstItem() (item itemmodel.ItemAgg, err error) {
	itemModel := psqlmodel.ItemModel{}
	result := repo.gormDb.First(&itemModel)
	if result.Error != nil {
		return item, result.Error
	}

	return itemModel.ToAggregate(), nil
}

func (repo *itemRepository) Add(item itemmodel.ItemAgg) error {
	itemModel := psqlmodel.NewItemModel(item)
	res := repo.gormDb.Create(&itemModel)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (repo *itemRepository) Update(item itemmodel.ItemAgg) error {
	itemModel := psqlmodel.NewItemModel(item)
	res := repo.gormDb.Save(&itemModel)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
