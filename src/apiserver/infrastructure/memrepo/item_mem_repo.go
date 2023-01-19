package memrepo

import (
	"fmt"
	"os"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
	"github.com/google/uuid"
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
				itemmodel.NewItemAgg(stoneItemDefaultId, "stone", fmt.Sprintf("%s/assets/items/stone.png", serverUrl)),
				itemmodel.NewItemAgg(torchItemDefaultId, "torch", fmt.Sprintf("%s/assets/items/torch.png", serverUrl)),
			},
		}
		return itemMemRepoSingleton
	}
	return itemMemRepoSingleton
}

func (repo *itemMemRepo) GetAllItems() []itemmodel.ItemAgg {
	return repo.items
}
