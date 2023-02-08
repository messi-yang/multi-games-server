package commonmemrepo

import (
	"errors"
	"fmt"
	"os"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
	"github.com/samber/lo"
)

type itemMemRepo struct {
	items []itemmodel.ItemAgg
}

var itemMemRepoSingleton *itemMemRepo

func NewItemMemRepo() itemmodel.Repo {
	if itemMemRepoSingleton == nil {
		stoneItemDefaultId := itemmodel.NewItemIdVo(0)
		torchItemDefaultId := itemmodel.NewItemIdVo(1)

		serverUrl := os.Getenv("SERVER_URL")

		itemMemRepoSingleton = &itemMemRepo{
			items: []itemmodel.ItemAgg{
				itemmodel.NewItemAgg(stoneItemDefaultId, "stone", false, fmt.Sprintf("%s/assets/items/stone.png", serverUrl)),
				itemmodel.NewItemAgg(torchItemDefaultId, "torch", true, fmt.Sprintf("%s/assets/items/torch.png", serverUrl)),
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
