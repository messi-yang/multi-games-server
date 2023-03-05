package memrepo

import (
	"errors"
	"fmt"
	"os"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
	"github.com/samber/lo"
)

type itemMemRepo struct {
	items []itemmodel.ItemAgg
}

var itemMemRepoSingleton *itemMemRepo

func NewItemMemRepo() itemmodel.Repo {
	if itemMemRepoSingleton == nil {
		stoneItemDefaultId, _ := itemmodel.ParseItemIdVo("3c28537a-80c2-4ac1-917b-b1cd517c6b5e")
		torchItemDefaultId, _ := itemmodel.ParseItemIdVo("34af14ab-42c5-4c55-a787-44f32012354e")
		treeItemDefaultId, _ := itemmodel.ParseItemIdVo("414b5703-91d1-42fc-a007-36dd8f25e329")

		serverUrl := os.Getenv("SERVER_URL")

		itemMemRepoSingleton = &itemMemRepo{
			items: []itemmodel.ItemAgg{
				itemmodel.NewItemAgg(stoneItemDefaultId, "stone", false, fmt.Sprintf("%s/assets/items/stone.png", serverUrl), "/items/stone.gltf"),
				itemmodel.NewItemAgg(torchItemDefaultId, "torch", true, fmt.Sprintf("%s/assets/items/torch.png", serverUrl), "/items/torch.gltf"),
				itemmodel.NewItemAgg(treeItemDefaultId, "tree", false, fmt.Sprintf("%s/assets/items/tree.png", serverUrl), "/items/tree.gltf"),
			},
		}
		return itemMemRepoSingleton
	}
	return itemMemRepoSingleton
}

func (repo *itemMemRepo) GetAll() []itemmodel.ItemAgg {
	return repo.items
}

func (repo *itemMemRepo) Get(itemId itemmodel.ItemIdVo) (itemmodel.ItemAgg, error) {
	item, found := lo.Find(repo.items, func(item itemmodel.ItemAgg) bool {
		return item.GetId().IsEqual(itemId)
	})
	if !found {
		return itemmodel.ItemAgg{}, errors.New("item with given id not found")
	}
	return item, nil
}
