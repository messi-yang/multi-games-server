package postgres

import (
	"errors"
	"fmt"
	"os"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type itemRepository struct {
	gormDb *gorm.DB
	items  []itemmodel.ItemAgg
}

var itemRepositorySingleton *itemRepository

func NewItemRepository() (itemmodel.Repository, error) {
	if itemRepositorySingleton != nil {
		return itemRepositorySingleton, nil
	} else {
		gormDb, err := NewSession()
		if err != nil {
			return nil, err
		}
		stoneItemDefaultId, _ := itemmodel.ParseItemIdVo("3c28537a-80c2-4ac1-917b-b1cd517c6b5e")
		torchItemDefaultId, _ := itemmodel.ParseItemIdVo("34af14ab-42c5-4c55-a787-44f32012354e")
		treeItemDefaultId, _ := itemmodel.ParseItemIdVo("414b5703-91d1-42fc-a007-36dd8f25e329")

		serverUrl := os.Getenv("SERVER_URL")

		itemRepositorySingleton = &itemRepository{
			gormDb: gormDb,
			items: []itemmodel.ItemAgg{
				itemmodel.NewItemAgg(stoneItemDefaultId, "stone", false, fmt.Sprintf("%s/asset/item/stone/thumbnail.png", serverUrl), fmt.Sprintf("%s/asset/item/stone/model.gltf", serverUrl)),
				itemmodel.NewItemAgg(torchItemDefaultId, "torch", true, fmt.Sprintf("%s/asset/item/torch/thumbnail.png", serverUrl), fmt.Sprintf("%s/asset/item/torch/model.gltf", serverUrl)),
				itemmodel.NewItemAgg(treeItemDefaultId, "tree", false, fmt.Sprintf("%s/asset/item/tree/thumbnail.png", serverUrl), fmt.Sprintf("%s/asset/item/tree/model.gltf", serverUrl)),
			},
		}
		return itemRepositorySingleton, nil
	}
}

func (repo *itemRepository) GetAll() ([]itemmodel.ItemAgg, error) {
	return repo.items, nil
}

func (repo *itemRepository) Get(itemId itemmodel.ItemIdVo) (itemmodel.ItemAgg, error) {
	item, found := lo.Find(repo.items, func(item itemmodel.ItemAgg) bool {
		return item.GetId().IsEqual(itemId)
	})
	if !found {
		return itemmodel.ItemAgg{}, errors.New("item with given id not found")
	}
	return item, nil
}
