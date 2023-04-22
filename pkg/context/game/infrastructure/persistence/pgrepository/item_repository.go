package pgrepository

import (
	"fmt"
	"os"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/pgmodel"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

func newItemModel(item itemmodel.ItemAgg) pgmodel.ItemModel {
	return pgmodel.ItemModel{
		Id:           item.GetId().Uuid(),
		Name:         item.GetName(),
		Traversable:  item.GetTraversable(),
		ThumbnailSrc: item.GetThumbnailSrc(),
		ModelSrc:     item.GetModelSrc(),
	}
}

func parseItemModel(itemModel pgmodel.ItemModel) itemmodel.ItemAgg {
	serverUrl := os.Getenv("SERVER_URL")
	return itemmodel.NewItemAgg(
		commonmodel.NewItemIdVo(itemModel.Id),
		itemModel.Name,
		itemModel.Traversable,
		fmt.Sprintf("%s%s", serverUrl, itemModel.ThumbnailSrc),
		fmt.Sprintf("%s%s", serverUrl, itemModel.ModelSrc),
	)
}

type itemRepository struct {
	dbClient *gorm.DB
}

func NewItemRepository() (repository itemmodel.Repository, err error) {
	dbClient, err := pgmodel.NewClient()
	if err != nil {
		return repository, err
	}
	return &itemRepository{dbClient: dbClient}, nil
}

func (repo *itemRepository) GetAll() (items []itemmodel.ItemAgg, err error) {
	var itemModels []pgmodel.ItemModel
	result := repo.dbClient.Find(&itemModels)
	if result.Error != nil {
		err = result.Error
		return items, err
	}

	items = lo.Map(itemModels, func(itemModel pgmodel.ItemModel, _ int) itemmodel.ItemAgg {
		return parseItemModel(itemModel)
	})
	return items, nil
}

func (repo *itemRepository) Get(itemId commonmodel.ItemIdVo) (item itemmodel.ItemAgg, err error) {
	itemModel := pgmodel.ItemModel{Id: itemId.Uuid()}
	result := repo.dbClient.First(&itemModel)
	if result.Error != nil {
		return item, result.Error
	}

	return parseItemModel(itemModel), nil
}

func (repo *itemRepository) GetFirstItem() (item itemmodel.ItemAgg, err error) {
	itemModel := pgmodel.ItemModel{}
	result := repo.dbClient.First(&itemModel)
	if result.Error != nil {
		return item, result.Error
	}

	return parseItemModel(itemModel), nil
}

func (repo *itemRepository) Add(item itemmodel.ItemAgg) error {
	itemModel := newItemModel(item)
	res := repo.dbClient.Create(&itemModel)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (repo *itemRepository) Update(item itemmodel.ItemAgg) error {
	itemModel := newItemModel(item)
	res := repo.dbClient.Save(&itemModel)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
