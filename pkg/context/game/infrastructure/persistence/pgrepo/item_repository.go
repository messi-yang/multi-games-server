package pgrepo

import (
	"fmt"
	"os"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/pgmodel"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

func newItemModel(item itemmodel.Item) pgmodel.ItemModel {
	return pgmodel.ItemModel{
		Id:           item.GetId().Uuid(),
		Name:         item.GetName(),
		Traversable:  item.GetTraversable(),
		ThumbnailSrc: item.GetThumbnailSrc(),
		ModelSrc:     item.GetModelSrc(),
	}
}

func parseItemModel(itemModel pgmodel.ItemModel) itemmodel.Item {
	serverUrl := os.Getenv("SERVER_URL")
	return itemmodel.NewItem(
		commonmodel.NewItemId(itemModel.Id),
		itemModel.Name,
		itemModel.Traversable,
		fmt.Sprintf("%s%s", serverUrl, itemModel.ThumbnailSrc),
		fmt.Sprintf("%s%s", serverUrl, itemModel.ModelSrc),
	)
}

type itemRepo struct {
	db *gorm.DB
}

func NewItemRepo() (repository itemmodel.Repo, err error) {
	db, err := pgmodel.NewClient()
	if err != nil {
		return repository, err
	}
	return &itemRepo{db: db}, nil
}

func (repo *itemRepo) GetAll() (items []itemmodel.Item, err error) {
	var itemModels []pgmodel.ItemModel
	result := repo.db.Find(&itemModels)
	if result.Error != nil {
		err = result.Error
		return items, err
	}

	items = lo.Map(itemModels, func(itemModel pgmodel.ItemModel, _ int) itemmodel.Item {
		return parseItemModel(itemModel)
	})
	return items, nil
}

func (repo *itemRepo) Get(itemId commonmodel.ItemId) (item itemmodel.Item, err error) {
	itemModel := pgmodel.ItemModel{Id: itemId.Uuid()}
	result := repo.db.First(&itemModel)
	if result.Error != nil {
		return item, result.Error
	}

	return parseItemModel(itemModel), nil
}

func (repo *itemRepo) GetFirstItem() (item itemmodel.Item, err error) {
	itemModel := pgmodel.ItemModel{}
	result := repo.db.First(&itemModel)
	if result.Error != nil {
		return item, result.Error
	}

	return parseItemModel(itemModel), nil
}

func (repo *itemRepo) Add(item itemmodel.Item) error {
	itemModel := newItemModel(item)
	res := repo.db.Create(&itemModel)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (repo *itemRepo) Update(item itemmodel.Item) error {
	itemModel := newItemModel(item)
	res := repo.db.Save(&itemModel)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
