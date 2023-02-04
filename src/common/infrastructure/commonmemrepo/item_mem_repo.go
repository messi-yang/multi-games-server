package commonmemrepo

import (
	"errors"
	"fmt"
	"os"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type itemMemRepo struct {
	items []itemmodel.ItemAgg
}

var itemMemRepoSingleton *itemMemRepo

func NewItemMemRepo() itemmodel.Repo {
	if itemMemRepoSingleton == nil {
		stoneItemDefaultId, _ := itemmodel.NewItemIdVo("2a0c8f7f-48dc-4553-86b4-2bbc3786bb66")
		torchItemDefaultId, _ := itemmodel.NewItemIdVo("31c8cc9e-42a3-4a42-86d2-905ca37305ba")

		fmt.Println(uuid.New())
		fmt.Println(uuid.New())
		fmt.Println(uuid.New())

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
